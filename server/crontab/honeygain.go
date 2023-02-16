package crontab

import (
	"app/global"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eddieivan01/nic"
	"go.uber.org/zap"
)

func Honeygain() {
	token := global.Config.AdminConfig.Script.HoneygainToken
	myFlow, err := MyFlow(token)
	if err != nil {
		global.Logger.Error("Honeygain", zap.String("error", err.Error()))
		return
	}
	if myFlow.WinningCredits > 0 {
		global.Logger.Info("Honeygain", zap.String("msg", "今天奖励已获取"))
		return
	}
	needFlow, err := NeedFlow(token)
	if err != nil {
		global.Logger.Error("Honeygain", zap.String("error", err.Error()))
		return
	}
	if myFlow.GatheringBytes > needFlow.Data.MaxBytes {
		openJar, err := OpenJar(token)
		if err != nil {
			global.Logger.Error("Honeygain", zap.String("error", err.Error()))
		} else {
			global.Logger.Info("Honeygain", zap.String("msg", fmt.Sprintf("开罐成功 获取%d", openJar.Data.Credits)))
		}
	} else {
		global.Logger.Info("Honeygain", zap.String("msg", fmt.Sprintf("未完成目标流量 目标%d 完成%d", needFlow.Data.MaxBytes, myFlow.GatheringBytes)))
	}
}

type MyflowResponse struct {
	GatheringBytes int     `json:"gathering_bytes"` //当天获取的流量
	WinningCredits float64 `json:"winning_credits"` //当天获得的开罐奖励
}

type NeedFlowResponse struct {
	Data struct {
		MaxBytes int `json:"max_bytes"` //开罐需要的流量
	} `json:"data"`
}
type OpenJarResponse struct {
	Data struct {
		Credits int `json:"credits"` //开罐获得的流量
	} `json:"data"`
}

// MyFlow 获取账号当天流量
func MyFlow(token string) (*MyflowResponse, error) {
	myflowResponse := &MyflowResponse{}
	url := "https://dashboard.honeygain.com/api/v1/earnings/today"
	response, err := nic.Get(url, nic.H{
		Headers: map[string]interface{}{
			"authorization": token,
		},
	})
	if err != nil {
		return myflowResponse, errors.New("MyFlow 请求失败 error:" + err.Error())
	}
	if response.StatusCode != 200 {
		return myflowResponse, errors.New("MyFlow 请求失败 error:http_code " + fmt.Sprintf("%d", response.StatusCode))
	}
	err = json.Unmarshal(response.Bytes, myflowResponse)
	if err != nil {
		return myflowResponse, errors.New("MyFlow json反序列化失败 error:" + err.Error())
	}
	return myflowResponse, nil
}

// NeedFlow 获取需要开罐的流量
func NeedFlow(token string) (*NeedFlowResponse, error) {
	needflowResponse := &NeedFlowResponse{}
	url := "https://dashboard.honeygain.com/api/v1/contest_winnings"
	response, err := nic.Get(url, nic.H{
		Headers: map[string]interface{}{
			"authorization": token,
		},
	})
	if err != nil {
		return needflowResponse, errors.New("NeedFlow 请求失败 error:" + err.Error())
	}
	if response.StatusCode != 200 {
		return needflowResponse, errors.New("NeedFlow 请求失败 error:http_code " + fmt.Sprintf("%d", response.StatusCode))
	}
	err = json.Unmarshal(response.Bytes, needflowResponse)
	if err != nil {
		return needflowResponse, errors.New("NeedFlow json反序列化失败 error:" + err.Error())
	}
	return needflowResponse, nil
}

// OpenJar 开罐
func OpenJar(token string) (*OpenJarResponse, error) {
	openJarResponse := &OpenJarResponse{}
	url := "https://dashboard.honeygain.com/api/v1/contest_winnings"
	response, err := nic.Post(url, nic.H{
		Headers: map[string]interface{}{
			"authorization": token,
		},
	})
	if err != nil {
		return openJarResponse, errors.New("OpenJar 请求失败 error:" + err.Error())
	}
	if response.StatusCode != 200 {
		return openJarResponse, errors.New("OpenJar 请求失败 error:http_code " + fmt.Sprintf("%d", response.StatusCode))
	}
	err = json.Unmarshal(response.Bytes, openJarResponse)
	if err != nil {
		return openJarResponse, errors.New("OpenJar json反序列化失败 error:" + err.Error())
	}
	return openJarResponse, nil
}
