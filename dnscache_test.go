package dnscache

import (
	"github.com/markchadwick/spec"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

var _ = spec.Suite("DNS Cache", func(c *spec.C) {
	c.It("should provide a net/http RoundTripper", func(c *spec.C) {
		c.Skip("not implemented")
	})
})

func Test(t *testing.T) {
	spec.Run(t)
}
