package ride

import (
	"net/http"

	"api-gateway/internal/util"
	"api-gateway/internal/proxy"
	"api-gateway/internal/svc"
)

func ResumeRideHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.Forward(svcCtx.RideProxy, w, r, util.ExtractUserID(r))
	}
}
