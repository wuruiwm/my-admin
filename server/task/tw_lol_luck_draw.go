package task

import (
	"app/global"
	"app/model"
	"app/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eddieivan01/nic"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"time"
)

func TwLolLuckDraw() {
	twLolLuckDraw()
	mq, err := util.NewRabbitmq()
	if err != nil {
		panic(err)
	}
	mq.Consume("tw_lol_luck_draw", "delay", func(delivery amqp.Delivery, rabbitmq *util.Rabbitmq) {
		twLolLuckDraw()
		_ = delivery.Ack(true)
	})
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

func twLolLuckDraw() {
	log.Println(1111)
	var (
		prize  string
		err    error
		errNum int
		mq     *util.Rabbitmq
	)
	_ = prize
	for {
		if prize, err = twLolLuckDrawRun(); err != nil {
			global.Logger.Error("tw_lol_luck_draw", zap.Any("error", err.Error()))
			errNum++
			if errNum > 5 {
				_ = util.Notice("台服lol幸运抽奖", fmt.Sprintf("抽奖失败: %s", err.Error()))
				errNum = 0
			}
			time.Sleep(time.Second * 1200)
		} else {
			break
		}
	}
	for {
		mq, err = util.NewRabbitmq()
		if err != nil {
			global.Logger.Error("tw_lol_luck_draw", zap.String("error", err.Error()))
			continue
		}
		if err = mq.Publish("tw_lol_luck_draw", "delay", "1", 8*3601*1000); err != nil {
			global.Logger.Error("tw_lol_luck_draw", zap.String("error", err.Error()))
			continue
		}
		mq.Close()
		break
	}
	_ = util.Notice("台服lol幸运抽奖", prize)
}

func twLolLuckDrawRun() (string, error) {
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
