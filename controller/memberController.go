package controller

import (
	"svenBlog/dao"
	"svenBlog/middlewares"
	"svenBlog/model"
	"svenBlog/service"
	"svenBlog/tool/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func Register(ctx *gin.Context) {
	var member dao.MemberDao
	if err := ctx.ShouldBind(&member); err == nil {
		username := member.MemDao.Username
		password := member.MemDao.Password
		// 判断用户名长度
		if strings.Count(username, "")-1 >= 15 || strings.Count(username, "")-1 <= 6 {
			response.Response(ctx, http.StatusUnprocessableEntity, 200, nil, "用户名必须在6-15个字符长度")
			// 判断密码长度
		} else if strings.Count(password, "")-1 >= 20 || strings.Count(password, "")-1 <= 10 {
			response.Response(ctx, http.StatusUnprocessableEntity, 200, nil, "密码必须在6-15个字符长度")
		} else {
			// 用户注册
			if member.IsUserExist(username) {
				// 密码加密 (输入的密码)
				hashedPassword, err := passwordEncryption(password)
				if err != nil {
					response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
				}
				// 新建用户对象
				newMember := model.Member{
					Username: member.MemDao.Username,
					Password: string(hashedPassword),
				}
				if err := member.CreateTable(&newMember); err == nil {
					ctx.JSON(http.StatusOK, gin.H{
						"code":          200,
						"username":      member.MemDao.Username,
						"password":      password,
						"hasedPassword": hashedPassword,
					})
				} else {
					response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户注册错误")
				}
			} else {
				response.Response(ctx, http.StatusUnprocessableEntity, 200, nil, "用户名已存在")
			}
		}
	} else {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "服务器内部错误")
	}
}

// 用户登录
func Login(ctx *gin.Context) {
	var member service.MemberService
	if err := ctx.ShouldBind(&member); err == nil {
		username := member.MbService.MemDao.Username
		password := member.MbService.MemDao.Password
		loginData := member.MemberLogin(username)
		if loginData != nil {
			// 判断密码收否正确
			if err := passwordDecode(loginData.Password, password); err == nil {
				token, err := middlewares.GenToken(username)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "失败"})
				} else {
					ctx.JSON(http.StatusOK, gin.H{
						"code": 200,
						"user": username,
						"token": token,
					})
				}
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
			}
		}
	} else {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "服务器内部错误")
	}
}

// 用户修改密码
func CheckPassword(ctx *gin.Context) {
	var member dao.MemberDao
	if err := ctx.ShouldBind(&member); err == nil {
		encryption, err := passwordEncryption(member.MemDao.Password)
		if err != nil {
			response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		}
		if err := member.UpdateMember(string(encryption)); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"user": member.MemDao.Username,
				"pass": string(encryption),
				"msg": "密码更新完成",
			})
		}
	} else {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "服务器内部错误")
	}
}

// 密码加密
func passwordEncryption(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, err

}

// 密码解密
func passwordDecode(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}


