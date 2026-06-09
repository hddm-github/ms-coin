package logic

import (
	"context"
	"errors"
	"grpc-common/ucenter/types/login"
	"mscoin-common/tools"
	"time"
	"ucenter/internal/domain"

	"ucenter/internal/svc"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

const LoginCacheKey = "REGISTER:"

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MemberDomain *domain.MemberDomain
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:          ctx,
		svcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
		MemberDomain: domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *LoginLogic) Login(req *login.LoginReq) (*login.LoginRes, error) {

	// 校验密码
	member, err := l.MemberDomain.FindByPhone(context.Background(), req.GetUsername())
	if err != nil {
		logx.Error(err)
		return nil, errors.New("登录失败")
	}
	if member == nil {
		return nil, errors.New("此用户未注册")
	}
	password := member.Password
	salt := member.Salt
	verify := tools.Verify(req.Password, salt, password, nil)
	if !verify {
		return nil, errors.New("密码不正确")
	}
	// 登录成功，生成 token
	// jwt
	key := l.svcCtx.Config.JWT.AccessSecret
	expire := l.svcCtx.Config.JWT.AccessExpire

	token, err := l.getJwtToken(key, time.Now().Unix(), expire, member.Id)
	if err != nil {
		return nil, errors.New("token生成错误")
	}
	// 返回登录信息
	loginCount := member.LoginCount + 1
	go func() {
		l.MemberDomain.UpdateLoginCount(context.Background(), member.Id, 1)
	}()
	return &login.LoginRes{
		Token:         token,
		Id:            member.Id,
		Username:      member.Username,
		MemberLevel:   member.MemberLevelStr(),
		MemberRate:    member.MemberRate(),
		RealName:      member.RealName,
		Country:       member.Country,
		Avatar:        member.Avatar,
		PromotionCode: member.PromotionCode,
		SuperPartner:  member.SuperPartner,
		LoginCount:    int32(loginCount),
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
