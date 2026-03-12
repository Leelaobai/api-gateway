// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package ride

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRidesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 骑行记录列表
func NewListRidesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRidesLogic {
	return &ListRidesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRidesLogic) ListRides(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
