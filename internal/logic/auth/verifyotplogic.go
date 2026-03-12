// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"context"

	"api-gateway/internal/svc"
	"api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyOTPLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证码登录/注册
func NewVerifyOTPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyOTPLogic {
	return &VerifyOTPLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyOTPLogic) VerifyOTP(req *types.ProxyReq) (resp *types.ProxyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
