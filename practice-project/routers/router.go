package routers

import (
	
	"practice-project/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.CatController{}, "get:ShowCat", )
	beego.Router("/breeds", &controllers.CatController{}, "get:GetBreedsHandler")
	beego.Router("/breed-images", &controllers.CatController{}, "get:GetBreedImagesHandler")
	beego.Router("/vote", &controllers.CatController{}, "post:VoteOnImage")

	// beego.Router("/breeds", &controllers.CatController{}, "get:ShowBreed")
}
