package crontab

import (
	"app/global"
	"app/model"
	"app/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eddieivan01/nic"
	"time"
)

func TwLolLuckDraw() {
	twLolLuckDrawModel := &model.TwLolLuckDraw{}
	global.Db.Order("create_time desc").Take(&twLolLuckDrawModel)
	unix, _ := util.DateToUnix(twLolLuckDrawModel.CreateTime)
	nowUnix := util.Unix()
	if unix > 0 && nowUnix-unix < (8*3600-600) {
		return
	}
	var (
		prize string
		err   error
	)
	for i := 0; i < 120; i++ {
		if prize, err = twLolLuckDraw(); err != nil {
			util.Logger.Error("tw_lol_luck_draw", util.Map{
				"error": err.Error(),
			})
			time.Sleep(time.Second * 5)
		} else {
			break
		}
	}
	if err != nil {
		_ = util.Notice("台服lol幸运抽奖", fmt.Sprintf("抽奖失败: %s", err.Error()))
	} else {
		_ = util.Notice("台服lol幸运抽奖", prize)
	}
}

type twLolLuckDrawError struct {
	Error  int    `json:"error"`
	Detail string `json:"detail"`
}

func (t *twLolLuckDrawError) Err() error {
	if t.Error == 11 {
		return errors.New("用户鉴权失败,请更新sk")
	} else if t.Error == 20 {
		return errors.New("抽奖cd中")
	} else if t.Error != 0 {
		return errors.New("其他错误:" + t.Detail)
	}
	return nil
}

type twLolLuckDrawVersionResponse struct {
	Result struct {
		Settings []struct {
			Code    string `json:"code"`
			Version int    `json:"version"`
		} `json:"settings"`
	} `json:"result"`
}

func twLolLuckDrawVersion(sk string) (int, error) {
	url := fmt.Sprintf("https://luckydraw.gamehub.garena.tw/service/luckydraw/?sk=%s&region=TW&tid=%d", sk, util.Unix())
	resp, err := nic.Get(url, nil)
	if err != nil {
		return 0, errors.New("获取抽奖版本号请求失败 error:" + err.Error())
	}
	luckDrawError := &twLolLuckDrawError{}
	if err := json.Unmarshal(resp.Bytes, luckDrawError); err != nil {
		return 0, errors.New("json反序列化失败 error:" + err.Error())
	}
	if err := luckDrawError.Err(); err != nil {
		return 0, err
	}
	versionResponse := &twLolLuckDrawVersionResponse{}
	if err := json.Unmarshal(resp.Bytes, versionResponse); err != nil {
		return 0, errors.New("json反序列化失败 error:" + err.Error())
	}
	version := 0
	for _, v := range versionResponse.Result.Settings {
		if v.Code == "lol" {
			version = v.Version
		}
	}
	if version == 0 {
		return 0, errors.New("获取抽奖版本号失败")
	}
	return version, nil
}

type twLolLuckDrawResponse struct {
	Result struct {
		Prize struct {
			Item struct {
				Name string `json:"name"`
			} `json:"item"`
		} `json:"prize"`
	} `json:"result"`
}

func twLolLuckDraw() (string, error) {
	version, err := twLolLuckDrawVersion(global.Config.AdminConfig.Script.TwLolLuckDrawSk)
	if err != nil {
		return "", err
	}
	resp, err := nic.Post("https://luckydraw.gamehub.garena.tw/service/luckydraw", nic.H{
		Data: nic.KV{
			"game":    "lol",
			"sk":      global.Config.AdminConfig.Script.TwLolLuckDrawSk,
			"region":  "TW",
			"version": fmt.Sprintf("%d", version),
			"tid":     fmt.Sprintf("%d", util.Unix()),
		},
	})
	if err != nil {
		return "", errors.New("抽奖请求失败 error:" + err.Error())
	}
	luckDrawError := &twLolLuckDrawError{}
	if err = json.Unmarshal(resp.Bytes, luckDrawError); err != nil {
		return "", errors.New("json反序列化失败 error:" + err.Error())
	}
	if err = luckDrawError.Err(); err != nil {
		return "", err
	}
	luckDrawResponse := &twLolLuckDrawResponse{}
	if err = json.Unmarshal(resp.Bytes, luckDrawResponse); err != nil {
		return "", errors.New("json反序列化失败 error:" + err.Error())
	}
	if err = global.Db.Create(&model.TwLolLuckDraw{
		Prize:    luckDrawResponse.Result.Prize.Item.Name,
		Response: resp.Text,
	}).Error; err != nil {
		return "", errors.New("抽奖记录存入数据库 error:" + err.Error())
	}
	return luckDrawResponse.Result.Prize.Item.Name, nil
}
