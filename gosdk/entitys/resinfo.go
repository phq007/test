package entitys

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ResInfo struct {
	Id             string `orm:"column(id);pk;size(60);type(varchar);" json:"id" valid:"MaxSize(60)"`
	Name           string `orm:"column(Name);type(varchar);null;size(100);" json:"Name" valid:"MaxSize(100)"`
	ResType        string `orm:"column(ResType);type(varchar);null;size(100);" json:"ResType" valid:"MaxSize(100)"`
	CtrlId         string `orm:"column(CtrlId);type(varchar);null;size(100);" json:"CtrlId" valid:"MaxSize(100)"`
	AttachmentId   string `orm:"column(AttachmentId);type(varchar);null;size(100);" json:"AttachmentId" valid:"MaxSize(100)"`
	Hash           string `orm:"column(Hash);type(varchar);null;size(100);" json:"Hash" valid:"MaxSize(100)"`
	StartTime      string `orm:"column(StartTime);type(varchar);null;size(100);" json:"StartTime" valid:"MaxSize(100)"`
	EndTime        string `orm:"column(EndTime);type(varchar);null;size(100);" json:"EndTime" valid:"MaxSize(100)"`
	Playcount      int    `orm:"column(Playcount);type(int);null;" json:"Playcount" valid:"MaxSize(100)"`
	TmplData       string `orm:"column(TmplData);type(varchar);null;size(100);" json:"TmplData" valid:"MaxSize(1000)"`
	AttachmentPath string `orm:"column(AttachmentPath);type(varchar);null;size(100);" json:"AttachmentPath" valid:"MaxSize(1000)"`
}

type Tmplinfo struct {
	TmplId string
	Hash   string
}

type AdInfoStruct struct {
	TmplData Tmplinfo
	Res      []ResInfo
}

func init() {
	orm.RegisterModel(new(ResInfo))
	time.Now()
}

func (this *ResInfo) TableName() string {
	return "res_info"
}
