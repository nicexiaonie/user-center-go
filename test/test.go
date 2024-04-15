package main

import (
	"fmt"
	"github.com/nicexiaonie/gtype"
	user_center_go "github.com/nicexiaonie/user-center-go"
)

func main() {

	user := user_center_go.Init("http://127.0.0.1:8081", "XHY-SERVER")

	user.SetLogHook(func(logContext string) {
		fmt.Println(logContext)
	})

	//userInfo, err := user.GetBaseInfo(10151)
	//fmt.Println(err)
	//fmt.Println(gtype.ToString(userInfo))

	//r1, err := user.ServiceSmsSendLogin("13693189022")
	//fmt.Println(err)
	//fmt.Println(gtype.ToString(r1))
	//
	//r2, err := user.ApiUserLogin(user_center_go.ApiUserLoginReq{
	//    LoginType:          "sms",
	//    PhoneNumber:        r1.Body.PhoneNumber,
	//    SmsCode:            r1.Body.Code,
	//    WechatOpenId:       "",
	//    WechatCode:         "",
	//    BirthTime:          "",
	//    Ip:                 "",
	//    AppVersion:         "",
	//    MachineCode:        "",
	//    PhoneType:          "",
	//    PhoneVersion:       "",
	//    PhoneSystemVersion: "",
	//})
	//fmt.Println(err)
	//fmt.Println(gtype.ToString(r2))

	//r3, err := user.ServiceSmsChangeMobile(10151, "13693189099")
	//fmt.Println(err)
	//fmt.Println(gtype.ToString(r3))
	//
	//r4, err := user.UserChangeMobile(10151, "13693189099", r3.Body.Code)
	//fmt.Println(err)
	//fmt.Println(gtype.ToString(r4))

	//r5, err := user.SetBaseInfo(user_center_go.ApiSetBaseInfo{
	//	UserId:   10151,
	//	NickName: "aefs",
	//})
	//fmt.Println(err)
	//fmt.Println(gtype.ToString(r5))

	r5, err := user.RealName(user_center_go.ApiRealNameReq{
		UserId:   10258,
		RealName: "聂元培",
		IdCard:   "410421199505072511",
	})

	fmt.Println(err)
	fmt.Println(gtype.ToString(r5))

}
