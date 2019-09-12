package main

import (
	_ "MyGoTest/boot"
	_ "MyGoTest/router"

	"github.com/gogf/gf/g"
)

func main() {
	g.Server().Run()
}
