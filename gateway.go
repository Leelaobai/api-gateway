package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"api-gateway/internal/config"
	"api-gateway/internal/handler"
	"api-gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gateway-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		c.Auth.AccessSecret = secret
	}
	if addr := os.Getenv("REDIS_ADDR"); addr != "" {
		c.RedisAddr = addr
	}
	if url := os.Getenv("USER_SERVICE_URL"); url != "" {
		c.UserServiceUrl = url
	}
	if url := os.Getenv("RIDE_SERVICE_URL"); url != "" {
		c.RideServiceUrl = url
	}
	if url := os.Getenv("COMMUNITY_SERVICE_URL"); url != "" {
		c.CommunityServiceUrl = url
	}
	if v := os.Getenv("RATE_LIMIT_WINDOW_SEC"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			c.RateLimit.WindowSec = n
		}
	}
	if v := os.Getenv("RATE_LIMIT_GLOBAL_PER_IP"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			c.RateLimit.GlobalPerIP = n
		}
	}
	if v := os.Getenv("RATE_LIMIT_AUTH_PER_IP"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			c.RateLimit.AuthPerIP = n
		}
	}

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting api-gateway at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
