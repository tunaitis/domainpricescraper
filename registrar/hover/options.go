package hover

import (
	"net/http"
	"net/url"
)

type Option func(*Hover)

func WithProxy(proxy *url.URL) Option {
	return func(h *Hover) {
		h.transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}
}
