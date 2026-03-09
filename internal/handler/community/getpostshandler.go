package community

import (
	"net/http"

	"api-gateway/internal/util"
	"api-gateway/internal/proxy"
	"api-gateway/internal/svc"
)

func GetPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.Forward(svcCtx.CommunityProxy, w, r, util.ExtractUserID(r))
	}
}
