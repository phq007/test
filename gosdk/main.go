//package gosdk

package main

import (
	_ "gosdk/routers"
	"gosdk/sysconfig"
	"gosdk/util"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	//Callback external var
	Callback CallBackListener
)

//CallBackListener export CallBackListener
type CallBackListener interface {
	OnReciveMessage(msg string)
}

//CallBackListener export CallBackListener
type GetDataListener interface {
	GetDataPath(t int) string
}

type Getdata_test struct {
	WorkDir     string
	SystemURL   string
	SSOURL      string
	DownLoadDir string
}

func (c *Getdata_test) GetDataPath(t int) string {
	switch t {
	case 0:
		return (c.WorkDir)
	case 1:
		return (c.SystemURL)
	case 2:
		return (c.SSOURL)
	case 3:
		return (c.DownLoadDir)
	}
	return ""
}

func main() {
	var config Getdata_test
	config.WorkDir = "F:/github/workspace"
	//config.SystemURL = "http://idea-code.com:9696"
	config.SystemURL = "http://10.107.55.14:8080"
	config.SSOURL = "https://app.idea-code.com:1443"
	config.DownLoadDir = "F:/github/workspace"
	var callback GetDataListener
	callback = &config
	Init(callback)
	Start()
}

//Init export Init
func Init(callback GetDataListener) error {
	sysconfig.WorkDir = callback.GetDataPath(0)
	sysconfig.SystemURL = callback.GetDataPath(1)
	sysconfig.SSOURL = callback.GetDataPath(2)
	sysconfig.DownLoadDir = callback.GetDataPath(3)
	log.Println("gosdk Init WorkDir:" + sysconfig.WorkDir)
	log.Println("gosdk Init SystemURL:" + sysconfig.SystemURL)
	log.Println("gosdk Init SSOURL:" + sysconfig.SSOURL)
	log.Println("gosdk Init DownLoadDir:" + sysconfig.DownLoadDir)
	return nil
}

//Start export Start
func Start() {
	beego.BConfig.RunMode = "prod"
	orm.RegisterDataBase("default", "sqlite3", sysconfig.WorkDir+"/data.db")
	orm.RunSyncdb("default", false, true)
	go beego.Run()
	go LogicProcess()
	<-sysconfig.EndRunning
}

//Exec export Exec
func Exec(url string, body string) string {
	log.Println("Exec:" + url + " \nbody:" + body)
	if body != "" {
		return util.HTTPPost("http://127.0.0.1:8080"+url, body)
	}
	return util.HTTPGet("http://127.0.0.1:8080" + url)
}

//SetCallBack export SetCallBack
func SetCallBack(callback CallBackListener) {
	Callback = callback
}

//Stop export Stop
func Stop() {
	sysconfig.EndRunning <- true
}

//LogicProcess main message dispatcher
func LogicProcess() {
	for {
		select {
		case event := <-sysconfig.ChanBusiness:
			{
				log.Println("process_message...")
				if event.Cmd == "quit" {
					return
				} else if event.Cmd == "callback" {
					go Callback.OnReciveMessage(event.Param)
				}
			}
		}
	}
}
