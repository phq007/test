package routers

import (
	"gosdk/controllers/ad"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ad", &controllers.ADController{})
	beego.Router("/res/:hash", &controllers.ResInfoController{}, "Get:Get")
	beego.Router("/playlist", &controllers.PlayListController{}, "Get:Get")
	beego.Router("/tmpl/:hash", &controllers.PlayListController{}, "Get:GetTmplData")

}
