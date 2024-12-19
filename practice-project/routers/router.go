package routers

import (
	"practice-project/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.CatController{}, "get:ShowCat")
}
