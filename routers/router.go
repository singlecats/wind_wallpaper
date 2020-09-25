package routers

import (
	"wind_wallpaper/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/list", &controllers.ListController{}, "get:Index")
}
