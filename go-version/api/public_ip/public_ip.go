package public_ip

import (
	"kot/api"
	"net/url"
)

const API_BASE_URL = "https://api.ipify.org"

func GetIp() string {
	data := url.Values{}
	data.Set("format", "json")
	res_data, _ := api.NewClient(API_BASE_URL, nil).Get("", data)
	return res_data.Get("ip").String()
}
