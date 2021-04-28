package routers

import (
	"fyoukuApi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.VideoController{})
}
