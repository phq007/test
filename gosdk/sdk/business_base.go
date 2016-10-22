package sdk

import (
	"crypto/tls"
	"encoding/json"
	"gosdk/sysconfig"
	"gosdk/util"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//StartBusiness start business
func StartBusiness() {
	if sysconfig.Device.LoginId != "" && sysconfig.Device.Password != "" {
		result := Login(sysconfig.SSOURL, sysconfig.Device.LoginId, sysconfig.Device.Password, "")
		if result < 0 {
			log.Fatal("login failed..")
			return
		}
	}
	go util.HTTPGet("http://127.0.0.1:8080/" + sysconfig.BusinessURL)
	go WebSocketClient()
}

//Callback send message to host application
func Callback(body sysconfig.BusinessMsg) {
	b, _ := json.Marshal(body)
	sysconfig.ChanBusiness <- sysconfig.BusinessMsg{Cmd: "callback", Param: string(b)}
}

func Login(loginUrl string, username string, password string, macaddress string) int {
	loginUrl = loginUrl + "/openid/login?username=" + username + "&password=" + password
	loginData := url.Values{"usrname": {username}, "password": {password}}
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	log.Println("loginUrl:", username, password, loginUrl, loginData)
	client := &http.Client{Transport: tr}
	resp, err := client.PostForm(loginUrl, loginData)
	if err != nil {
		log.Println("login fail")
		log.Println(err.Error())
		return -1
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%s\n", data)
		json.Unmarshal(data, &sysconfig.User)
	}
	return 0
}
