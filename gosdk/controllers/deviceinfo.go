package controllers

import (
	"encoding/base64"
	"encoding/json"
	"gosdk/models"
	"gosdk/sdk"
	"gosdk/sysconfig"
	"gosdk/util"
	"log"
	"strings"

	"github.com/astaxie/beego"
)

type DeviceInfoController struct {
	beego.Controller
}

func (c *DeviceInfoController) Get() {
	deviceid := c.Ctx.Input.Param(":deviceid")
	devicemanager := &models.DeviceInfoManager{}
	device, _ := devicemanager.GetDeviceInfo(deviceid)
	sysconfig.Device = *device
	c.Data["json"] = device
	c.ServeJSON()
}

func (c *DeviceInfoController) Start() {
	if sysconfig.Device.DeviceToken == "" {
		deviceid := c.Ctx.Input.Param(":deviceid")
		devicemanager := &models.DeviceInfoManager{}
		device, _ := devicemanager.GetDeviceInfo(deviceid)
		sysconfig.Device = *device
	}
	sdk.StartBusiness()
}

func (c *DeviceInfoController) Bind() {
	//deviceid := c.GetString("deviceid")
	//token := c.GetString("token")
}

func (c *DeviceInfoController) CheckBind() {
	deviceid := c.Ctx.Input.Param(":deviceid")
	token := c.GetString("token")
	devicemanager := &models.DeviceInfoManager{}
	device, _ := devicemanager.GetDeviceInfo(deviceid)
	if device.LoginId == "" || device.Password == "" {
		devicejson := util.HTTPGet(sysconfig.SystemURL + "/ad_api/checkbind?deviceid=" + deviceid + "&bind_device_token=" + token)
		json.Unmarshal([]byte(devicejson), &device)
		if device.BindDeviceToken != "" {
			b, _ := base64.StdEncoding.DecodeString(device.BindDeviceToken)
			s := strings.SplitN(string(b), "|", 2)
			if len(s) == 2 {
				device.LoginId = s[0]
				device.Password = s[1]
				devicemanager.InsertDeviceInfo(*device)
				log.Println("insert device info:" + device.Id)
			}
		}
	}
	sysconfig.Device = *device
	c.Data["json"] = device
	c.ServeJSON()
}
