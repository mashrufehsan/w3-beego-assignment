package routers

import (
	"w3-beego-assignment/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// HTTP Routes
	beego.Router("/", &controllers.MainController{})
	beego.Router("/tab_breeds", &controllers.ShowBreedsController{})
	beego.Router("/tab_favs", &controllers.ShowFavsController{})
	beego.Router("/tab_myVotes", &controllers.ShowMyVotesController{})

	// API Routes
	beego.Router("/breeds", &controllers.BreedController{}, "get:GetBreeds")
	beego.Router("/breeds/:breed_id", &controllers.BreedController{}, "get:GetBreedsByID")
	beego.Router("/breeds/images/:breed_id", &controllers.BreedController{}, "get:GetImagesByBreed")
	beego.Router("/getACat", &controllers.CatController{}, "get:GetACat")
	beego.Router("/createAFavourite", &controllers.CatController{}, "post:CreateAFavourite")
	beego.Router("/getFavourites", &controllers.CatController{}, "get:GetFavourites")
	beego.Router("/deleteAFavourite/:favourite_id", &controllers.CatController{}, "delete:DeleteAFavourite")
	beego.Router("/getVotes", &controllers.VoteController{}, "get:GetVotes")
	beego.Router("/vote", &controllers.VoteController{}, "post:Vote")
}
