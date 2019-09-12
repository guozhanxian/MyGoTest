package user

import (
	"errors"
	"fmt"

	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/util/gvalid"
)

// UserSessionMark
const (
	UserSessionMark = "user_info"
)

var (
	// 表对象
	table = g.DB().Table("user").Safe()
)

//SignUp 用户注册
func SignUp(data g.MapStrStr) error {
	// 数据校验
	rules := []string{
		"passport @required|length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password2@required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间",
		"password @required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
	}
	if e := gvalid.CheckMap(data, rules); e != nil {
		return errors.New(e.String())
	}
	if _, ok := data["nickname"]; !ok {
		data["nickname"] = data["passport"]
	}
	// 唯一性数据检查
	if !CheckPassport(data["passport"]) {
		return fmt.Errorf(fmt.Sprintf("账号 %s 已经存在", data["passport"]))
	}
	if !CheckNickName(data["nickname"]) {
		return fmt.Errorf(fmt.Sprintf("昵称 %s 已经存在", data["nickname"]))
	}
	// 记录账号创建/注册时间
	if _, ok := data["create_time"]; !ok {
		data["create_time"] = gtime.Now().String()
	}
	if _, err := table.Filter().Data(data).Save(); err != nil {
		return err
	}
	return nil
}

//IsSignedIn 判断用户是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains(UserSessionMark)
}

//SignIn 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func SignIn(passport, password string, session *ghttp.Session) error {
	record, err := table.Where("passport=? and password=?", passport, password).One()
	if err != nil {
		return err
	}
	if record == nil {
		return errors.New("账号或密码错误")
	}
	session.Set(UserSessionMark, record)
	return nil
}

//SignOut 用户注销
func SignOut(session *ghttp.Session) {
	session.Remove(UserSessionMark)
}

//CheckPassport 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckPassport(passport string) bool {
	i, err := table.Where("passport", passport).Count()
	if err != nil {
		return false
	}
	return i == 0
}

//CheckNickName 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckNickName(nickname string) bool {
	i, err := table.Where("nickname", nickname).Count()
	if err != nil {
		return false
	}
	return i == 0
}
