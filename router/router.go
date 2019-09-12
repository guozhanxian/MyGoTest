package router

import (
	"MyGoTest/app/api/user"

	"github.com/gogf/gf/g"
)

func init() {
	// 用户模块 路由注册 - 使用执行对象注册方式
	g.Server().BindObject("/user", new(user.Controller))
}
