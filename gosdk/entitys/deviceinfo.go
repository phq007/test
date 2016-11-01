package entitys

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(DeviceInfo))
	time.Now()
}

type DeviceInfo struct {
	Id                string    `orm:"column(id);pk;size(36);"json:"id"valid:"MaxSize(36)"`
	Name              string    `orm:"column(name);type(varchar);null;size(100);"json:"name"valid:"MaxSize(100)"`
	UserId            string    `orm:"column(user_id);type(varchar);null;size(36);"json:"user_id"valid:"MaxSize(36)"`
	DeviceToken       string    `orm:"column(device_token);type(varchar);size(100);"json:"device_token"valid:"MaxSize(100)"`
	SystemType        string    `orm:"column(system_type);type(varchar);size(100);"json:"system_type"valid:"MaxSize(100)"`
	AreaId            string    `orm:"column(area_id);type(varchar);size(36);"json:"area_id"valid:"MaxSize(36)"`
	GroupId           string    `orm:"column(group_id);type(varchar);size(36);"json:"group_id"valid:"MaxSize(36)"`
	LatitudeLongitude string    `orm:"column(latitude_longitude);type(varchar);size(30);"json:"latitude_longitude"valid:"MaxSize(30)"`
	Address           string    `orm:"column(address);type(varchar);size(200);"json:"address"valid:"MaxSize(200)"`
	Domain            string    `orm:"column(domain);type(varchar);null;size(100);"json:"domain"valid:"MaxSize(100)"`
	CreateBy          string    `orm:"column(create_by);type(varchar);null;size(36);"json:"create_by"valid:"MaxSize(36)"`
	CreateDate        time.Time `orm:"column(create_date);type(datetime);null;"json:"create_date"`
	LastLoginDate     time.Time `orm:"column(last_login_date);type(datetime);null;"json:"last_login_date"`
	LastRequestDate   time.Time `orm:"column(last_request_date);type(datetime);null;"json:"last_request_date"`
	Status            string    `orm:"column(status);type(varchar);null;size(36);"json:"status"valid:"MaxSize(36)"`
	Expire            string    `orm:"column(expire);type(varchar);null;size(100)"json:"expire"valid:"MaxSize(100)"`
	Appversion        string    `orm:"column(appversion);type(varchar);null;size(20)"json:"version"`
	Appshortversion   int       `orm:"column(appshortversion);type(int);null;size(8)"json:"shortversion"`
	BindDeviceToken   string    `orm:"column(bind_device_token);type(varchar);null;size(20)"json:"bind_device_token"`
	LoginId           string    `orm:"column(login_id);type(varchar);null;size(20)"json:"-"`
	Password          string    `orm:"column(password);type(varchar);null;size(20)"json:"-"`
}

func (this *DeviceInfo) TableName() string {
	return "device_info"
}
