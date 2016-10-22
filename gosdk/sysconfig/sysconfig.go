package sysconfig

import (
	"gosdk/entitys"
)

//BusinessMsg struct
type BusinessMsg struct {
	Cmd   string
	Param string
}

//UserLogin user login information struct
type UserLogin struct {
	User_id       string
	Access_token  string
	Refresh_token string
	Login_id      string
	Time_stamp    int64
	Server_ip     string
	Domain        string
}

var (
	//SystemURL server url
	SystemURL string
	//SSOURL sso url
	SSOURL string
	//BusinessURL local business start url
	BusinessURL string
	//WorkDir working dir
	WorkDir string
	//DownLoadDir working dir
	DownLoadDir string
	//EndRunning stop sdk chan
	EndRunning = make(chan bool, 1)
	//ChanBusiness for call back use
	ChanBusiness = make(chan BusinessMsg, 1000)
	//User user login infomation
	User UserLogin
	//Device device info
	Device entitys.DeviceInfo
)

func init() {
	BusinessURL = "ad"
	SystemURL = "http://idea-code.com:9696"
	SSOURL = "https://app.idea-code.com:1443"
}
