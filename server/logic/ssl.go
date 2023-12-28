package logic

import (
	"app/api/request"
	"app/api/response"
	"app/util"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func Ssl(param *request.Ssl) (*response.Ssl, error) {
	keyPath := param.Domain + ".key"
	pemPath := param.Domain + ".pem"
	keyByte, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, errors.New("读取证书私钥失败 error:" + err.Error())
	}
	pemByte, err := os.ReadFile(pemPath)
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
