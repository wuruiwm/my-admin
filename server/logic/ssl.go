package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/util"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Ssl(param *request.Ssl) (*response.Ssl, error) {
	sslPath := strings.TrimRight(global.Config.AdminConfig.Ssl.Path, "/")
	keyPath := fmt.Sprintf("%s/%s.key", sslPath, param.Domain)
	pemPath := fmt.Sprintf("%s/%s.pem", sslPath, param.Domain)
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
