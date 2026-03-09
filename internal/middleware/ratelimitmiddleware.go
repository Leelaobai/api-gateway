package middleware

import (
	"net"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RateLimitMiddleware struct {
	globalLimiter *limit.PeriodLimit
}

func NewRateLimitMiddleware(rds *redis.Redis, windowSec, globalPerIP int) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		globalLimiter: limit.NewPeriodLimit(windowSec, globalPerIP, rds, "rate:ip:"),
	}
}

func (m *RateLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := clientIP(r)
		code, _ := m.globalLimiter.Take(ip)
		if code == limit.OverQuota || code == limit.HitQuota {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"code":"RATE_LIMIT_EXCEEDED","message":"操作过于频繁，请稍后再试"}`))
			return
		}
		next(w, r)
	}
}

func clientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	if xri := r.Header.Get("X-Real-Ip"); xri != "" {
		return xri
	}
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}
