package dnscache

import (
	"github.com/markchadwick/spec"
	"time"
)

func nextAddr(r *record) string {
	addr, _ := r.NextAddr()
	return addr
}

func isExpired(r *record) bool {
	timeout := time.After(time.Second)
	// checking for expiry is a little racey since time.After does not fire
	// instantly for a 0 TTL. So we check in loop for up to a second.
	for {
		select {
		case <-timeout:
			return false
		default:
			if _, isExpired := r.NextAddr(); isExpired {
				return isExpired
			}
		}
	}
}

var _ = spec.Suite("Record", func(c *spec.C) {
	c.It("should cycle throught hosts", func(c *spec.C) {
		rec := Record(time.Second, []string{
			"192.168.1.1",
			"192.168.1.2",
			"192.168.1.3",
		})
		c.Assert(nextAddr(rec)).Equals("192.168.1.1")
		c.Assert(nextAddr(rec)).Equals("192.168.1.2")
		c.Assert(nextAddr(rec)).Equals("192.168.1.3")
		c.Assert(nextAddr(rec)).Equals("192.168.1.1")
		c.Assert(nextAddr(rec)).Equals("192.168.1.2")
		c.Assert(nextAddr(rec)).Equals("192.168.1.3")
		c.Assert(nextAddr(rec)).Equals("192.168.1.1")
	})

	c.It("should know when its inside its TTL", func(c *spec.C) {
		rec := Record(time.Minute, []string{"192.168.1.1"})
		c.Assert(isExpired(rec)).IsFalse()
		rec = Record(time.Duration(-1), []string{"192.168.1.1"})
		c.Assert(isExpired(rec)).IsTrue()
	})

	c.It("should (probably) ignore ipv6 hosts", func(c *spec.C) {
		rec := Record(time.Second, []string{
			"74.125.228.32",
			"2607:f8b0:4004:801::1003",
		})
		c.Assert(rec.addrs).HasLen(1)
		c.Assert(nextAddr(rec)).Equals("74.125.228.32")
		c.Assert(nextAddr(rec)).Equals("74.125.228.32")
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

	c.It("should lookup an IP", func(c *spec.C) {
		_, err := cache.LookupHost("192.168.1.1")
		c.Assert(err).IsNil()
	})

	c.It("should look up localhost", func(c *spec.C) {
		addr, err := cache.LookupHost("localhost")
		c.Assert(err).IsNil()
		c.Assert(addr).Equals("127.0.0.1")
	})
})
