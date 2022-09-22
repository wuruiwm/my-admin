package crontab

import (
	"app/global"
	"app/model"
	"app/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eddieivan01/nic"
	"go.uber.org/zap"
	"time"
)

func TwLolLuckDraw() {
	retryTime := []int{0, 15, 15, 15, 15, 15}
	errList := make([]error, 0)
	var (
		prize string
		err   error
	)
	for i := 0; i < 6; i++ {
		time.Sleep(time.Duration(retryTime[i]) * time.Second)
		if prize, err = twLolLuckDraw(); err != nil {
			errList = append(errList, err)
			global.Logger.Error("tw_lol_luck_draw", zap.Any("error", err.Error()))
		} else {
			break
		}
	}
	errCount := len(errList)
	if errCount > 0 {
		content := ""
		for k, v := range errList {
			content = content + fmt.Sprintf("第%d次抽奖失败原因：%s\n", k+1, v.Error())
		}
		if errCount < 6 {
			content = content + fmt.Sprintf("第%d次抽奖成功：%s\n", errCount+1, prize)
		}
		_ = util.Notice("台服lol幸运抽奖", content)
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
	sk := global.Config.AdminConfig.Script.TwLolLuckDrawSk
	version, err := twLolLuckDrawVersion(sk)
	if err != nil {
		return "", err
	}
	resp, err := nic.Post("https://luckydraw.gamehub.garena.tw/service/luckydraw", nic.H{
		Data: nic.KV{
			"game":    "lol",
			"sk":      sk,
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
