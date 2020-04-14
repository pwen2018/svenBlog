package model

type Member struct {
	Id       int64  `form:"id"`
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}
