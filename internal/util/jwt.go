package util

import (
	"fmt"
	"net/http"
)

func ExtractUserID(r *http.Request) string {
	val := r.Context().Value("user_id")
	if val == nil {
		return ""
	}
	return fmt.Sprintf("%v", val)
}
