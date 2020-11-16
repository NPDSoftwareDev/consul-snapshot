package consul

import (
	"log"
	"strings"

	consulapi "github.com/hashicorp/consul/api"
)

// Consul struct is used for consul client such as the client
// and the actual key data.
type Consul struct {
	Client     consulapi.Client
	KeyData    consulapi.KVPairs
	KeyDataLen int
	PQData     []*consulapi.PreparedQueryDefinition
	PQDataLen  int
	ACLData    []*consulapi.ACLTokenListEntry
	ACLDataLen int
}

// Client creates a consul client from the consul api
func Client() *consulapi.Client {
	consul, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		log.Fatalf("[ERR] Unable to create a consul client: %v", err)
	}
	return consul
}

// ListKeys lists all the keys from consul with no prefix.
func (c *Consul) ListKeys() {
	listOpt := &consulapi.QueryOptions{
		AllowStale:        false,
		RequireConsistent: true,
	}
	keys, _, err := c.Client.KV().List("/", listOpt)
	if err != nil {
		log.Fatalf("[ERR] Unable to list keys: %v", err)
	}
	c.KeyData = keys
	c.KeyDataLen = len(keys)
}

// ListPQs lists all the prepared queries from consul
func (c *Consul) ListPQs() {
	listOpt := &consulapi.QueryOptions{
		AllowStale:        false,
		RequireConsistent: true,
	}
	pqs, _, err := c.Client.PreparedQuery().List(listOpt)
	if err != nil {
		log.Fatalf("[ERR] Unable to list PQs: %v", err)
	}

	c.PQData = pqs
	c.PQDataLen = len(pqs)
}

// ListTokens lists all the prepared queries from consul
func (c *Consul) ListTokens() {
	listOpt := &consulapi.QueryOptions{
		AllowStale:        false,
		RequireConsistent: true,
	}

	acls, _, err := c.Client.ACL().TokenList(listOpt)
	if err != nil {
		// Really don't like this but seems to be the only way to detect
		if strings.Contains(err.Error(), "401 (ACL support disabled)") {
			log.Print("[INFO] ACL support detected as disbaled, skipping")
			c.ACLData = []*consulapi.ACLTokenListEntry{}
			c.ACLDataLen = 0
		} else {
			log.Fatalf("[ERR] Unable to list ACL tokens: %v", err)
		}
	} else {
		c.ACLData = acls
		c.ACLDataLen = len(acls)
	}

}
