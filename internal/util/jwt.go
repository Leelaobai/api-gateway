package util

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
)

func ExtractUserID(r *http.Request) string {
	val := r.Context().Value("user_id")
	if val == nil {
		return ""
	}
	return fmt.Sprintf("%v", val)
}

// ExtractUserIDOptionalStrict 用于“未登录可访问”的接口：
// - 若请求未携带 Authorization，则返回 ("", "")，表示按未登录处理
// - 若携带 Authorization，则必须解析成功；失败返回 ("", <errorCode>)
// - 解析成功返回 (userID, "")
//
// errorCode 取值：
// - AUTH_TOKEN_EXPIRED
// - AUTH_TOKEN_INVALID
func ExtractUserIDOptionalStrict(r *http.Request, secret string) (userID string, errorCode string) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", ""
	}
	if !strings.HasPrefix(auth, "Bearer ") {
		return "", "AUTH_TOKEN_INVALID"
	}
	tokenStr := strings.TrimPrefix(auth, "Bearer ")
	if tokenStr == "" {
		return "", "AUTH_TOKEN_INVALID"
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return "", "AUTH_TOKEN_EXPIRED"
			}
		}
		return "", "AUTH_TOKEN_INVALID"
	}
	if token == nil || !token.Valid {
		return "", "AUTH_TOKEN_INVALID"
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "AUTH_TOKEN_INVALID"
	}
	raw, ok := claims["user_id"]
	if !ok {
		return "", "AUTH_TOKEN_INVALID"
	}
	uid := fmt.Sprintf("%v", raw)
	if uid == "" || uid == "<nil>" {
		return "", "AUTH_TOKEN_INVALID"
	}
	return uid, ""
}
