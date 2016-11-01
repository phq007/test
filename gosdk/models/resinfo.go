package models

import (
	"gosdk/entitys"
	"log"

	"github.com/astaxie/beego/orm"
)

type ResInfoManager struct {
}

func (c *ResInfoManager) Insert(info entitys.ResInfo) error {
	o := orm.NewOrm()
	id, err := o.Insert(&info)
	if err == nil {
		log.Println(id)
	} else {
		log.Println(err.Error())
	}
	return err
}

func (c *ResInfoManager) Get(did string) (d *entitys.ResInfo, err error) {
	o := orm.NewOrm()
	info := entitys.ResInfo{Id: did}
	err = o.Read(&info)
	return &info, err
}
