package service

import (
	"svenBlog/dao"
	"svenBlog/model"
)

type MemberService struct {
	MbService *dao.MemberDao
}

// 用户登录
func (mbService *MemberService) MemberLogin(username string) *model.Member {
	var member dao.MemberDao
	if err := member.SelectMember(username); err == nil {
		if member.MemDao.Id == 0 {
			return nil
		} else {
			return &member.MemDao
		}
	}
	return nil
}
