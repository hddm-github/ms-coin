package logic

import (
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"mscoin-common/tools"
	"time"

	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "REGISTER:"

type RegisterByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByPhoneLogic {
	return &RegisterByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterByPhoneLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	// todo: add your logic here and delete this line
	logx.Info("ucenter register by phone, in: %v", in)
	return &register.RegRes{}, nil
}

func (l *RegisterByPhoneLogic) SendCode(req *register.CodeReq) (*register.NoRes, error) {
	code := tools.Gen4Number()
	go func() {
		logx.Infof("调用短信验证码接口，手机号：%s，验证码：%s", req.Phone, code)
	}()
	logx.Infof("验证码：%s \n", code)
	ctx, cancel := context.WithTimeout(l.ctx, 15*time.Second)
	defer cancel()
	err := l.svcCtx.Cache.SetWithExpireCtx(ctx, RegisterCacheKey+req.Phone, code, 15*time.Second)
	if err != nil {
		return nil, errors.New("验证码存入 Cache 失败")
	}
	return &register.NoRes{}, nil
}
