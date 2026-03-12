package vehicle

import (
	"net/http"

	"api-gateway/internal/proxy"
	"api-gateway/internal/svc"
	"api-gateway/internal/util"
)

func DeleteVehicleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.Forward(svcCtx.UserProxy, w, r, util.ExtractUserID(r))
	}
}
