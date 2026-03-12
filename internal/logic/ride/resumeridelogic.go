// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package ride

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResumeRideLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 继续骑行
func NewResumeRideLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResumeRideLogic {
	return &ResumeRideLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResumeRideLogic) ResumeRide(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
