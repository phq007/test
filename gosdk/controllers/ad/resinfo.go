package controllers

import (
	"gosdk/models"

	"github.com/astaxie/beego"
)

type ResInfoController struct {
	beego.Controller
}

func (c *ResInfoController) Get() {
	id := c.Ctx.Input.Param(":hash")
	manager := &models.ResInfoManager{}
	info, _ := manager.Get(id)
	c.Data["json"] = info
	c.ServeJSON()
}
