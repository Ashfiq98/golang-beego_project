package main

import (
	_ "practice-project/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

