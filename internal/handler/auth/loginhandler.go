package auth

import (
	"net/http"

	"api-gateway/internal/proxy"
	"api-gateway/internal/svc"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !svcCtx.AuthRateLimiter.Check(r) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"code":"RATE_LIMIT_EXCEEDED","message":"操作过于频繁，请稍后再试"}`))
			return
		}
		proxy.Forward(svcCtx.UserProxy, w, r, "")
	}
}
