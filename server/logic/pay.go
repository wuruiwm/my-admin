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

type PayResult struct {
	Card string `json:"card"`
	List []*pay `json:"list"`
}

func Pay() (*PayResult, error) {
	list := make([]*pay, 0)
	err := json.Unmarshal([]byte(global.Config.AdminConfig.Pay.Config), &list)
	return &PayResult{
		Card: global.Config.AdminConfig.Pay.Card,
		List: list,
	}, err
}
