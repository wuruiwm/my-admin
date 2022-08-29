package util

import (
	"app/global"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eddieivan01/nic"
	"gopkg.in/gomail.v2"
)

func Notice(title, content string) error {
	noticeType := global.Config.AdminConfig.Notice.Type
	if noticeType == "email" {
		return NoticeEmail(title, content)
	} else if noticeType == "gotify" {
		return NoticeGotify(title, content)
	}
	return errors.New("通知类型错误")
}

func NoticeEmail(title, content string) error {
	email := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	email.SetHeader("From", email.FormatAddress(global.Config.AdminConfig.Notice.EmailUsername, "后台管理系统"))
	email.SetHeader("To", global.Config.AdminConfig.Notice.EmailReceiveUser)
	email.SetHeader("Subject", title)
	email.SetBody("text/plain", content)
	conn := gomail.NewDialer(
		global.Config.AdminConfig.Notice.EmailServerHost,
		global.Config.AdminConfig.Notice.EmailPort,
		global.Config.AdminConfig.Notice.EmailUsername,
		global.Config.AdminConfig.Notice.EmailPassword,
	)
	conn.TLSConfig = &tls.Config{InsecureSkipVerify: global.Config.AdminConfig.Notice.EmailIsEncrypt == 1}
	if err := conn.DialAndSend(email); err != nil {
		return errors.New("邮件发送失败 error:" + err.Error())
	}
	return nil
}

type NoticeGotifyError struct {
	Error            string `json:"error"`
	ErrorCode        int    `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

func NoticeGotify(title, content string) error {
	url := fmt.Sprintf("%s/message?token=%s", global.Config.AdminConfig.Notice.GotifyServerUrl, global.Config.AdminConfig.Notice.GotifyServerToken)
	resp, err := nic.Post(url, nic.H{
		JSON: nic.KV{
			"title":   title,
			"message": content,
		},
	})
	if err != nil {
		return errors.New("发送消息,请求失败 error:" + err.Error())
	}
	noticeGotifyError := &NoticeGotifyError{}
	if err = json.Unmarshal(resp.Bytes, noticeGotifyError); err != nil {
		return errors.New("json反序列化 error:" + err.Error())
	}
	if noticeGotifyError.ErrorCode != 0 {
		return errors.New("发送消息错误 error:" + noticeGotifyError.ErrorDescription)
	}
	return nil
}
