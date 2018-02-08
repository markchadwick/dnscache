package dnscache

import (
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/rcrowley/go-metrics"
)

var (
	cacheHit      = metrics.NewMeter()
	cacheMiss     = metrics.NewMeter()
	lookupErr     = metrics.NewMeter()
	lookupRecover = metrics.NewMeter()
)

func init() {
	metrics.DefaultRegistry.Register("dnscache.cache.hit", cacheHit)
	metrics.DefaultRegistry.Register("dnscache.cache.miss", cacheMiss)
	metrics.DefaultRegistry.Register("dnscache.lookup.err", lookupErr)
	metrics.DefaultRegistry.Register("dnscache.lookup.recover", lookupRecover)
}

type record struct {
	addrs []string
	next  chan string
}

func Record(ttl time.Duration, addrs []string) *record {
	rec := &record{
		addrs: make([]string, 0),
		next:  make(chan string),
	}
	for _, addr := range addrs {
		if !strings.Contains(addr, ":") {
			rec.addrs = append(rec.addrs, addr)
		}
	}

	go func() {
		i := 0
		numAddrs := len(rec.addrs)
		expired := time.After(ttl)
		for {
			select {
			case <-expired:
				close(rec.next)
				return
			default:
				if i >= numAddrs {
					i = 0
				}
				rec.next <- rec.addrs[i]
				i += 1
			}
		}
	}()

	return rec
}

// Returns the next address to use if the record is not expired.
func (r *record) NextAddr() (addr string, expired bool) {
	addr, ok := <-r.next
	// if the channel is closed, the record has expired
	if !ok {
		return "", true
	}
	return addr, false
}

// Returns a random address ignoring the record expiry.
func (r *record) RandomAddr() string {
	return r.addrs[rand.Intn(len(r.addrs))]
}

type Cache struct {
	recs map[string]*record
	ttl  time.Duration
	sync.RWMutex
}

func New(ttl time.Duration) *Cache {
	return &Cache{
		recs: make(map[string]*record),
		ttl:  ttl,
	}
}

func (c *Cache) LookupHost(host string) (addr string, err error) {
	c.RLock()
	rec, haveRecord := c.recs[host]
	c.RUnlock()
	if haveRecord {
		// if we have a cached & active record, return the next address
		if cached, expired := rec.NextAddr(); !expired {
			cacheHit.Mark(1)
			return cached, nil
		}
	}

	// Value is not cached, look it up. We synchronize this section to prevent
	// many goroutines from doing redundant lookups.
	c.Lock()
	defer c.Unlock()
	hosts, err := net.LookupHost(host)

	// In the case there was an error looking up the value AND we have a cached
	// value (that has expired), just return the cached value.
	if err != nil && haveRecord {
		lookupRecover.Mark(1)
		return rec.RandomAddr(), nil
	}

	// no cache and the lookup failed, just proxy the error
	if err != nil {
		lookupErr.Mark(1)
		return "", err
	}

	// store the new record
	rec = Record(c.ttl, hosts)
	c.recs[host] = rec
	cacheMiss.Mark(1)
	addr, _ = rec.NextAddr()
	return addr, nil
}
