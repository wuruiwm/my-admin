package cron

import (
	"app/global"
	"app/model"
	"app/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eddieivan01/nic"
)

type TwLolLuckDraw struct {
}

func NewTwLolLuckDraw() *TwLolLuckDraw {
	return &TwLolLuckDraw{}
}
func (t *TwLolLuckDraw) Run() {
	if err := t.luckDraw(); err != nil {
		fmt.Println(err)
	}
}

func (t *TwLolLuckDraw) initSk() string {
	return util.AdminConfig("script.tw_lol_luck_draw_sk")
}

type TwLolLuckDrawError struct {
	Error  int    `json:"error"`
	Detail string `json:"detail"`
}

func (t *TwLolLuckDrawError) Err() error {
	if t.Error == 11 {
		return errors.New("用户鉴权失败,请更新sk")
	} else if t.Error == 20 {
		return errors.New("抽奖cd中")
	} else if t.Error != 0 {
		return errors.New("其他错误:" + t.Detail)
	}
	return nil
}

type InitVersion struct {
	Result struct {
		Settings []struct {
			Code    string `json:"code"`
			Version int    `json:"version"`
		} `json:"settings"`
	} `json:"result"`
}

func (t *TwLolLuckDraw) initVersion(sk string) (int, error) {
	url := fmt.Sprintf("https://luckydraw.gamehub.garena.tw/service/luckydraw/?sk=%s&region=TW&tid=%d", sk, util.GetUnix())
	resp, err := nic.Get(url, nil)
	if err != nil {
		return 0, errors.New("获取抽奖版本号请求失败 error:" + err.Error())
	}
	twLolLuckDrawError := &TwLolLuckDrawError{}
	if err := json.Unmarshal(resp.Bytes, twLolLuckDrawError); err != nil {
		return 0, errors.New("json反序列化失败 error:" + err.Error())
	}
	if err := twLolLuckDrawError.Err(); err != nil {
		return 0, err
	}
	initVersion := &InitVersion{}
	if err := json.Unmarshal(resp.Bytes, initVersion); err != nil {
		return 0, errors.New("json反序列化失败 error:" + err.Error())
	}
	version := 0
	for _, v := range initVersion.Result.Settings {
		if v.Code == "lol" {
			version = v.Version
		}
	}
	if version == 0 {
		return 0, errors.New("获取抽奖版本号失败")
	}
	return version, nil
}

type LuckDraw struct {
	Result struct {
		Prize struct {
			Item struct {
				Name string `json:"name"`
			} `json:"item"`
		} `json:"prize"`
	} `json:"result"`
}

func (t *TwLolLuckDraw) luckDraw() error {
	sk := t.initSk()
	version, err := t.initVersion(sk)
	if err != nil {
		return err
	}
	resp, err := nic.Post("https://luckydraw.gamehub.garena.tw/service/luckydraw", nic.H{
		Data: nic.KV{
			"game":    "lol",
			"sk":      sk,
			"region":  "TW",
			"version": fmt.Sprintf("%d", version),
			"tid":     fmt.Sprintf("%d", util.GetUnix()),
		},
	})
	if err != nil {
		return errors.New("抽奖请求失败 error:" + err.Error())
	}
	twLolLuckDrawError := &TwLolLuckDrawError{}
	if err := json.Unmarshal(resp.Bytes, twLolLuckDrawError); err != nil {
		return errors.New("json反序列化失败 error:" + err.Error())
	}
	if err := twLolLuckDrawError.Err(); err != nil {
		return err
	}
	luckDraw := &LuckDraw{}
	if err := json.Unmarshal(resp.Bytes, luckDraw); err != nil {
		return errors.New("json反序列化失败 error:" + err.Error())
	}
	twLolLuckDraw := &model.TwLolLuckDraw{
		Prize:    luckDraw.Result.Prize.Item.Name,
		Response: resp.Text,
	}
	if err := global.Db.Create(twLolLuckDraw).Error; err != nil {
		return errors.New("抽奖记录存入数据库 error:" + err.Error())
	}
	return nil
}
