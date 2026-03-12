// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package vehicle

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMyVehiclesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 我的座驾列表
func NewListMyVehiclesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMyVehiclesLogic {
	return &ListMyVehiclesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMyVehiclesLogic) ListMyVehicles(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
