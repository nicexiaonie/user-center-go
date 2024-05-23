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
	hook   func(logContext string)
}

func Init(url, source string) User {
	return User{
		Url:    url,
		Source: source,
		hook: func(logContext string) {

		},
	}
}

func (u *User) SetLogHook(f func(logContext string)) {
	u.hook = f
}

func (u User) ServiceSmsSendLogin(phoneNumber string) (ResponseServiceSmsSendLogin, error) {
	res := ResponseServiceSmsSendLogin{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", map[string]interface{}{
		"phone_number": phoneNumber,
	})
	u.hook(fmt.Sprintf("Request Uri:/api/service_sms/send_login, body:%s", request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/service_sms/send_login", request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:/api/service_sms/send_login, body:%s", gr.Body))

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
	if res.Code != 0 {
		return res, errors.New(gtype.ToString(res.Message))
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
	u.hook(fmt.Sprintf("Request Uri:/api/service_sms/send_change_mobil, body:%s", request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/service_sms/send_change_mobil", request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:/api/service_sms/send_change_mobil, body:%s", gr.Body))
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
	if res.Code != 0 {
		return res, errors.New(gtype.ToString(res.Message))
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
	u.hook(fmt.Sprintf("Request Uri:/api/user/change_mobil, body:%s", request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/user/change_mobil", request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:/api/user/change_mobil, body:%s", gr.Body))

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
	if res.Code != 0 {
		return res, errors.New(gtype.ToString(res.Message))
	}
	return res, nil
}

func (u User) ApiUserLogin(body ApiUserLoginReq) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", body)
	uri := "/api/user/login"
	u.hook(fmt.Sprintf("Request Uri:%s, body:%s", uri, request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+uri, request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:%s, body:%s", uri, gr.Body))

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
	if res.Code != 0 {
		return res, errors.New(gtype.ToString(res.Message))
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
	uri := "/api/user/base_info"
	u.hook(fmt.Sprintf("Request Uri:%s, body:%s", uri, request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+uri, request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:%s, body:%s", uri, gr.Body))
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
	if res.Code != 0 {
		return res, errors.New(gtype.ToString(res.Message))
	}
	return res, nil
}

func (u User) GetBaseInfoSecret(userId int64) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("secret", 1)
	request.Add("body", map[string]interface{}{
		"user_id": userId,
	})
	uri := "/api/user/base_info"
	u.hook(fmt.Sprintf("Request Uri:%s, body:%s", uri, request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+uri, request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:%s, body:%s", uri, gr.Body))
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
	if res.Code != 0 {
		return res, errors.New(gtype.ToString(res.Message))
	}
	return res, nil
}

func (u User) SetBaseInfo(body ApiSetBaseInfo) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", body)
	uri := "/api/set/base_info"
	u.hook(fmt.Sprintf("Request Uri:%s, body:%s", uri, request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/set/base_info", request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:%s, body:%s", uri, gr.Body))
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
	if res.Code != 0 {
		return res, errors.New(gtype.ToString(res.Message))
	}
	return res, nil
}

func (u User) BindWeChat(body ApiBindWeChatReq) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", body)
	uri := "/api/service_wechat/open_id_bind"
	u.hook(fmt.Sprintf("Request Uri:%s, body:%s", uri, request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+uri, request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:%s, body:%s", uri, gr.Body))
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
	if res.Code != 0 {
		return res, errors.New(gtype.ToString(res.Message))
	}
	return res, nil
}

func (u User) SyncWeChatData(body ApiBindWeChatReq) (ResponseUserBaseInfo, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", body)
	uri := "/api/service_wechat/sync_nick_profile"
	u.hook(fmt.Sprintf("Request Uri:%s, body:%s", uri, request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+"/api/service_wechat/sync_nick_profile", request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:%s, body:%s", uri, gr.Body))
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
	if res.Code != 0 {
		return res, errors.New(gtype.ToString(res.Message))
	}
	return res, nil
}

func (u User) RealName(body ApiRealNameReq) (bool, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", body)
	uri := "/api/authentication/real_name"
	u.hook(fmt.Sprintf("Request Uri:%s, body:%s", uri, request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+uri, request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:%s, body:%s", uri, gr.Body))
	if err != nil {
		return false, err
	}
	if gr.StatusCode != 200 {
		return false, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
	}
	err = json.Unmarshal([]byte(gr.Body), &res)
	if err != nil {
		return false, errors.New(fmt.Sprintf("解析失败. %s", err))
	}
	if res.Code != 0 {
		return false, errors.New(gtype.ToString(res.Message))
	}
	return true, nil
}

func (u User) GetTreeUser(body ApiGetTreeUserReq) (bool, error) {
	res := ResponseUserBaseInfo{}
	request := ghttp.FromValues{}
	request.Add("request_id", gtype.UniqueId())
	request.Add("source", u.Source)
	request.Add("body", body)
	uri := "/api/org/get_tree_user"
	u.hook(fmt.Sprintf("Request Uri:%s, body:%s", uri, request.EncodeJson()))
	gr, err := ghttp.PostJsonRetry(u.Url+uri, request, nil, time.Second*3, 3)
	u.hook(fmt.Sprintf("Response Uri:%s, body:%s", uri, gr.Body))
	if err != nil {
		return false, err
	}
	if gr.StatusCode != 200 {
		return false, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
	}
	err = json.Unmarshal([]byte(gr.Body), &res)
	if err != nil {
		return false, errors.New(fmt.Sprintf("解析失败. %s", err))
	}
	if res.Code != 0 {
		return false, errors.New(gtype.ToString(res.Message))
	}
	return true, nil
}
