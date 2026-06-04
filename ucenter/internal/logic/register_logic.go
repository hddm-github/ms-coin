package logic

import (
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"mscoin-common/tools"
	"time"
	"ucenter/internal/domain"

	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "REGISTER:"

type RegisterByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MemberDomain *domain.MemberDomain
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByPhoneLogic {
	return &RegisterByPhoneLogic{
		ctx:          ctx,
		svcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
		MemberDomain: domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *RegisterByPhoneLogic) RegisterByPhone(req *register.RegReq) (*register.RegRes, error) {

	// 校验验证码
	redisValue := ""
	err := l.svcCtx.Cache.GetCtx(context.Background(), RegisterCacheKey+req.Phone, &redisValue)
	if err != nil {
		return nil, errors.New("验证码获取错误")
	}
	if req.Code != redisValue {
		return nil, errors.New("验证码错误") //
	}
	// 验证码通过 进行注册 手机号受限验证是否注册过
	mem, err := l.MemberDomain.FindByPhone(context.Background(), req.Phone)
	if err != nil {
		return nil, errors.New("服务异常，请联系管理员")
	}
	if mem != nil {
		return nil, errors.New("手机号已注册")
	}
	// 4. 生成 member 模型，存入数据库
	err = l.MemberDomain.Register(context.Background(),
		req.Phone,
		req.Password,
		req.Country,
		req.Username,
		req.SuperPartner,
		req.Promotion)

	if err != nil {
		return nil, errors.New("注册失败，请联系管理员")
	}
	return &register.RegRes{}, nil
}

func (l *RegisterByPhoneLogic) SendCode(req *register.CodeReq) (*register.NoRes, error) {
	code := tools.Rand4Num()
	go func() {
		logx.Infof("调用短信验证码接口，手机号：%s，验证码：%s", req.Phone, code)
	}()
	logx.Infof("验证码：%s \n", code)
	ctx, cancel := context.WithTimeout(l.ctx, 60*time.Second)
	defer cancel()
	err := l.svcCtx.Cache.SetWithExpireCtx(ctx, RegisterCacheKey+req.Phone, code, 60*time.Second)
	if err != nil {
		return nil, errors.New("验证码存入 Cache 失败")
	}
	return &register.NoRes{}, nil
}
