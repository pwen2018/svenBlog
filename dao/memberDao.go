package dao

import (
	"svenBlog/model"
	"svenBlog/tool"
)

type MemberDao struct {
	MemDao model.Member
}

// 创建用户
func (md *MemberDao) CreateTable(member *model.Member) error {
	err := tool.DB.Table("member").Create(&member).Error
	if err != nil {
		return err
	}
	return nil
}

// 判断用户是否存在
func (mb *MemberDao) IsUserExist(username string) bool {
	tool.DB.Table("member").Where("username = ?", username).First(&mb.MemDao)
	if mb.MemDao.Id == 0 {
		return true
	}
	return false
}

// 查询用户
func (mb *MemberDao) SelectMember(username string) error {
	err := tool.DB.Table("member").Where("username = ?", username).First(&mb.MemDao).Error
	if err != nil {
		return err
	}
	return nil
}

// 更改密码
func (mb *MemberDao) UpdateMember(password string) error {
	err := tool.DB.Table("member").Update(model.Member{Password: password}).Error
	if err != nil {
		return err
	}
	return nil
}
