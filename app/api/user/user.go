package user

import (
	"MyGoTest/app/service/user"
	"MyGoTest/library/response"

	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/util/gvalid"
)

//Controller 用户API管理对象
type Controller struct{}

//SignUp 用户注册接口
func (c *Controller) SignUp(r *ghttp.Request) {
	if err := user.SignUp(r.GetPostMap()); err != nil {
		response.JSON(r, 1, err.Error())
	} else {
		response.JSON(r, 0, "ok")
	}
}

//SignIn 用户登录接口
func (c *Controller) SignIn(r *ghttp.Request) {
	data := r.GetPostMap()
	rules := map[string]string{
		"passport": "required",
		"password": "required",
	}
	msgs := map[string]interface{}{
		"passport": "账号不能为空",
		"password": "密码不能为空",
	}
	if e := gvalid.CheckMap(data, rules, msgs); e != nil {
		response.JSON(r, 1, e.String())
	}
	if err := user.SignIn(data["passport"], data["password"], r.Session); err != nil {
		response.JSON(r, 1, err.Error())
	} else {
		response.JSON(r, 0, "ok")
	}
}

//IsSignedIn 判断用户是否已经登录
func (c *Controller) IsSignedIn(r *ghttp.Request) {
	if user.IsSignedIn(r.Session) {
		response.JSON(r, 0, "ok")
	} else {
		response.JSON(r, 1, "")
	}
}

//SignOut 用户注销/退出接口
func (c *Controller) SignOut(r *ghttp.Request) {
	user.SignOut(r.Session)
	response.JSON(r, 0, "ok")
}

//CheckPassport 检测用户账号接口(唯一性校验)
func (c *Controller) CheckPassport(r *ghttp.Request) {
	passport := r.Get("passport")
	if e := gvalid.Check(passport, "required", "请输入账号"); e != nil {
		response.JSON(r, 1, e.String())
	}
	if user.CheckPassport(passport) {
		response.JSON(r, 0, "ok")
	}
	response.JSON(r, 1, "账号已经存在")
}

//CheckNickName 检测用户昵称接口(唯一性校验)
func (c *Controller) CheckNickName(r *ghttp.Request) {
	nickname := r.Get("nickname")
	if e := gvalid.Check(nickname, "required", "请输入昵称"); e != nil {
		response.JSON(r, 1, e.String())
	}
	if user.CheckNickName(r.Get("nickname")) {
		response.JSON(r, 0, "ok")
	}
	response.JSON(r, 1, "昵称已经存在")
}
