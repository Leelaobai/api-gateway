// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package ride

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartRideLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 开始骑行
func NewStartRideLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartRideLogic {
	return &StartRideLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartRideLogic) StartRide(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
