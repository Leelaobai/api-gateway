// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package ride

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PauseRideLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 暂停骑行
func NewPauseRideLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PauseRideLogic {
	return &PauseRideLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PauseRideLogic) PauseRide(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
