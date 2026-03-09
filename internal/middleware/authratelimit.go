package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type AuthRateLimiter struct {
	limiter *limit.PeriodLimit
}

func NewAuthRateLimiter(rds *redis.Redis, windowSec, authPerIP int) *AuthRateLimiter {
	return &AuthRateLimiter{
		limiter: limit.NewPeriodLimit(windowSec, authPerIP, rds, "rate:auth:"),
	}
}

func (l *AuthRateLimiter) Check(r *http.Request) bool {
	ip := clientIP(r)
	code, _ := l.limiter.Take(ip)
	return code != limit.OverQuota && code != limit.HitQuota
}
