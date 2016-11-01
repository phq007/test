package util

import (
	"encoding/base64"
	"gosdk/sysconfig"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//HTTPGet http get
func HTTPGet(serverURL string) string {
	var data []byte
	req, _ := http.NewRequest("GET", serverURL, nil)
	if sysconfig.User.Access_token != "" {
		header := base64.StdEncoding.EncodeToString([]byte(sysconfig.User.Login_id + ":" + sysconfig.User.Access_token + ":" + sysconfig.User.Domain))
		req.Header.Add("Authorization", "SSO "+header)
	}
	client2 := &http.Client{}
	resp, err := client2.Do(req)
	if err != nil {
		log.Println("http get failed:", err)
	} else {
		data, _ = ioutil.ReadAll(resp.Body)
	}
	return string(data)
}

//HTTPpPost http post method
func HTTPPost(serverURL string, bodydata string) string {
	var data []byte
	body := strings.NewReader(bodydata)
	req, _ := http.NewRequest("POST", serverURL, body)
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("Content-Type", "application/json")
	if sysconfig.User.Access_token != "" {
		header := base64.StdEncoding.EncodeToString([]byte(sysconfig.User.Login_id + ":" + sysconfig.User.Access_token + ":" + sysconfig.User.Domain))
		req.Header.Add("Authorization", "SSO "+header)
	}
	client2 := &http.Client{}
	resp, err := client2.Do(req)
	if err != nil {
		log.Println("http get failed:", err)
	} else {
		data, _ = ioutil.ReadAll(resp.Body)
	}
	return string(data)
	/*resp, err := httpPost(serverUrl, "application/json;charset=utf-8", bytes.NewBuffer([]byte(bodydata)))
	if err != nil {
		log.Println("http get failed:", err)
	} else {
		data, _ = ioutil.ReadAll(resp.Body)
	}
	return string(data)*/
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//Download file
func Download(serverurl string, localfile string, hash string) error {
	var len int64
	var file *os.File
	fInfo, err := os.Stat(localfile)
	if os.IsExist(err) {
		len = fInfo.Size()
		file, err = os.OpenFile(localfile, os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		file.Seek(len, 0)
	} else {
		file, err = os.OpenFile(localfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
	}
	var req http.Request
	req.Method = "GET"
	req.Close = true
	req.URL, err = url.Parse(serverurl)
	if err != nil {
		panic(err)
	}
	header := http.Header{}
	header.Set("Range", "bytes="+strconv.FormatInt(len, 10)+"-")
	header.Set("user-agent", "gosdk")
	req.Header = header
	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		panic(err)
	}
	written, err := io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
	println("written: ", written)
	return nil
}
