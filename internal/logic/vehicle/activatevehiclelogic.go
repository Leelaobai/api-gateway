// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package vehicle

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivateVehicleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 切换当前座驾
func NewActivateVehicleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivateVehicleLogic {
	return &ActivateVehicleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivateVehicleLogic) ActivateVehicle(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
