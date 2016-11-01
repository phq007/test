package controllers

import (
	"gosdk/sysconfig"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego"
)

type PlayListController struct {
	beego.Controller
}

func (c *PlayListController) GetTmplData() {
	hash := c.Ctx.Input.Param(":hash")
	fi, err := os.Open(sysconfig.WorkDir + "/ui/" + hash + "/AD_UI.json")
	if err != nil {
		c.Ctx.WriteString("{}")
		return
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	c.Ctx.WriteString(string(fd))
}

func (c *PlayListController) Get() {
	var s string = `{
  "TmplData": {
    "TmplId": "b1d392d7-db60-479c-ae96-e7cb7d6fa89b",
    "Hash": "b8da2e53618cbfc99538cddb7ee497f0"
  },
  "Res": [
    {
      "Id": "d6a87208-3d11-40e4-a90c-1f1797ea07a9",
      "Name": "1中环二手车15秒.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "6f9d4c65-f4eb-4aee-8b61-44062c702421",
      "Hash": "d39532750241827641d04ca7f0c0c94a",
      "StartTime": "2016-06-27 09:56:36",
      "EndTime": "2019-08-27 09:56:44",
      "Playcount": 9
    },
    {
      "Id": "2876d6b8-539d-11e6-9ffa-0242c0a80008",
      "Name": "dahongfanqichezuodian-5.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "9c0e7097-ba04-495b-8040-a14a3cc87466",
      "Hash": "31442033546980ebe2f0aa9dd2b297f1",
      "StartTime": "2016-07-06 05:06:17",
      "EndTime": "2019-08-06 05:06:19",
      "Playcount": 0
    },
    {
      "Id": "28770b87-539d-11e6-9ffa-0242c0a80008",
      "Name": "denglupu.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "2aaa2be1-b00f-40fa-a781-cfe2e385f878",
      "Hash": "2a90e58f36422bdf894cec3602a0e119",
      "StartTime": "2016-07-06 05:06:52",
      "EndTime": "2019-08-06 05:06:50",
      "Playcount": 0
    },
    {
      "Id": "2876eeac-539d-11e6-9ffa-0242c0a80008",
      "Name": "donghao-14.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "d5e2be20-db35-4812-a441-5b76ce83a8be",
      "Hash": "0bdea136089be23321184c82318ee27c",
      "StartTime": "2016-07-06 05:07:06",
      "EndTime": "2019-08-06 05:07:08",
      "Playcount": 0
    },
    {
      "Id": "2876da16-539d-11e6-9ffa-0242c0a80008",
      "Name": "huawangzhiniaoku-60.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "9e9825c3-bd7a-41e8-99fd-2d5f9404d1de",
      "Hash": "4c7977b5f0bd57ba47b8630c9136e33e",
      "StartTime": "2016-07-06 05:07:32",
      "EndTime": "2018-08-06 05:07:34",
      "Playcount": 0
    },
    {
      "Id": "2e05d911-33f8-4b0e-9da0-56332779d026",
      "Name": "huikangmingche1-57.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "7c976aaa-3c0b-4e5b-adc2-8263457795d4",
      "Hash": "f09872136e9152aa434ec8def3e7d1c6",
      "StartTime": "2016-07-26 09:58:06",
      "EndTime": "2019-08-27 09:58:11",
      "Playcount": 8
    },
    {
      "Id": "2876f486-539d-11e6-9ffa-0242c0a80008",
      "Name": "JADO捷渡中国-行车记录仪.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "fc7dacbe-50b1-426c-8437-2bce9e614166",
      "Hash": "f16d1aa6f65b50eac11d7a13ddb3b58c",
      "StartTime": "2016-07-06 05:24:19",
      "EndTime": "2019-08-06 05:24:21",
      "Playcount": 0
    },
    {
      "Id": "2876f5bf-539d-11e6-9ffa-0242c0a80008",
      "Name": "jiliangmojing-14.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "03cc42c7-c1ba-4eb4-b4a9-0625a450f48b",
      "Hash": "e0541482b13cbfb4f8f0d11dc53ebb78",
      "StartTime": "2016-07-06 05:08:27",
      "EndTime": "2019-08-06 05:08:28",
      "Playcount": 0
    },
    {
      "Id": "2876eb93-539d-11e6-9ffa-0242c0a80008",
      "Name": "libaoshi_ruihuayou_10.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "e6ea6e9d-739e-46f9-802a-8cd53bdfc9a6",
      "Hash": "dcaccab8d4c087b58c53fbbf2ce16e70",
      "StartTime": "2016-07-06 05:08:49",
      "EndTime": "2022-06-06 05:08:51",
      "Playcount": 0
    },
    {
      "Id": "2876fff0-539d-11e6-9ffa-0242c0a80008",
      "Name": "majisiluntai.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "ab288040-d096-4a1b-a093-1996263c7b1a",
      "Hash": "04f214f4b30d1489797b50fe7aadc87f",
      "StartTime": "2016-07-06 05:09:09",
      "EndTime": "2018-08-06 05:09:11",
      "Playcount": 0
    },
    {
      "Id": "2876fa54-539d-11e6-9ffa-0242c0a80008",
      "Name": "OCT华侨城旅行社01.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "9d34b9d2-7ded-4959-b993-3cc69472fe17",
      "Hash": "01727435a182e17fa88f5aee1e113166",
      "StartTime": "2016-07-06 22:57:15",
      "EndTime": "2019-08-06 22:57:16",
      "Playcount": 0
    },
    {
      "Id": "2876cc01-539d-11e6-9ffa-0242c0a80008",
      "Name": "OCT华侨城旅行社02.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "88ad6190-939c-4e15-b403-72d13d6ca4bc",
      "Hash": "65f990c33b3dcb23daebba6152449739",
      "StartTime": "2016-07-06 23:04:13",
      "EndTime": "2019-08-06 23:04:15",
      "Playcount": 0
    },
    {
      "Id": "2876f3ed-539d-11e6-9ffa-0242c0a80008",
      "Name": "OCT华侨城旅行社04.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "93b6e565-d983-458b-a139-db8f7da05802",
      "Hash": "41cf0f39fbbfcf2dbf4128d8be1e5a7f",
      "StartTime": "2016-07-06 22:46:27",
      "EndTime": "2019-08-06 22:46:29",
      "Playcount": 0
    },
    {
      "Id": "2876ee11-539d-11e6-9ffa-0242c0a80008",
      "Name": "OCT华侨城旅行社05.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "05ff5538-1d4e-42e1-92b4-2a04ec1528b3",
      "Hash": "e178330fe1602b3521b359f6480770c9",
      "StartTime": "2016-07-06 05:27:28",
      "EndTime": "2019-08-06 05:27:33",
      "Playcount": 0
    },
    {
      "Id": "2876fb92-539d-11e6-9ffa-0242c0a80008",
      "Name": "qianshanchaye-10.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "94814d01-fa33-4894-bcf6-5992bb8eb190",
      "Hash": "9bc0ba563a12d57be983b3293b7b1226",
      "StartTime": "2016-07-06 05:09:26",
      "EndTime": "2019-08-06 05:09:28",
      "Playcount": 0
    },
    {
      "Id": "2876ed75-539d-11e6-9ffa-0242c0a80008",
      "Name": "qicheshensuoshilashua.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "10b71003-a687-4da4-b1dd-aff44c6a9f9a",
      "Hash": "e48b8530ab1f595ac88f1194eb1d0e85",
      "StartTime": "2016-07-06 05:04:34",
      "EndTime": "2019-08-06 05:04:36",
      "Playcount": 0
    },
    {
      "Id": "2876db25-539d-11e6-9ffa-0242c0a80008",
      "Name": "SAST先科-双镜头录制.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "3a4ced4d-cc19-4660-9111-02fa93d3dece",
      "Hash": "0f8678822dbb21de293eb91dd2e8f62d",
      "StartTime": "2016-07-06 23:00:26",
      "EndTime": "2019-08-06 23:00:28",
      "Playcount": 0
    },
    {
      "Id": "2877142c-539d-11e6-9ffa-0242c0a80008",
      "Name": "TBBtires通用轮胎.png",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "ac8a041f-882c-4978-bb0e-401916174538",
      "Hash": "60c82ea2a5f1359c7ddadbf880d283ff",
      "StartTime": "2016-07-06 22:44:27",
      "EndTime": "2019-08-06 22:44:28",
      "Playcount": 0
    },
    {
      "Id": "e0840724-6cd1-48d0-bea5-606fb1036966",
      "Name": "TOYO TIRES轮胎.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "5562260e-4c79-41f5-bcb2-e4d1d7c1750c",
      "Hash": "acc2a48590a3a65523fe7b10aaf9b6c6",
      "StartTime": "2016-07-30 14:50:04",
      "EndTime": "2019-08-30 14:50:06",
      "Playcount": 9
    },
    {
      "Id": "28770dfd-539d-11e6-9ffa-0242c0a80008",
      "Name": "yilinaifen-15.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "471d0ed6-a53b-4ea1-9b2a-3e12b9d77651",
      "Hash": "3ea1153270d9ef9b7b88658897f708b9",
      "StartTime": "2016-07-06 05:04:58",
      "EndTime": "2019-08-06 05:05:00",
      "Playcount": 0
    },
    {
      "Id": "2877156d-539d-11e6-9ffa-0242c0a80008",
      "Name": "youlifangrunhuayou-15.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "f6030b98-c087-4b81-a8a0-4de3cdd3bdbc",
      "Hash": "11b7d31de497f525376eaf1e3be5851b",
      "StartTime": "2016-07-06 05:05:41",
      "EndTime": "2019-08-06 05:05:43",
      "Playcount": 0
    },
    {
      "Id": "2876e40d-539d-11e6-9ffa-0242c0a80008",
      "Name": "yuelaimei-zuodian-15.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "eb55dd21-d136-449c-8f2a-4617664cf53a",
      "Hash": "f39e80ad6e2cba98f5f4db875bedc782",
      "StartTime": "2016-07-06 05:05:59",
      "EndTime": "2019-08-06 05:06:01",
      "Playcount": 0
    },
    {
      "Id": "2876d75f-539d-11e6-9ffa-0242c0a80008",
      "Name": "东仆-掸子蜡.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "f2d99596-03a9-465d-81c1-cac196fed54c",
      "Hash": "99ee571b9a5727b8d029913bf39701d6",
      "StartTime": "2016-07-06 05:23:17",
      "EndTime": "2019-08-06 05:23:19",
      "Playcount": 0
    },
    {
      "Id": "fa1c42f9-343e-4665-9feb-7a279673c921",
      "Name": "乌海市钰洋轮胎-富力通.png",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "dc88912d-150e-41a2-b6a5-b162a18227fb",
      "Hash": "f76cb2ac08f9f0ebf4ffe710020aa1b9",
      "StartTime": "2016-07-26 17:58:40",
      "EndTime": "2019-08-27 17:58:43",
      "Playcount": 0
    },
    {
      "Id": "287714cb-539d-11e6-9ffa-0242c0a80008",
      "Name": "五羊进口狐狸皮羊皮车椅垫.png",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "30c7a452-a284-4dbf-abb7-3ea7a0e5c125",
      "Hash": "0d37d87d41624b6218bfd05d128433d1",
      "StartTime": "2016-07-06 22:48:00",
      "EndTime": "2019-08-06 22:48:02",
      "Playcount": 0
    },
    {
      "Id": "ec3896c7-9085-479f-bae7-bb7755fd6440",
      "Name": "伊利奶粉.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "74bc945b-7212-45be-bb8f-9a25d04df250",
      "Hash": "62adbbf8e0b4def211300175be3808be",
      "StartTime": "2016-07-30 14:48:03",
      "EndTime": "2019-08-30 14:48:05",
      "Playcount": 9
    },
    {
      "Id": "6df8b6f8-ef98-4b0c-b4bf-d5dddfbb51ca",
      "Name": "优步车载移动互联系统.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "7aa6ec79-635c-48fa-8ac1-72278aeed895",
      "Hash": "afa76883d0cf5102f4235521a3c9daa5",
      "StartTime": "2016-07-30 14:48:27",
      "EndTime": "2019-08-30 14:48:30",
      "Playcount": 9
    },
    {
      "Id": "14cc2170-2615-48d5-aaa1-8da67892306c",
      "Name": "xiangcunmuge01.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "5bc3e06c-2f2a-4446-87ee-e5e52eed684a",
      "Hash": "7b690cc83485d4fc5ac4df6aeb5516ff",
      "StartTime": "2016-08-04 11:14:50",
      "EndTime": "2019-09-04 11:14:51",
      "Playcount": 9
    },
    {
      "Id": "05e546a8-38d1-40d6-aaa1-76f55c4be1b0",
      "Name": "xiangcunmuge01.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "5bc3e06c-2f2a-4446-87ee-e5e52eed684a",
      "Hash": "7b690cc83485d4fc5ac4df6aeb5516ff",
      "StartTime": "2016-08-04 14:29:32",
      "EndTime": "2019-09-04 14:29:34",
      "Playcount": 8
    },
    {
      "Id": "19e37090-e6d3-4c59-acfe-ed73d7fc528e",
      "Name": "xiangcunmuge02.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "4b9c0cea-3696-45b8-8a63-27be82344d2c",
      "Hash": "c6ff7fe409d5676f6d70e8cd7548decb",
      "StartTime": "2016-08-04 11:15:16",
      "EndTime": "2019-09-04 11:15:18",
      "Playcount": 9
    },
    {
      "Id": "1cb48911-1c9a-4dbb-ae9d-459a7aff02dc",
      "Name": "shengdiantengfeng01.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "70115266-432b-438f-bf04-8e45661efc65",
      "Hash": "4bfda70f1a8353b8392b0e0cc0f7f1de",
      "StartTime": "2016-08-04 14:30:05",
      "EndTime": "2019-09-04 14:30:06",
      "Playcount": 9
    },
    {
      "Id": "281645bd-e6c6-46f3-ac9d-2ae606333870",
      "Name": "shengdiantengfeng01.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "70115266-432b-438f-bf04-8e45661efc65",
      "Hash": "4bfda70f1a8353b8392b0e0cc0f7f1de",
      "StartTime": "2016-08-04 11:13:35",
      "EndTime": "2019-09-04 11:13:41",
      "Playcount": 9
    },
    {
      "Id": "12051598-5c5d-48b7-be5d-1fce0e4c7e19",
      "Name": "shengdiantengfeng02.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "7a5b950a-6852-4464-b3ae-10dfbd41ee68",
      "Hash": "64de940fc7357106e3d8a674fee5a233",
      "StartTime": "2016-08-04 11:14:06",
      "EndTime": "2019-09-04 11:14:07",
      "Playcount": 9
    },
    {
      "Id": "3ce1a183-e0e5-41a5-8356-c091a3ef99a4",
      "Name": "八方达大屏导航1.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "4105cbc8-5172-4674-b575-085958651a52",
      "Hash": "12e51e90f9acc6027343416fd2291826",
      "StartTime": "2016-07-30 14:52:29",
      "EndTime": "2019-08-30 14:52:30",
      "Playcount": 9
    },
    {
      "Id": "1a969f9f-522d-4e61-9073-e8a0068bed40",
      "Name": "八方达大屏导航2.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "a6bbc362-6fb2-4fdc-91cc-c540da79b469",
      "Hash": "0f88e9b551f297915f4549bfc316cec9",
      "StartTime": "2016-07-30 16:04:46",
      "EndTime": "2019-08-30 16:04:48",
      "Playcount": 9
    },
    {
      "Id": "2877160b-539d-11e6-9ffa-0242c0a80008",
      "Name": "南阳市康氏茗茶.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "657e5fb5-d60b-4a5d-ab32-02a821469663",
      "Hash": "64934dfee2e0d7ca41a210618718649e",
      "StartTime": "2016-07-06 05:27:49",
      "EndTime": "2019-08-06 05:27:52",
      "Playcount": 0
    },
    {
      "Id": "2876e7c3-539d-11e6-9ffa-0242c0a80008",
      "Name": "博晟汽修-东昊润滑油.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "d31123dc-1098-4f27-b1e0-fe110dc0c4f4",
      "Hash": "4779adb98085cc0a3b344724d2830c76",
      "StartTime": "2016-07-06 05:25:55",
      "EndTime": "2019-08-06 05:25:57",
      "Playcount": 0
    },
    {
      "Id": "a9aee6b8-0543-403d-8326-1e04b34708c5",
      "Name": "善领-车品旗舰店-电子狗记录仪.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "324296c0-115b-4c0d-8d58-d5f7727f8c6c",
      "Hash": "8815b297e903ab4ee98fcca55380d0c1",
      "StartTime": "2016-07-30 16:17:27",
      "EndTime": "2019-08-30 16:17:28",
      "Playcount": 9
    },
    {
      "Id": "2876f917-539d-11e6-9ffa-0242c0a80008",
      "Name": "四方-慕尚轮胎.png",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "a970447b-e3cf-4fbb-9cf7-ab365e2c7cfe",
      "Hash": "06904cba998ae3de61bd21d1635fea0b",
      "StartTime": "2016-07-06 22:44:49",
      "EndTime": "2019-08-06 22:44:51",
      "Playcount": 0
    },
    {
      "Id": "2876d42d-539d-11e6-9ffa-0242c0a80008",
      "Name": "圣元优聪奶粉.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "70a4ab1c-eb7f-4fea-a2ea-25a25ed9ad90",
      "Hash": "e389324f7f302e2e256c3c7a2c1e9d80",
      "StartTime": "2016-07-06 23:04:50",
      "EndTime": "2019-08-06 23:04:52",
      "Playcount": 0
    },
    {
      "Id": "49f3cd66-0329-4d9a-a3a9-06a1db7b8d30",
      "Name": "天意轮胎.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "786b9836-5e2d-48c9-9258-49edb7d6c09e",
      "Hash": "551d8db2158516987f39b22fa04f244b",
      "StartTime": "2016-07-30 16:01:25",
      "EndTime": "2019-08-30 16:01:27",
      "Playcount": 9
    },
    {
      "Id": "2876d7fc-539d-11e6-9ffa-0242c0a80008",
      "Name": "奇奇正-掸子蜡.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "73fcf7b4-6e7b-42d7-9d55-bfc67ba05cbc",
      "Hash": "91c0d2e0bdd2f8cca61f5798b5bb65c7",
      "StartTime": "2016-07-06 05:26:15",
      "EndTime": "2019-08-06 05:26:18",
      "Playcount": 0
    },
    {
      "Id": "cfbafe5a-1113-4f0f-ab6f-ad3bb0d3bdb6",
      "Name": "奔宝名车-二手名车.mp4",
      "ResType": "0",
      "CtrlId": "A_Video",
      "AttachmentId": "0607ed46-b040-4e1d-82e0-b127e2b13fec",
      "Hash": "2f2715bbce7d8518e7f33552a889f0d4",
      "StartTime": "2016-07-26 17:57:08",
      "EndTime": "2019-08-27 17:57:14",
      "Playcount": 9
    },
    {
      "Id": "784c2af0-ae31-4b5f-98cc-79380bfea58e",
      "Name": "奥特嘉防漏安全轮胎.png",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "e9e4365c-830a-453c-a2e8-83efc6ef7ce9",
      "Hash": "54b2421e3af9db64f15b50e4060273c5",
      "StartTime": "2016-07-30 16:02:43",
      "EndTime": "2019-08-30 16:02:44",
      "Playcount": 9
    },
    {
      "Id": "2876dd4d-539d-11e6-9ffa-0242c0a80008",
      "Name": "妈咪宝贝-婴儿纸尿裤.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "c97e479c-2816-4e8b-94cb-1fe5aa8241dd",
      "Hash": "f530a7f38276f8009abac54ee514f81a",
      "StartTime": "2016-07-06 23:02:17",
      "EndTime": "2019-09-06 23:02:19",
      "Playcount": 0
    },
    {
      "Id": "071422f3-36a6-4e75-b96c-cf32a262c3ae",
      "Name": "安儿乐-婴儿纸尿裤.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "016f36bc-3296-4786-8562-0ae22aa0767f",
      "Hash": "59e3d4e8f6ca58d114fd164e3acd9179",
      "StartTime": "2016-07-30 14:50:52",
      "EndTime": "2019-08-30 14:50:53",
      "Playcount": 9
    },
    {
      "Id": "2876d61c-539d-11e6-9ffa-0242c0a80008",
      "Name": "广州丰田-凯美瑞配套轮胎.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "3cdf4bad-f76c-48fe-8ffc-15717680866d",
      "Hash": "e180e815d5d882f96cc4ec95edec696f",
      "StartTime": "2016-07-06 23:01:11",
      "EndTime": "2019-08-06 23:01:13",
      "Playcount": 0
    },
    {
      "Id": "28770e9c-539d-11e6-9ffa-0242c0a80008",
      "Name": "广州益中-陆战旅户外弯刀猎刀.png",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "6ba3343d-10bb-489f-acc4-4ab4db1f3a20",
      "Hash": "893c911d928d49fbc3db203f629301d3",
      "StartTime": "2016-07-06 05:26:52",
      "EndTime": "2019-08-06 05:26:53",
      "Playcount": 0
    },
    {
      "Id": "2876e36f-539d-11e6-9ffa-0242c0a80008",
      "Name": "广州益中-陆战旅户外移动照明.png",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "99c5e0b5-dd67-47e0-a3ea-5127d1f2def1",
      "Hash": "9c2f762852f9e61d6c6f4afa518a882d",
      "StartTime": "2016-07-06 22:42:22",
      "EndTime": "2019-08-06 22:42:23",
      "Playcount": 0
    },
    {
      "Id": "90782956-f6c6-4596-ac67-8be080d4dd79",
      "Name": "广州益中-陆战旅户外装备.png",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "e91c0094-d44d-4d64-9fe0-1bea0a5b692b",
      "Hash": "51631960b5761c527f851affb1cae4c7",
      "StartTime": "2016-07-30 16:15:21",
      "EndTime": "2019-08-30 16:15:23",
      "Playcount": 9
    },
    {
      "Id": "287716a8-539d-11e6-9ffa-0242c0a80008",
      "Name": "广州益中-陆战旅汽车应急启动电源.png",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "14afc4bb-d944-400f-bfd5-1e7c85b80d9c",
      "Hash": "4be554eff1e0602daff8a1a2fb22d199",
      "StartTime": "2016-07-06 05:28:56",
      "EndTime": "2019-08-06 05:28:58",
      "Playcount": 0
    },
    {
      "Id": "287711b6-539d-11e6-9ffa-0242c0a80008",
      "Name": "广州益中-陆战旅超级强光手电.png",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "c3871c78-f47a-4802-a5ac-02d9374dec64",
      "Hash": "7af72c65d379a85afd73b623f28a0f58",
      "StartTime": "2016-07-06 22:48:21",
      "EndTime": "2019-08-06 22:48:22",
      "Playcount": 0
    },
    {
      "Id": "2876f85b-539d-11e6-9ffa-0242c0a80008",
      "Name": "得益新鲜牛奶.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "490ca642-121d-4260-8752-a6f6215aeb70",
      "Hash": "6c7df9585fdaf7e9aca8af9747591755",
      "StartTime": "2016-07-06 23:03:56",
      "EndTime": "2019-08-06 23:03:57",
      "Playcount": 0
    },
    {
      "Id": "427c3cd1-d862-42ab-902e-10c41420780a",
      "Name": "御品传媒1.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "fd14dbea-3ac7-4dd1-a909-5f39bd99b656",
      "Hash": "3429dbd08fd9be3883b9e73c060b7306",
      "StartTime": "2016-07-30 14:48:56",
      "EndTime": "2019-08-30 14:48:57",
      "Playcount": 9
    },
    {
      "Id": "005a4658-a911-4b2c-90ab-20953feb4dea",
      "Name": "御品传媒2.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "3ef0687e-a5b1-46a3-b4af-6cf6f138962c",
      "Hash": "9fe8d35e895aa46b90c9e983e3a2f456",
      "StartTime": "2016-07-30 16:03:41",
      "EndTime": "2019-08-30 16:03:42",
      "Playcount": 9
    },
    {
      "Id": "43358f15-e421-4684-9872-0168f7f7f314",
      "Name": "御品传媒3.jpg",
      "ResType": "1",
      "CtrlId": "C_Logo",
      "AttachmentId": "47ebe4bf-2914-46dc-89c8-787d04091939",
      "Hash": "320186e0fca43734f645168ed088e00c",
      "StartTime": "2016-07-30 14:22:05",
      "EndTime": "2019-08-30 14:22:06",
      "Playcount": 9
    },
    {
      "Id": "28770301-539d-11e6-9ffa-0242c0a80008",
      "Name": "御品传媒4.jpg",
      "ResType": "1",
      "CtrlId": "C_Qrcode",
      "AttachmentId": "f314db78-bc43-4e43-9066-c59548471b51",
      "Hash": "feebf67319ded7f52a0acf5d811edaed",
      "StartTime": "2016-07-06 23:13:02",
      "EndTime": "2019-08-06 23:13:03",
      "Playcount": 0
    },
    {
      "Id": "2876f080-539d-11e6-9ffa-0242c0a80008",
      "Name": "德运进口牛奶.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "462aa693-66b2-41a5-888c-518091497cad",
      "Hash": "29a33dc1b61c8528c8ad2a188f73561e",
      "StartTime": "2016-07-06 05:23:35",
      "EndTime": "2019-08-06 05:23:37",
      "Playcount": 0
    },
    {
      "Id": "2876f120-539d-11e6-9ffa-0242c0a80008",
      "Name": "恒利集团-爽儿宝纸尿裤.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "67f435c9-03f3-4259-a536-162c2758b39b",
      "Hash": "adceb985816f5672facbb1e1a4989752",
      "StartTime": "2016-07-06 22:45:24",
      "EndTime": "2019-08-06 22:45:26",
      "Playcount": 0
    },
    {
      "Id": "2876efe7-539d-11e6-9ffa-0242c0a80008",
      "Name": "恒发汽车坐垫.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "79afa87d-8e15-4f8c-b62c-dae4fd3c2a68",
      "Hash": "fa177a36ac49d9b5d19576afe0513a8c",
      "StartTime": "2016-07-06 22:39:42",
      "EndTime": "2019-08-06 22:39:44",
      "Playcount": 0
    },
    {
      "Id": "2876e232-539d-11e6-9ffa-0242c0a80008",
      "Name": "旺仔牛奶.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "fcd3483e-2178-471a-b224-ddec9f5db284",
      "Hash": "6cacc0ee3b0d403a7cafa03e897f82ce",
      "StartTime": "2016-07-06 23:03:36",
      "EndTime": "2020-08-06 23:03:37",
      "Playcount": 0
    },
    {
      "Id": "2876fe1a-539d-11e6-9ffa-0242c0a80008",
      "Name": "旺仔牛奶.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "f711e80d-f626-4a88-a0b0-4b9971238935",
      "Hash": "3d81dc289422e5a0d362ab2c70a62a26",
      "StartTime": "2016-07-06 23:03:17",
      "EndTime": "2019-08-06 23:03:18",
      "Playcount": 0
    },
    {
      "Id": "2877048a-539d-11e6-9ffa-0242c0a80008",
      "Name": "朝阳好运橡胶-正霸铁骑士轮胎.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "f5f88fa9-608c-4a53-bb45-11c82feef472",
      "Hash": "36def4a1f0c9e165e797e8ab73001a77",
      "StartTime": "2016-07-07 07:01:59",
      "EndTime": "2019-08-07 07:02:00",
      "Playcount": 0
    },
    {
      "Id": "2876d89e-539d-11e6-9ffa-0242c0a80008",
      "Name": "杜邦-高隔热防爆膜.png",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "ffce91a6-c927-4085-9e6f-7d3ab7f1fce3",
      "Hash": "ae4f4b18f079c5c6819615f00e0e7f33",
      "StartTime": "2016-07-06 23:02:53",
      "EndTime": "2019-08-06 23:02:55",
      "Playcount": 0
    },
    {
      "Id": "2876e0f1-539d-11e6-9ffa-0242c0a80008",
      "Name": "查道客.png",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "b541a8ff-c780-498a-8141-e3d2553b7e45",
      "Hash": "5c9502fae8bcb88c89d1638a490d965d",
      "StartTime": "2016-07-06 22:55:57",
      "EndTime": "2019-08-06 22:55:59",
      "Playcount": 0
    },
    {
      "Id": "287706b9-539d-11e6-9ffa-0242c0a80008",
      "Name": "欧蒂斯导航专卖店.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "0fbc744f-21f5-49ad-a934-8b749eef82ca",
      "Hash": "4b1672c643c8a07cae819091eb141756",
      "StartTime": "2016-07-06 22:54:47",
      "EndTime": "2019-08-06 22:54:48",
      "Playcount": 0
    },
    {
      "Id": "2876e2d4-539d-11e6-9ffa-0242c0a80008",
      "Name": "汇康名车馆03.png",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "f730ef2b-d69b-42cd-96ca-b08b921ebd8d",
      "Hash": "b0d6399aad9dfbacc2ba8ac1eae535c1",
      "StartTime": "2016-07-06 05:25:02",
      "EndTime": "2019-08-06 05:25:04",
      "Playcount": 0
    },
    {
      "Id": "49ff486b-187f-4f91-8207-21265894cd48",
      "Name": "汇康名车馆03.png",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "b72a82ba-7aed-441c-90ce-a64626e713ae",
      "Hash": "b0d6399aad9dfbacc2ba8ac1eae535c1",
      "StartTime": "2016-08-04 11:09:33",
      "EndTime": "2019-09-04 11:09:34",
      "Playcount": 9
    },
    {
      "Id": "2876fc37-539d-11e6-9ffa-0242c0a80008",
      "Name": "汇康名车馆04.png",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "bd820595-e307-4a62-ad58-a37974d092d3",
      "Hash": "dd9369546c085f4bad50e00c6e94246a",
      "StartTime": "2016-07-06 22:43:37",
      "EndTime": "2019-08-06 22:43:43",
      "Playcount": 0
    },
    {
      "Id": "2876e5e5-539d-11e6-9ffa-0242c0a80008",
      "Name": "河南鼎宏实业-雪峰汽车坐垫.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "a9d46c39-6d46-4aa5-84b0-f114864f5573",
      "Hash": "feb71bd6695d4855727969e46ef9f080",
      "StartTime": "2016-07-06 22:44:04",
      "EndTime": "2019-08-06 22:44:05",
      "Playcount": 0
    },
    {
      "Id": "28770fd6-539d-11e6-9ffa-0242c0a80008",
      "Name": "深圳鸿富裕-E滴油.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "06d0b17c-4862-4127-b8ae-108ec836f0de",
      "Hash": "c49aa54acceb312591c045bab544d0b4",
      "StartTime": "2016-07-06 22:43:16",
      "EndTime": "2019-08-06 22:43:18",
      "Playcount": 0
    },
    {
      "Id": "28771742-539d-11e6-9ffa-0242c0a80008",
      "Name": "深圳鸿富裕-E滴油.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "c4d49a67-1925-40c7-b0b9-80e41a893d40",
      "Hash": "90107d3342377e0d9d084637267645a2",
      "StartTime": "2016-07-06 22:58:35",
      "EndTime": "2019-08-06 22:58:37",
      "Playcount": 0
    },
    {
      "Id": "2876de48-539d-11e6-9ffa-0242c0a80008",
      "Name": "深圳鸿富裕-美孚.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "602273bc-bb6c-4f36-80bb-fe330b9faec9",
      "Hash": "388ec2233a93162e39675ff852146a07",
      "StartTime": "2016-07-06 22:55:24",
      "EndTime": "2019-08-06 22:55:26",
      "Playcount": 0
    },
    {
      "Id": "2876ec31-539d-11e6-9ffa-0242c0a80008",
      "Name": "深圳鸿富裕-美孚.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "b87b5b64-f68a-49f5-9afb-5f88cc6553aa",
      "Hash": "0ff68ed662254d980918e81ef779fd9c",
      "StartTime": "2016-07-06 05:23:00",
      "EndTime": "2019-08-06 05:23:01",
      "Playcount": 0
    },
    {
      "Id": "2876eccf-539d-11e6-9ffa-0242c0a80008",
      "Name": "清净园-宗家府泡菜.png",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "ff45d19a-c991-418b-a071-4d70168276a7",
      "Hash": "a947a2f72ce2698e4af9c16cd422068f",
      "StartTime": "2016-07-06 22:38:02",
      "EndTime": "2019-08-06 22:38:04",
      "Playcount": 0
    },
    {
      "Id": "cacf0b02-d641-4e85-96d0-778cad16ce39",
      "Name": "爱贝隆-日本花王-婴儿纸尿裤.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "0961bb4c-1706-4540-86e9-9dff1c5f5b3a",
      "Hash": "be15de673d55f17e97d2bc6c7ffdf690",
      "StartTime": "2016-07-30 14:50:30",
      "EndTime": "2019-08-30 14:50:32",
      "Playcount": 9
    },
    {
      "Id": "f9604802-0593-4a35-8138-555d156d8bc9",
      "Name": "爱贝隆-日本花王-婴儿纸尿裤01.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "dbb2c2b4-2a91-4ba4-8226-5e2871496f2a",
      "Hash": "00ddeb5487eee62a99a1bc58bd630dc3",
      "StartTime": "2016-07-30 14:52:08",
      "EndTime": "2019-08-30 14:52:09",
      "Playcount": 9
    },
    {
      "Id": "d855c52d-54b8-467c-b76f-08a2e51143a0",
      "Name": "璋蘭茶庄.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "6a80f5bd-3418-4fdd-99ff-8b22e18cb48b",
      "Hash": "ef553786dee4769c0f974a02c52ccb2f",
      "StartTime": "2016-07-30 16:03:11",
      "EndTime": "2019-08-30 16:03:12",
      "Playcount": 9
    },
    {
      "Id": "2876fd7a-539d-11e6-9ffa-0242c0a80008",
      "Name": "福勤SUPERTY轮胎1.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "3cfe2444-bc89-49c5-baf2-d79919cc6f2b",
      "Hash": "3da616deb68d8a616c310397629a5d7d",
      "StartTime": "2016-07-06 05:27:11",
      "EndTime": "2019-08-06 05:27:12",
      "Playcount": 0
    },
    {
      "Id": "2876d57d-539d-11e6-9ffa-0242c0a80008",
      "Name": "福勤SUPERTY轮胎2.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "4fa8425f-a732-495d-ae26-f3f41df447e1",
      "Hash": "1e491580857c6861e58d1aacdeee298f",
      "StartTime": "2016-07-06 22:41:53",
      "EndTime": "2019-08-06 22:41:55",
      "Playcount": 0
    },
    {
      "Id": "28770a39-539d-11e6-9ffa-0242c0a80008",
      "Name": "福勤SUPERTY轮胎LOGO.jpg",
      "ResType": "1",
      "CtrlId": "C_Qrcode",
      "AttachmentId": "8ad2517e-21a8-4f17-817b-10cbe9560c59",
      "Hash": "8185937ed2e87adef03fb889c58419b5",
      "StartTime": "2016-07-06 23:14:08",
      "EndTime": "2019-08-06 23:14:10",
      "Playcount": 0
    },
    {
      "Id": "8cc5145a-b04c-4916-bf85-eae9dd101326",
      "Name": "福勤SUPERTY轮胎LOGO.jpg",
      "ResType": "1",
      "CtrlId": "C_Logo",
      "AttachmentId": "7b051133-06e2-496f-9a9a-ddd8a1df2872",
      "Hash": "8185937ed2e87adef03fb889c58419b5",
      "StartTime": "2016-07-30 14:21:38",
      "EndTime": "2019-08-30 14:21:40",
      "Playcount": 9
    },
    {
      "Id": "73392578-5c00-4db9-a555-c080b87f0bbd",
      "Name": "笑笑童.png",
      "ResType": "1",
      "CtrlId": "C_Logo",
      "AttachmentId": "ee484510-9f6e-4fc5-8d72-41d41d45f251",
      "Hash": "c6ec432535a732bcae59070d3dc34928",
      "StartTime": "2016-06-30 14:18:54",
      "EndTime": "2019-08-30 14:18:58",
      "Playcount": 9
    },
    {
      "Id": "2876e18f-539d-11e6-9ffa-0242c0a80008",
      "Name": "美洲豹智能温控汽车座垫.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "9ad3f6db-5104-4f80-8342-2a52f07b7f09",
      "Hash": "9c24c4d9109da558de5ca134d622d68f",
      "StartTime": "2016-07-06 23:04:32",
      "EndTime": "2019-08-06 23:04:33",
      "Playcount": 0
    },
    {
      "Id": "2876d383-539d-11e6-9ffa-0242c0a80008",
      "Name": "美美99.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "c677b111-cfea-4625-b29f-37bb33ab71eb",
      "Hash": "a4d11d431a8a256ab6ebcecd7f9b4eea",
      "StartTime": "2016-07-06 22:45:46",
      "EndTime": "2019-08-06 22:45:48",
      "Playcount": 0
    },
    {
      "Id": "2876faf2-539d-11e6-9ffa-0242c0a80008",
      "Name": "美赞臣安婴宝.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "6f659f9d-914d-4653-bd1e-0f3658e6c342",
      "Hash": "d9c46d650e1f35a4a23a679c7d63bc54",
      "StartTime": "2016-07-06 23:00:48",
      "EndTime": "2019-08-06 23:00:49",
      "Playcount": 0
    },
    {
      "Id": "b1117e40-990d-4dbf-b65d-c6490c6a106c",
      "Name": "聚贤茶叶.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "4df9dc5f-d9ce-40c7-aee6-8eebf2bfa1eb",
      "Hash": "11bff06d74d514bfb95c544554d99d8a",
      "StartTime": "2016-07-30 16:16:38",
      "EndTime": "2019-08-30 16:16:39",
      "Playcount": 9
    },
    {
      "Id": "2876fcd6-539d-11e6-9ffa-0242c0a80008",
      "Name": "航帅-行车记录仪A360.png",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "a30fb400-7e9b-46fb-9ed0-9abdeb896443",
      "Hash": "2cbbd9edae792ebd60cbc3e6cdd1848c",
      "StartTime": "2016-07-07 06:46:08",
      "EndTime": "2019-08-07 06:46:09",
      "Playcount": 0
    },
    {
      "Id": "287717ec-539d-11e6-9ffa-0242c0a80008",
      "Name": "航帅-行车记录仪A360六大看点.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "2fe6f78e-1da6-4078-b3a5-a6bb73b0cb54",
      "Hash": "aba92bbaddcc64aab1a9b0875938aa4e",
      "StartTime": "2016-07-06 22:47:06",
      "EndTime": "2019-08-06 22:47:07",
      "Playcount": 0
    },
    {
      "Id": "28771253-539d-11e6-9ffa-0242c0a80008",
      "Name": "航帅-行车记录仪A360独家首发.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "847dd33e-51c4-4ad3-8da1-ac7fa6fe9135",
      "Hash": "cf11ca6421529f299a35a9c0c10585b3",
      "StartTime": "2016-07-06 22:48:39",
      "EndTime": "2019-08-06 22:48:40",
      "Playcount": 0
    },
    {
      "Id": "9c961a66-26a6-4549-9347-c894bad0cd2c",
      "Name": "花王纸尿裤.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "997563af-f98c-48a6-b6e8-9a64a5ff2b43",
      "Hash": "4ef1763710bccf46687944d6b3578339",
      "StartTime": "2016-07-30 16:15:42",
      "EndTime": "2019-08-30 16:15:43",
      "Playcount": 9
    },
    {
      "Id": "2876ff56-539d-11e6-9ffa-0242c0a80008",
      "Name": "辉大轮胎商行-韩泰轮胎.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "8bc45fdd-e4c8-4c35-801a-74a560f0ad94",
      "Hash": "bce23560e376dafffca9db0aa46b2ad0",
      "StartTime": "2016-07-06 22:40:27",
      "EndTime": "2019-08-06 22:40:29",
      "Playcount": 0
    },
    {
      "Id": "3afbaeb1-2ed0-4d12-8762-e9666a87c765",
      "Name": "达力嘉-润滑油.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "78655dcd-0be8-405f-b137-1b5fb5287d51",
      "Hash": "18f4d363c8826319b618b078d0716eba",
      "StartTime": "2016-07-30 16:05:12",
      "EndTime": "2019-08-30 16:05:14",
      "Playcount": 9
    },
    {
      "Id": "78890887-4163-49f9-8e28-956b9cb48d83",
      "Name": "逍客智商全能跨界车.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "93157422-fe2c-4add-b157-591040fcd331",
      "Hash": "e772e06867d0a69e9b03c7a5d887f09f",
      "StartTime": "2016-07-30 16:17:49",
      "EndTime": "2019-08-30 16:17:50",
      "Playcount": 9
    },
    {
      "Id": "764ebc45-4b2e-4abd-85b2-38f6c65ce3bb",
      "Name": "长城润滑油.png",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "8526f846-4a5a-45fb-a2d7-09081f765572",
      "Hash": "c11cb649ca281a8c587f6319f526f8a2",
      "StartTime": "2016-07-30 16:14:32",
      "EndTime": "2019-08-30 16:14:34",
      "Playcount": 9
    },
    {
      "Id": "2876d1fe-539d-11e6-9ffa-0242c0a80008",
      "Name": "雅培奶粉.jpg",
      "ResType": "1",
      "CtrlId": "B_Picture",
      "AttachmentId": "0e5377b4-3062-4418-b934-94b3efec2db5",
      "Hash": "8743ff6c48c00b804d7c985c30935364",
      "StartTime": "2016-07-06 05:26:35",
      "EndTime": "2019-08-06 05:26:37",
      "Playcount": 0
    },
    {
      "Id": "28770cc6-539d-11e6-9ffa-0242c0a80008",
      "Name": "雅士润滑油.jpg",
      "ResType": "1",
      "CtrlId": "C_Picture",
      "AttachmentId": "95affa87-62f8-48d0-90e3-a2b988a0a0f4",
      "Hash": "9f22bde808eee67838dfb4049b7b6e4b",
      "StartTime": "2016-07-06 22:59:21",
      "EndTime": "2019-08-06 22:59:23",
      "Playcount": 0
    },
    {
      "Id": "dafc265f-80fb-48de-b69c-7f1eb5bb0910",
      "Name": "飞德勒轮胎.png",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "0990f9e7-1b9f-4222-a2d4-2c1e3fe0221f",
      "Hash": "3d21da60e772ba24ba302b03fe5f824b",
      "StartTime": "2016-07-30 16:05:39",
      "EndTime": "2019-08-30 16:05:40",
      "Playcount": 9
    },
    {
      "Id": "0eaefc46-8e12-449d-92cd-2ad4abda67b5",
      "Name": "飞鹤奶粉.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "5ef7ff5c-6af9-434c-aeb2-ffa09a4688ce",
      "Hash": "4d57e6159a05750dfb0e2bf8a431bdfa",
      "StartTime": "2016-07-30 16:06:02",
      "EndTime": "2019-08-30 16:06:04",
      "Playcount": 9
    },
    {
      "Id": "2876f521-539d-11e6-9ffa-0242c0a80008",
      "Name": "龟派车品-龟牌-掸子蜡.jpg",
      "ResType": "1",
      "CtrlId": "E_Customer",
      "AttachmentId": "4807d417-2d61-4ab4-abd2-93a1482d3c4a",
      "Hash": "4f610b63f0c306a6031515aa2f74bf41",
      "StartTime": "2016-07-06 22:41:12",
      "EndTime": "2019-08-06 22:41:15",
      "Playcount": 0
    }
  ]
}`
	c.Ctx.WriteString(s)
}
