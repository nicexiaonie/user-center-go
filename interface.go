package user_center_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nicexiaonie/ghttp"
	"github.com/nicexiaonie/gtype"
	"time"
)

type User struct {
	Url    string
	Source string
}

func Init(url, source string) User {
	return User{
		Url:    url,
		Source: source,
	}
}

func (u User) ServiceSmsSendLogin(phoneNumber string) (ResponseServiceSmsSendLogin, error) {
	res := ResponseServiceSmsSendLogin{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", map[string]interface{}{
		"phone_number": phoneNumber,
	})
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/service_sms/send_login", request, nil, time.Second*3, 3)
	if err != nil {
		return res, err
	}
	if gr.StatusCode != 200 {
		return res, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
	}
	err = json.Unmarshal([]byte(gr.Body), &res)
	if err != nil {
		return res, errors.New(fmt.Sprintf("解析失败. %s", err))
	}
	return res, nil
}
func (u User) ServiceSmsChangeMobile(userId int64, phoneNumber string) (ResponseServiceSmsSendLogin, error) {
	res := ResponseServiceSmsSendLogin{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", map[string]interface{}{
		"phone_number": phoneNumber,
		"user_id":      userId,
	})
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/service_sms/send_change_mobil", request, nil, time.Second*3, 3)
	if err != nil {
		return res, err
	}
	if gr.StatusCode != 200 {
		return res, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
	}
	err = json.Unmarshal([]byte(gr.Body), &res)
	if err != nil {
		return res, errors.New(fmt.Sprintf("解析失败. %s", err))
	}
	return res, nil
}
func (u User) UserChangeMobile(userId int64, phoneNumber, code string) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", map[string]interface{}{
		"phone_number": phoneNumber,
		"user_id":      userId,
		"sms_code":     code,
	})
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/user/change_mobil", request, nil, time.Second*3, 3)
	if err != nil {
		return res, err
	}
	if gr.StatusCode != 200 {
		return res, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
	}
	err = json.Unmarshal([]byte(gr.Body), &res)
	if err != nil {
		return res, errors.New(fmt.Sprintf("解析失败. %s", err))
	}
	return res, nil
}

func (u User) ApiUserLogin(body ApiUserLoginReq) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", body)
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/user/login", request, nil, time.Second*3, 3)
	if err != nil {
		return res, err
	}
	if gr.StatusCode != 200 {
		return res, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
	}
	err = json.Unmarshal([]byte(gr.Body), &res)
	if err != nil {
		return res, errors.New(fmt.Sprintf("解析失败. %s", err))
	}
	return res, nil
}

func (u User) GetBaseInfo(userId int64) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("secret", 2)
	request.Add("body", map[string]interface{}{
		"user_id": userId,
	})
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/user/base_info", request, nil, time.Second*3, 3)
	if err != nil {
		return res, err
	}
	if gr.StatusCode != 200 {
		return res, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
	}
	err = json.Unmarshal([]byte(gr.Body), &res)
	if err != nil {
		return res, errors.New(fmt.Sprintf("解析失败. %s", err))
	}
	return res, nil
}

func (u User) SetBaseInfo(body ApiSetBaseInfo) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", body)
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/set/base_info", request, nil, time.Second*3, 3)
	if err != nil {
		return res, err
	}
	if gr.StatusCode != 200 {
		return res, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
	}
	err = json.Unmarshal([]byte(gr.Body), &res)
	if err != nil {
		return res, errors.New(fmt.Sprintf("解析失败. %s", err))
	}
	return res, nil
}

func (u User) BindWeChat(body ApiBindWeChatReq) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", body)
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/service_wechat/open_id_bind", request, nil, time.Second*3, 3)
	if err != nil {
		return res, err
	}
	if gr.StatusCode != 200 {
		return res, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
	}
	err = json.Unmarshal([]byte(gr.Body), &res)
	if err != nil {
		return res, errors.New(fmt.Sprintf("解析失败. %s", err))
	}
	return res, nil
}

func (u User) SyncWeChatData(body ApiBindWeChatReq) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", body)
	gr, err := ghttp.PostJsonRetry(u.Url+"/service_wechat/sync_nick_profile", request, nil, time.Second*3, 3)
	if err != nil {
		return res, err
	}
	if gr.StatusCode != 200 {
		return res, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
	}
	err = json.Unmarshal([]byte(gr.Body), &res)
	if err != nil {
		return res, errors.New(fmt.Sprintf("解析失败. %s", err))
	}
	return res, nil
}