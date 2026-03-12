package community

import (
	"net/http"

	"api-gateway/internal/proxy"
	"api-gateway/internal/svc"
	"api-gateway/internal/util"
)

func GetPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, code := util.ExtractUserIDOptionalStrict(r, svcCtx.Config.Auth.AccessSecret)
		if code != "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			if code == "AUTH_TOKEN_EXPIRED" {
				w.Write([]byte(`{"code":"AUTH_TOKEN_EXPIRED","message":"登录已过期，请重新登录"}`))
			} else {
				w.Write([]byte(`{"code":"AUTH_TOKEN_INVALID","message":"认证失效，请重新登录"}`))
			}
			return
		}
		proxy.Forward(svcCtx.CommunityProxy, w, r, userID)
	}
}
