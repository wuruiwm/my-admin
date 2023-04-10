package crontab

import (
	"app/global"
	"app/util"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/eddieivan01/nic"
	"strconv"
	"time"
)

// 国内ip
var interiorIp = "1.117.115.94"

// 国外ip
var externalIp = "13.224.249.121"

// dns域名
var domainName = "video.nikm.cn"

type CloudflareListResponse struct {
	Result []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"result"`
}

type CloudflareDnsRequest struct {
	Content string `json:"content"`
	Name    string `json:"name"`
	Proxied bool   `json:"proxied"`
	Type    string `json:"type"`
	Ttl     int    `json:"ttl"`
}

func Cloudflare() {
	response, err := CloudflareList()
	if err != nil {
		util.Logger.Error("Cloudflare", map[string]interface{}{
			"msg": err.Error(),
		})
		return
	}
	isExist := false
	for _, v := range response.Result {
		if v.Name == domainName {
			isExist = true
			err = CloudflareUpdate(v.Id)
			if err != nil {
				util.Logger.Error("Cloudflare", map[string]interface{}{
					"msg": "修改dns解析记录失败",
				})
			} else {
				util.Logger.Info("Cloudflare", map[string]interface{}{
					"msg": "修改dns解析记录成功",
				})
			}
			return
		}
	}
	if isExist {
		util.Logger.Error("Cloudflare", map[string]interface{}{
			"msg":    "未找到dns解析记录",
			"result": response.Result,
		})
	}
}

func CloudflareList() (*CloudflareListResponse, error) {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", CloudflareZoneId())
	header := CloudflareHeader()
	for i := 0; i < 3; i++ {
		result, err := nic.Get(url, nic.H{
			Headers: header,
		})
		if err != nil {
			util.Logger.Error("Cloudflare", map[string]interface{}{
				"url":    url,
				"header": header,
				"msg":    err.Error(),
				"num":    i + 1,
			})
			time.Sleep(5 * time.Second)
			continue
		}
		response := &CloudflareListResponse{}
		err = sonic.Unmarshal(result.Bytes, response)
		if err != nil {
			util.Logger.Error("Cloudflare", map[string]interface{}{
				"url":    url,
				"header": header,
				"msg":    err.Error(),
				"num":    i + 1,
			})
			time.Sleep(5 * time.Second)
			continue
		}
		return response, nil
	}
	return nil, errors.New("获取dns解析列表失败")
}

func CloudflareUpdate(id string) error {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", CloudflareZoneId(), id)
	header := CloudflareHeader()
	data := &CloudflareDnsRequest{
		Content: CloudflareIp(),
		Name:    domainName,
		Proxied: false,
		Type:    "A",
		Ttl:     60,
	}
	body, _ := sonic.Marshal(data)
	for i := 0; i < 3; i++ {
		_, err := nic.Put(url, nic.H{
			Headers: header,
			Raw:     string(body),
		})
		if err != nil {
			util.Logger.Error("Cloudflare", map[string]interface{}{
				"url":    url,
				"header": header,
				"data":   data,
				"msg":    err.Error(),
				"num":    i + 1,
			})
			time.Sleep(5 * time.Second)
			continue
		}
		return nil
	}
	return errors.New("修改dns解析失败")
}

func CloudflareHeader() map[string]interface{} {
	return map[string]interface{}{
		"Content-Type": "application/json",
		"X-Auth-Key":   global.Config.AdminConfig.Script.CloudflareKey,
		"X-Auth-Email": global.Config.AdminConfig.Script.CloudflareEmail,
	}
}

func CloudflareZoneId() string {
	return global.Config.AdminConfig.Script.CloudflareZoneId
}

func CloudflareIp() string {
	hour, _ := strconv.Atoi(time.Now().Format("15"))
	if hour >= 8 && hour < 18 {
		return externalIp
	} else {
		return interiorIp
	}
}
