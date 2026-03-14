package community

import (
	"net/http"

	"api-gateway/internal/proxy"
	"api-gateway/internal/svc"
)

// GetOssSignedUrlHandler 转发到 community-service，获取图片 presigned URL 并 302 重定向（公开，无需登录）
func GetOssSignedUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.Forward(svcCtx.CommunityProxy, w, r, "")
	}
}
