package routers

import (
	
	"practice-project/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    // beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.CatController{}, "get:ShowCat")
	beego.Router("/getcatdata", &controllers.CatController{}, "get:GetCatData")
	beego.Router("/breeds", &controllers.CatController{}, "get:GetBreedsHandler")
	beego.Router("/breed-images", &controllers.CatController{}, "get:GetBreedImagesHandler")
	beego.Router("/breed-images/:breedID", &controllers.CatController{}, "get:FetchImagesByBreedHandler") 
	beego.Router("/vote/up", &controllers.CatController{}, "post:VoteUp")
	beego.Router("/vote/down", &controllers.CatController{}, "post:VoteDown")
	beego.Router("/vote/history", &controllers.CatController{}, "get:VoteHistory")
	// beego.Router("/vote/history", &controllers.CatController{}, "delete:VoteHistory")

	// beego.Router("/vote", &controllers.CatController{}, "post:VoteOnImage")
    // beego.Router("/breeds", &controllers.CatController{}, "get:ShowBreed")
}
