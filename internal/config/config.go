package config

import "github.com/zeromicro/go-zero/rest"

type RateLimitConfig struct {
	WindowSec   int
	GlobalPerIP int
	AuthPerIP   int
}

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	UserServiceUrl      string
	RideServiceUrl      string
	CommunityServiceUrl string
	RedisAddr           string
	RateLimit           RateLimitConfig
}
