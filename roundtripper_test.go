package dnscache

import (
	"github.com/markchadwick/spec"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = spec.Suite("Roundtripper", func(c *spec.C) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer server.Close()
	client := &http.Client{Transport: DnsCachedRoundTripper(time.Minute)}

	hits := cacheHit.Count()
	misses := cacheMiss.Count()

	c.It("should proxy DNS queries", func(c *spec.C) {
		_, err := client.Get(server.URL)
		c.Assert(err).IsNil()
		c.Assert(cacheHit.Count()).Equals(hits)
		c.Assert(cacheMiss.Count()).Equals(misses + 1)
	})

	c.It("should cache DNS queries", func(c *spec.C) {
		cache := &Cache{
			recs: map[string]*record{
				"127.0.0.1": Record(time.Minute, []string{"127.0.0.1"}),
			},
			ttl: time.Minute,
		}
		client := &http.Client{Transport: RoundTripper(cache)}

		_, err := client.Get(server.URL)
		c.Assert(err).IsNil()
		c.Assert(cacheHit.Count()).Equals(hits + 1)
		c.Assert(cacheMiss.Count()).Equals(misses)
	})
})
