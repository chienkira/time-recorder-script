package api

import (
	"bytes"
	"encoding/json"
)

const API_BASE_URL = "https://s3.kingtime.jp/gateway/bprgateway"

func Login(userid, password string) {
	bodyData, _ := json.Marshal(map[string]string{
		"page_id":  "account_verify",
		"account":  userid,
		"password": password,
	})
	NewClient(API_BASE_URL, nil).post("", bytes.NewBuffer(bodyData))
}
