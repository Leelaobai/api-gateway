// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package vehicle

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyVehicleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 我的座驾
func NewGetMyVehicleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyVehicleLogic {
	return &GetMyVehicleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyVehicleLogic) GetMyVehicle(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
