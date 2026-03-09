package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var transport = &http.Transport{
	MaxIdleConns:        200,
	MaxIdleConnsPerHost: 50,
	IdleConnTimeout:     90 * time.Second,
}

func NewProxy(targetBase string) *httputil.ReverseProxy {
	target, _ := url.Parse(targetBase)
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host
		},
		Transport: transport,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(`{"code":"SYSTEM_UNAVAILABLE","message":"服务暂时不可用，请稍后重试"}`))
		},
	}
}

func Forward(p *httputil.ReverseProxy, w http.ResponseWriter, r *http.Request, userID string) {
	if userID != "" {
		r.Header.Set("X-User-Id", userID)
	}
	p.ServeHTTP(w, r)
}
