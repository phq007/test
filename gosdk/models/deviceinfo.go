package models

import (
	"gosdk/entitys"
	"log"

	"github.com/astaxie/beego/orm"
)

type DeviceInfoManager struct {
}

func (c *DeviceInfoManager) InsertDeviceInfo(device entitys.DeviceInfo) error {
	o := orm.NewOrm()
	id, err := o.Insert(&device)
	if err == nil {
		log.Println(id)
	} else {
		log.Println(err.Error())
	}
	return err
}

func (c *DeviceInfoManager) GetDeviceInfo(did string) (d *entitys.DeviceInfo, err error) {
	o := orm.NewOrm()
	device := entitys.DeviceInfo{}

	selector := o.Raw("select * from device_info where device_token =?", did)
	err = selector.QueryRow(&device)
	if err != nil {
	}
	return &device, err
}
