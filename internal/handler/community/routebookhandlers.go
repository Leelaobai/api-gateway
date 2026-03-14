package community

import (
	"net/http"

	"api-gateway/internal/proxy"
	"api-gateway/internal/svc"
	"api-gateway/internal/util"
)

// requireUserID 提取必须存在的 userID，否则返回 401
func requireUserID(w http.ResponseWriter, r *http.Request, secret string) (string, bool) {
	userID, code := util.ExtractUserIDOptionalStrict(r, secret)
	if code != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		if code == "AUTH_TOKEN_EXPIRED" {
			w.Write([]byte(`{"code":"AUTH_TOKEN_EXPIRED","message":"登录已过期，请重新登录"}`))
		} else {
			w.Write([]byte(`{"code":"AUTH_TOKEN_INVALID","message":"认证失效，请重新登录"}`))
		}
		return "", false
	}
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"code":"AUTH_TOKEN_INVALID","message":"请先登录"}`))
		return "", false
	}
	return userID, true
}

// GetPostDetailHandler 帖子详情（可选 JWT，用于带出 is_liked）
func GetPostDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return GetRouteBookHandler(svcCtx)
}

// GetRouteBookHandler 路书详情（可选 JWT）
func GetRouteBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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

// GetCommentsHandler 评论列表（可选 JWT）
func GetCommentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return GetRouteBookHandler(svcCtx)
}

// 用户主页（可选 JWT，用于带出 is_following 等状态）
func GetUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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

// GetUserPostsHandler 某用户的动态列表（可选 JWT，本人可看私密动态）
func GetUserPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return GetUserProfileHandler(svcCtx)
}

// authHandler 通用需要 JWT 的 handler 工厂
func authHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := requireUserID(w, r, svcCtx.Config.Auth.AccessSecret)
		if !ok {
			return
		}
		proxy.Forward(svcCtx.CommunityProxy, w, r, userID)
	}
}

func CreateRouteBookFromRideHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func UpdateRouteBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func DeleteRouteBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func LikeRouteBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func UnlikeRouteBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func FavoriteRouteBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func UnfavoriteRouteBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func PostCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func DeleteCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func MyRouteBooksHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func MyFavoritesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

// ── 以下为社区帖子 / 关注相关受保护接口 ──

func CreatePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func DeletePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func LikePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func UnlikePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func AddPostCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func DeletePostCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func FollowUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func UnfollowUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}

func OssPresignHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return authHandler(svcCtx)
}
