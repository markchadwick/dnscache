package dnscache

import (
	"net"
	"time"
)

type record struct {
	hosts []string
	next  chan string
}

func Record(hosts []string) *record {
	// TODO: ipv6?
	rec := &record{
		hosts: hosts,
		next:  make(chan string),
	}
	go func() {
		i := 0
		for {
			if i >= len(hosts) {
				i = 0
			}
			rec.next <- rec.hosts[i]
			i += 1
		}
	}()
	return rec
}

func (r *record) Host() string {
	return <-r.next
}

type Cache struct {
	hosts map[string]*record
	ttl   time.Duration
}

func New(ttl time.Duration) *Cache {
	return &Cache{
		hosts: make(map[string]*record),
		ttl:   ttl,
	}
}

// Analogous to net.LookupHost. It looks up the given host using the local
// resolver. It returns an array of that host's addresses.
func (c *Cache) LookupHost(addr string) (name []string, err error) {
	return net.LookupHost(addr)
}
