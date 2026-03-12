// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package ride

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRideLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 骑行详情
func NewGetRideLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRideLogic {
	return &GetRideLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRideLogic) GetRide(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
