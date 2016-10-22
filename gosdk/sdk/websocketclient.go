package sdk

import (
	"encoding/base64"
	"gosdk/sysconfig"

	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var Ws *websocket.Conn
var Wsconn_state = make(chan int, 2)
var mutex sync.Mutex

const (
	connected = iota
	disconnected
)

//export GetMacAddr
func GetMacAddr() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Poor soul, here is what you got: " + err.Error())
		return ""
	}
	for _, inter := range interfaces {
		if inter.HardwareAddr != nil && inter.Flags&net.FlagUp != 0 && inter.Flags&net.FlagBroadcast != 0 {
			var str string = inter.HardwareAddr.String()
			return str
		}
	}
	return ""
}

func Websocketconn_timer() {
	wsconn_timer := time.NewTicker(1000 * time.Millisecond)

	for {
		select {
		case <-wsconn_timer.C:
			if Ws == nil {
				Wsconn_state <- disconnected
				break
			} else {
				log.Println("ws connect has been success.")
			}

		}
	}
}

func Websocketconn_listen() {
	for {
		select {
		case state, _ := <-Wsconn_state:
			if state == disconnected {
				WsConnect(string([]byte(sysconfig.SystemURL)[7:]), sysconfig.Device.DeviceToken)
				break
			} else if state == connected {
				log.Println("ws connect success.")
				break
			}
		}
	}
}

func WebSocketClient() {
	go Websocketconn_listen()
	WsConnect(string([]byte(sysconfig.SystemURL)[7:]), sysconfig.Device.DeviceToken)
	go Websocketconn_timer()

	go func() {
		for {
			Recv()
		}
	}()
}

//export WsConnect
func WsConnect(wshost string, devicetoken string) {
	var header http.Header
	headerapi := base64.StdEncoding.EncodeToString([]byte(sysconfig.User.Login_id + ":" + sysconfig.User.Access_token + ":" + sysconfig.User.Domain))
	header = make(map[string][]string)
	header.Add("Authorization", "SSO "+headerapi)

	var err_ws error
	u := url.URL{Scheme: "ws", Host: wshost, Path: "/ws/join"}
	Ws, _, err_ws = websocket.DefaultDialer.Dial(u.String()+"?devicetoken="+devicetoken, header)
	log.Printf("ws connecting to %s", u.String())
	if err_ws != nil {
		log.Println("success")
	} else {
		log.Println("failed")
	}
}

//export Send
func Send(para string) {

	log.Println("send:", para)
	if Ws != nil {
		err := Ws.WriteMessage(websocket.TextMessage, []byte(para))
		if err != nil {
			log.Println("send failed:", err)
		}
	} else {
		log.Println("connect fail！")
		return
	}
}

//export Recv
func Recv() string {
	if Ws == nil {
		return ""
	}
	_, message, err := Ws.ReadMessage()

	if err != nil {
		log.Println("read:", err)
		Ws = nil
		//Wsconn_state <- disconnected
		return ""
	}
	//	receivemessage := AnalyzeMessage(message, len(message))
	//	if len(receivemessage) > 1 {
	//		log.Println("data", receivemessage[2])
	//	}
	//	if len(message) < 10 {
	//		return ""
	//	}
	log.Printf("recv:%s", message)
	return string(message)
}

func AnalyzeMessage(buff []byte, length int) []string {
	analMsg := make([]string, 0)
	if length < 9 {
		analMsg = append(analMsg, string(buff[0:]))
	} else {
		strNow := ""
		strNow = string(buff[0:5])
		analMsg = append(analMsg, strNow)
		strNow = string(buff[5:9])
		analMsg = append(analMsg, strNow)
		messqgelen, _ := strconv.Atoi(analMsg[1])
		strNow = string(buff[9 : 9+messqgelen])
		analMsg = append(analMsg, strNow)
		strNow = string(buff[9+messqgelen : 26+messqgelen])
		analMsg = append(analMsg, strNow)
		strNow = string(buff[26+messqgelen:])
		analMsg = append(analMsg, strNow)
	}

	return analMsg
}

func LenToString(length int) string {
	str := strconv.Itoa(length)
	tmp_byte := make([]byte, 4)
	tmp_byte = []byte("0000")
	str_byte := []byte(str)
	for i, j := len(str_byte)-1, len(tmp_byte)-1; i >= 0; i-- {
		tmp_byte[j] = str_byte[i]
		j--
	}
	return string(tmp_byte)
}

//export Wsclose
func Wsclose() {
	if Ws == nil {
		log.Println("连接失效")
		return
	}
	err := Ws.Close()
	if err != nil {
		log.Println("关闭异常：", err)
	}
}
