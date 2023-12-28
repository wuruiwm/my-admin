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

var data []Dns

type Dns struct {
	DomainName string `json:"domain_name"`
	InteriorIp string `json:"interior_ip"`
	ExternalIp string `json:"external_ip"`
}

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
	data = make([]Dns, 0)
	//初始化后台dns配置到data
	err := sonic.Unmarshal([]byte(global.Config.AdminConfig.Cloudflare.Dns), &data)
	if err != nil {
		util.Logger.Error("Cloudflare", util.Map{
			"msg": "初始化后台dns配置失败 error:" + err.Error() + " config:" + global.Config.AdminConfig.Cloudflare.Dns,
		})
		return
	}
	response, err := CloudflareList()
	if err != nil {
		util.Logger.Error("Cloudflare", util.Map{
			"msg": err.Error(),
		})
		return
	}
	for _, v := range response.Result {
		for _, dns := range data {
			if v.Name == dns.DomainName {
				err = dns.CloudflareUpdate(v.Id)
				if err == nil {
					util.Logger.Info("Cloudflare", util.Map{
						"msg": "修改dns解析记录成功",
					})
				}
			}
		}
	}
}

func CloudflareList() (*CloudflareListResponse, error) {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", CloudflareZoneId())
	header := CloudflareHeader()
	result, err := nic.Get(url, nic.H{
		Headers: header,
	})
	if err != nil {
		util.Logger.Error("Cloudflare", util.Map{
			"url":    url,
			"header": header,
			"msg":    err.Error(),
		})
		return nil, errors.New("请求列表失败 error:" + err.Error())
	}
	response := &CloudflareListResponse{}
	err = sonic.Unmarshal(result.Bytes, response)
	if err != nil {
		util.Logger.Error("Cloudflare", util.Map{
			"url":    url,
			"header": header,
			"msg":    errors.New("json反序列化失败 error:" + err.Error()),
		})
		return nil, errors.New("请求列表失败 error:" + err.Error())
	}
	return response, nil
}

func (d *Dns) CloudflareUpdate(id string) error {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", CloudflareZoneId(), id)
	header := CloudflareHeader()
	cloudflareDnsRequest := &CloudflareDnsRequest{
		Content: d.CloudflareIp(),
		Name:    d.DomainName,
		Proxied: false,
		Type:    "A",
		Ttl:     60,
	}
	body, _ := sonic.Marshal(cloudflareDnsRequest)
	_, err := nic.Put(url, nic.H{
		Headers: header,
		Raw:     string(body),
	})
	if err != nil {
		util.Logger.Error("Cloudflare", util.Map{
			"url":    url,
			"header": header,
			"data":   data,
			"msg":    "修改dns解析失败 error" + err.Error(),
		})
		return errors.New("修改dns解析失败")
	}
	return nil
}

func CloudflareHeader() map[string]interface{} {
	return map[string]interface{}{
		"Content-Type": "application/json",
		"X-Auth-Key":   global.Config.AdminConfig.Cloudflare.Key,
		"X-Auth-Email": global.Config.AdminConfig.Cloudflare.Email,
	}
}

func CloudflareZoneId() string {
	return global.Config.AdminConfig.Cloudflare.ZoneId
}

func (d *Dns) CloudflareIp() string {
	hour, _ := strconv.Atoi(time.Now().Format("15"))
	if hour >= 8 && hour < 18 {
		return d.ExternalIp
	} else {
		return d.InteriorIp
	}
}
