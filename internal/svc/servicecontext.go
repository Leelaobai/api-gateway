package svc

import (
	"net/http/httputil"

	"api-gateway/internal/config"
	"api-gateway/internal/middleware"
	"api-gateway/internal/proxy"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	RateLimit       rest.Middleware
	AuthRateLimiter *middleware.AuthRateLimiter
	UserProxy       *httputil.ReverseProxy
	RideProxy       *httputil.ReverseProxy
	CommunityProxy  *httputil.ReverseProxy
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.RedisAddr,
		Type: "node",
	})

	rl := c.RateLimit
	return &ServiceContext{
		Config:          c,
		RateLimit:       middleware.NewRateLimitMiddleware(rds, rl.WindowSec, rl.GlobalPerIP).Handle,
		AuthRateLimiter: middleware.NewAuthRateLimiter(rds, rl.WindowSec, rl.AuthPerIP),
		UserProxy:       proxy.NewProxy(c.UserServiceUrl),
		RideProxy:       proxy.NewProxy(c.RideServiceUrl),
		CommunityProxy:  proxy.NewProxy(c.CommunityServiceUrl),
	}
}
