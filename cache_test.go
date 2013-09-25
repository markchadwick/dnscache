package dnscache

import (
	"github.com/markchadwick/spec"
	"time"
)

var _ = spec.Suite("Record", func(c *spec.C) {
	c.It("should cycle throught hosts", func(c *spec.C) {
		rec := Record(time.Second, []string{
			"192.168.1.1",
			"192.168.1.2",
			"192.168.1.3",
		})
		c.Assert(rec.Addr()).Equals("192.168.1.1")
		c.Assert(rec.Addr()).Equals("192.168.1.2")
		c.Assert(rec.Addr()).Equals("192.168.1.3")
		c.Assert(rec.Addr()).Equals("192.168.1.1")
		c.Assert(rec.Addr()).Equals("192.168.1.2")
		c.Assert(rec.Addr()).Equals("192.168.1.3")
		c.Assert(rec.Addr()).Equals("192.168.1.1")
	})

	c.It("should know when its inside its TTL", func(c *spec.C) {
		rec := Record(time.Minute, []string{"192.168.1.1"})
		c.Assert(rec.Expired()).IsFalse()
		rec.expires = time.Now()
		c.Assert(rec.Expired()).IsTrue()
	})

	c.It("should (probably) ignore ipv6 hosts", func(c *spec.C) {
		rec := Record(time.Second, []string{
			"74.125.228.32",
			"2607:f8b0:4004:801::1003",
		})
		c.Assert(rec.addrs).HasLen(1)
		c.Assert(rec.Addr()).Equals("74.125.228.32")
		c.Assert(rec.Addr()).Equals("74.125.228.32")
	})
})

var _ = spec.Suite("Host Lookup", func(c *spec.C) {
	cache := New(1 * time.Minute)
	hit := cacheHit.Count()
	miss := cacheMiss.Count()

	c.It("should cache an address", func(c *spec.C) {
		_, err := cache.LookupHost("google.com")
		c.Assert(err).IsNil()

		c.Assert(cacheHit.Count()).Equals(hit)
		c.Assert(cacheMiss.Count()).Equals(miss + 1)

		_, err = cache.LookupHost("google.com")
		c.Assert(err).IsNil()

		c.Assert(cacheHit.Count()).Equals(hit + 1)
		c.Assert(cacheMiss.Count()).Equals(miss + 1)
	})
})
