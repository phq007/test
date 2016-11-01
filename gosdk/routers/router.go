// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gosdk/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/device/:deviceid", &controllers.DeviceInfoController{}, "Get:Get")
	beego.Router("/device/bind", &controllers.DeviceInfoController{}, "*:Bind")
	beego.Router("/device/checkbind/:deviceid", &controllers.DeviceInfoController{}, "Get:CheckBind")
	beego.Router("/device/start/:deviceid", &controllers.DeviceInfoController{}, "Get:Start")
}
