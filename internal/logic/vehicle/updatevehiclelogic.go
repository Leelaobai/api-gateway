// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package vehicle

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVehicleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改座驾昵称
func NewUpdateVehicleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVehicleLogic {
	return &UpdateVehicleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateVehicleLogic) UpdateVehicle(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
