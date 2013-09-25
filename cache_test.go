package dnscache

import (
	"github.com/markchadwick/spec"
	"log"
	"time"
)

var _ = spec.Suite("Record", func(c *spec.C) {
	c.It("should cycle throught hosts", func(c *spec.C) {
		rec := Record([]string{
			"192.168.1.1",
			"192.168.1.2",
			"192.168.1.3",
		})
		c.Assert(rec.Host()).Equals("192.168.1.1")
		c.Assert(rec.Host()).Equals("192.168.1.2")
		c.Assert(rec.Host()).Equals("192.168.1.3")
		c.Assert(rec.Host()).Equals("192.168.1.1")
		c.Assert(rec.Host()).Equals("192.168.1.2")
		c.Assert(rec.Host()).Equals("192.168.1.3")
		c.Assert(rec.Host()).Equals("192.168.1.1")
	})

	c.It("should know when its inside its TTL", func(c *spec.C) {
		c.Skip("not implemented")
	})

	c.It("should (probably) ignore ipv6 hosts", func(c *spec.C) {
		c.Skip("not implemented")
	})
})

var _ = spec.Suite("Host Lookup", func(c *spec.C) {
	cache := New(1 * time.Minute)

	c.It("should look up a basic address", func(c *spec.C) {
		name, err := cache.LookupHost("google.com")
		c.Assert(err).IsNil()

		log.Printf("-------------------------------------")
		log.Printf("Name: %#v", name)
		log.Printf("Err: %#v", err)
		log.Printf("-------------------------------------")
	})
})
