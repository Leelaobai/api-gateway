package auth

import (
	"net/http"

	"api-gateway/internal/proxy"
	"api-gateway/internal/svc"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !svcCtx.AuthRateLimiter.Check(r) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error":"too many requests"}`))
			return
		}
		proxy.Forward(svcCtx.UserProxy, w, r, "")
	}
}
