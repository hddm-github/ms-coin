// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"grpc-common/ucenter/types/register"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.Request) (resp *types.Response, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	if req.Captcha == nil {
		req.Captcha = &types.CaptchaReq{}
	}

	regReq := &register.RegReq{}
	if err := copier.Copy(regReq, req); err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UCRegisterRpc.RegisterByPhone(ctx, regReq)
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Message: "注册成功",
	}, nil
}

func (l *RegisterLogic) SendCode(t *types.CodeRequest) (resp *types.CodeResponse, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	l.svcCtx.UCRegisterRpc.SendCode(ctx, &register.CodeReq{
		Phone:   t.Phone,
		Country: t.Country,
	})
	return
}
