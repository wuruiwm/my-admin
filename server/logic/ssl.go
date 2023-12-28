package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/util"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/bytedance/sonic"
	"os"
)

type Domain struct {
	Domain string `json:"domain"`
	Key    string `json:"key"`
	Pem    string `json:"pem"`
}

func Ssl(param *request.Ssl) (*response.Ssl, error) {
	data := make([]*Domain, 0)
	//初始化后台dns配置到data
	err := sonic.Unmarshal([]byte(global.Config.AdminConfig.Ssl.Domain), &data)
	var domain *Domain
	for _, v := range data {
		if v.Domain == param.Domain {
			domain = v
		}
	}
	if domain == nil {
		return nil, errors.New("域名不存在 error:" + err.Error())
	}
	keyByte, err := os.ReadFile(domain.Key)
	if err != nil {
		return nil, errors.New("读取证书私钥失败 error:" + err.Error())
	}
	pemByte, err := os.ReadFile(domain.Pem)
	if err != nil {
		return nil, errors.New("读取证书公钥失败 error:" + err.Error())
	}
	pemBlock, _ := pem.Decode(pemByte)
	if pemBlock == nil {
		return nil, errors.New("解码证书失败")
	}
	x509Cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return nil, errors.New("读取证书失败 error:" + err.Error())
	}
	return &response.Ssl{
		Key:        string(keyByte),
		Pem:        string(pemByte),
		Domain:     x509Cert.DNSNames,
		ExpireTime: util.TimeToDate(x509Cert.NotAfter),
	}, nil
}
