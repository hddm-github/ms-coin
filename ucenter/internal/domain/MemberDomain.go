package domain

import (
	"context"
	"errors"
	"mscoin-common/msdb"
	"mscoin-common/tools"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberDomain struct {
	MemberRepo repo.MemberRepo
}

func (d *MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	//涉及到数据库查询
	mem, err := d.MemberRepo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func (d *MemberDomain) Register(
	ctx context.Context,
	phone string,
	password string,
	country string,
	username string,
	partner string,
	promotion string) error {

	mem := model.NewMember()
	_ = tools.Default(mem)
	salt, pwd := tools.Encode(password, nil)
	mem.Username = username
	mem.Password = pwd
	mem.Salt = salt
	mem.MobilePhone = phone
	mem.FillSuperPartner(partner)
	mem.Country = country
	mem.MemberLevel = model.GENERAL
	mem.PromotionCode = promotion
	err := d.MemberRepo.Save(ctx, mem)
	if err != nil {
		logx.Error(ctx, "注册失败, error=%v", err)
		return errors.New("数据库异常")
	}

	return nil
}
func NewMemberDomain(db *msdb.MsDB) *MemberDomain {
	return &MemberDomain{
		MemberRepo: dao.NewMemberRepo(db),
	}
}
