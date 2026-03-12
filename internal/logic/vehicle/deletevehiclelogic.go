// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package vehicle

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVehicleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除座驾
func NewDeleteVehicleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVehicleLogic {
	return &DeleteVehicleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteVehicleLogic) DeleteVehicle(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
