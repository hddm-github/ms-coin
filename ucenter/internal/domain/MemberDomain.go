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
	memberRepo repo.MemberRepo
}

func (d *MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	//涉及到数据库查询
	mem, err := d.memberRepo.FindByPhone(ctx, phone)
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
	mem.Avatar = "https://mszlu.oss-cn-beijing.aliyuncs.com/mscoin/defaultavatar.png"
	mem.MobilePhone = phone
	mem.FillSuperPartner(partner)
	mem.Country = country
	mem.MemberLevel = model.GENERAL
	mem.PromotionCode = promotion
	err := d.memberRepo.Save(ctx, mem)
	if err != nil {
		logx.Error(ctx, "注册失败, error=%v", err)
		return errors.New("数据库异常")
	}

	return nil
}

func (d *MemberDomain) UpdateLoginCount(background context.Context, id int64, i int) {
	err := d.memberRepo.UpdateLoginCount(background, id, i)
	if err != nil {
		logx.Error(err)
	}
}
func NewMemberDomain(db *msdb.MsDB) *MemberDomain {
	return &MemberDomain{
		memberRepo: dao.NewMemberRepo(db),
	}
}
