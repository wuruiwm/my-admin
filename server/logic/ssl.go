package logic

import (
	"app/api/response"
	"app/global"
	"app/util"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

func Ssl() (*response.Ssl, error) {
	keyPath := global.Config.AdminConfig.Ssl.Key
	pemPath := global.Config.AdminConfig.Ssl.Pem
	keyByte, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, errors.New("读取证书私钥失败 error:" + err.Error())
	}
	pemByte, err := ioutil.ReadFile(pemPath)
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
		ExpireTime: util.TimeToDate(x509Cert.NotAfter),
	}, nil
}
