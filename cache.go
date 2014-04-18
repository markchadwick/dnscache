package dnscache

import (
	"github.com/rcrowley/go-metrics"
	"net"
	"strings"
	"sync"
	"time"
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
	addrs   []string
	next    chan string
	expires time.Time
	stop    chan bool
}

func Record(ttl time.Duration, addrs []string) *record {
	rec := &record{
		addrs:   make([]string, 0),
		next:    make(chan string),
		expires: time.Now().Add(ttl),
		stop:    make(chan bool),
	}
	for _, addr := range addrs {
		if !strings.Contains(addr, ":") {
			rec.addrs = append(rec.addrs, addr)
		}
	}

	go func() {
		i := 0
		numAddrs := len(rec.addrs)
		for {
			select {
			default:
				if i >= numAddrs {
					i = 0
				}
				rec.next <- rec.addrs[i]
				i += 1
			case <-rec.stop:
				return
			}
		}
	}()

	return rec
}

func (r *record) Addr() string {
	return <-r.next
}

func (r *record) Expired() bool {
	valid := time.Now().Before(r.expires)
	if !valid {
		go func() { r.stop <- true }()
	}
	return !valid
}

type Cache struct {
	recs map[string]*record
	ttl  time.Duration
	l    sync.RWMutex
}

func New(ttl time.Duration) *Cache {
	return &Cache{
		recs: make(map[string]*record),
		ttl:  ttl,
	}
}

func (c *Cache) LookupHost(host string) (addr string, err error) {
	cached, expired, ok := c.cachedAddr(host)
	if ok && !expired {
		cacheHit.Mark(1)
		return cached, nil
	}

	// Value is not cached, look it up
	hosts, err := net.LookupHost(host)

	// In the case there was an error looking up the value AND we have a cached
	// value (that has expired), just return the cached value.
	if err != nil && ok {
		lookupRecover.Mark(1)
		return cached, nil
	}

	if err != nil {
		lookupErr.Mark(1)
		return "", err
	}

	rec := Record(c.ttl, hosts)
	c.l.Lock()
	c.recs[host] = rec
	c.l.Unlock()
	cacheMiss.Mark(1)
	return rec.Addr(), nil
}

func (c *Cache) cachedAddr(host string) (addr string, expired bool, ok bool) {
	c.l.RLock()
	rec, ok := c.recs[host]
	c.l.RUnlock()

	if ok {
		return rec.Addr(), rec.Expired(), true
	}

	return "", true, false
}
