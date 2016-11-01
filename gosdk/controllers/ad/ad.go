package controllers

import (
	"encoding/json"
	"gosdk/entitys"
	"gosdk/models"
	"gosdk/sdk"
	"gosdk/sysconfig"
	"gosdk/util"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego"
)

type ADController struct {
	beego.Controller
}

var s_playlist string

func playlist_timer() {
	playlist_timer := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-playlist_timer.C:
			playlist := util.HTTPGet(sysconfig.SystemURL + "/ad_api/list?device_token=" + sysconfig.Device.DeviceToken)
			if s_playlist != playlist {
				log.Println("playlist needed to update.")
				download_palylist(playlist)
				s_playlist = playlist
			} else {
				log.Println("playlist is the latest!")
			}
		}
	}
}

func download_palylist(playlist string) {
	var obj entitys.AdInfoStruct
	json.Unmarshal([]byte(playlist), &obj)
	log.Println(obj.TmplData.TmplId)
	os.Mkdir(sysconfig.DownLoadDir+"/godownload/", os.ModePerm)
	os.Mkdir(sysconfig.WorkDir+"/ui/", os.ModePerm)
	resmanager := models.ResInfoManager{}

	//download ui
	_, err := resmanager.Get(obj.TmplData.Hash)
	if err != nil && obj.TmplData.Hash != "" {
		util.Download(sysconfig.SystemURL+"/download/"+obj.TmplData.TmplId, sysconfig.DownLoadDir+"/godownload/"+obj.TmplData.Hash,
			obj.TmplData.Hash)
		util.Upzipfile(sysconfig.DownLoadDir+"/godownload/"+obj.TmplData.Hash, sysconfig.WorkDir+"/ui/"+obj.TmplData.Hash+"/")
		resmanager.Insert(entitys.ResInfo{Id: obj.TmplData.Hash, Name: "tmpl"})
	}

	//download res
	started := 0
	for _, res := range obj.Res {
		_, err := resmanager.Get(res.Hash)
		if err != nil {
			util.Download(sysconfig.SystemURL+"/download/"+res.AttachmentId, sysconfig.DownLoadDir+"/godownload/"+res.Name,
				res.Hash)
			resmanager.Insert(entitys.ResInfo{Id: res.Hash, ResType: res.ResType, Name: res.Name})

			//start play
			sdk.Callback(sysconfig.BusinessMsg{Cmd: "start"})
			started = 1
		}
	}
	if started == 0 {
		sdk.Callback(sysconfig.BusinessMsg{Cmd: "start"})
	}
	log.Println("download finished!")
}

func (this *ADController) Get() {
	playlist_timer()
}
