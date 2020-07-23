package api

import (
	"fmt"
	"net/url"
	"strings"
)

const API_BASE_URL = "https://s3.kingtime.jp/gateway/bprgateway"

type User struct {
	Name         string
	user_token   string
	auth_token   string
	clock_in_id  string
	clock_out_id string
}

func Login(userid, password string) *User {
	data := url.Values{}
	data.Set("page_id", "account_verify")
	data.Set("account", userid)
	data.Set("password", password)
	res_data, _ := NewClient(API_BASE_URL, nil).postForm("", strings.NewReader(data.Encode()))
	return &User{
		Name:         res_data.Get("user_data.user.name").String(),
		user_token:   res_data.Get("user_data.user.user_token").String(),
		auth_token:   res_data.Get("user_data.token.token_b").String(),
		clock_in_id:  res_data.Get("user_data.timerecorder.record_button.0.id").String(),
		clock_out_id: res_data.Get("user_data.timerecorder.record_button.1.id").String(),
	}
}

func ClockIn(user *User) {
	data := url.Values{}
	data.Set("id", user.clock_in_id)
	data.Set("user_token", user.user_token)
	data.Set("token", user.auth_token)
	res_data, _ := NewClient(API_BASE_URL, nil).postForm("", strings.NewReader(data.Encode()))
	fmt.Println(res_data)
}

func ClockOut(user *User) {
	data := url.Values{}
	data.Set("id", user.clock_out_id)
	data.Set("user_token", user.user_token)
	data.Set("token", user.auth_token)
	res_data, _ := NewClient(API_BASE_URL, nil).postForm("", strings.NewReader(data.Encode()))
	fmt.Println(res_data)
}
