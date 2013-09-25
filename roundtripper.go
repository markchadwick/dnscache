package dnscache

import (
	"net"
	"net/http"
	"time"
)

func DnsCachedRoundTripper(ttl time.Duration) http.RoundTripper {
	return RoundTripper(New(ttl))
}

func RoundTripper(cache *Cache) http.RoundTripper {
	transport := &http.Transport{Proxy: http.ProxyFromEnvironment}
	transport.Dial = func(network, addr string) (net.Conn, error) {
		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}
		resolved, err := cache.LookupHost(host)
		if err != nil {
			return nil, err
		}
		addr = net.JoinHostPort(resolved, port)
		return net.Dial(network, addr)
	}
	return transport
}
