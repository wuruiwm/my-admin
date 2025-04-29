package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/util"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"os"
)

type Domain struct {
	Domain string `json:"domain"`
	Key    string `json:"key"`
	Pem    string `json:"pem"`
}

func Ssl(param *request.Ssl) (*response.Ssl, error) {
	data, err := SslConfig()
	if err != nil {
		return nil, errors.New("读取配置失败 error:" + err.Error())
	}
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
		Domain:     domain.Domain,
		List:       x509Cert.DNSNames,
		ExpireTime: util.TimeToDate(x509Cert.NotAfter),
	}, nil
}

func SslConfig() ([]*Domain, error) {
	data := make([]*Domain, 0)
	//初始化后台dns配置到data
	err := json.Unmarshal([]byte(global.Config.AdminConfig.Ssl.Domain), &data)
	return data, err
}

func SslList() ([]*response.Ssl, error) {
	data, err := SslConfig()
	if err != nil {
		return nil, errors.New("读取配置失败 error:" + err.Error())
	}
	list := make([]*response.Ssl, 0)
	for _, v := range data {
		ssl, err := Ssl(&request.Ssl{
			Domain: v.Domain,
		})
		if err != nil {
			return nil, err
		}
		list = append(list, ssl)
	}
	return list, nil
}
