// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package community

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRouteBooksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 路书列表
func NewGetRouteBooksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRouteBooksLogic {
	return &GetRouteBooksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRouteBooksLogic) GetRouteBooks(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
