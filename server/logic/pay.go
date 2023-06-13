package logic

import (
	"app/global"
	"encoding/json"
)

type pay struct {
	Title  string `json:"title"`
	Qrcode string `json:"qrcode"`
	Type   string `json:"type"`
}

func Pay() ([]*pay, error) {
	list := make([]*pay, 0)
	err := json.Unmarshal([]byte(global.Config.AdminConfig.Pay.Config), &list)
	return list, err
}
