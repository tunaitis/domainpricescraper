package godaddy

import (
	"net/http"
	"net/url"
)

type Option func(*GoDaddy)

func WithProxy(proxy *url.URL) Option {
	return func(g *GoDaddy) {
		g.transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}
}

func WithApi(key string, secret string) Option {
	return func(g *GoDaddy) {
		g.apiKey = key
		g.apiSecret = secret
	}
}
