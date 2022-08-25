package util

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eddieivan01/nic"
	"gopkg.in/gomail.v2"
	"strconv"
)

func Notice(title, content string) {
	noticeType := AdminConfig("notice.type")
	if noticeType == "email" {
		NoticeEmail(title, content)
	} else if noticeType == "gotify" {
		NoticeGotify(title, content)
	}
}

func NoticeEmail(title, content string) {
	email := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	email.SetHeader("From", email.FormatAddress(AdminConfig("notice.email_username"), "后台管理系统"))
	email.SetHeader("To", AdminConfig("notice.email_receive_user"))
	email.SetHeader("Subject", title)
	email.SetBody("text/plain", content)
	port, err := strconv.Atoi(AdminConfig("notice.email_port"))
	if err != nil {
		fmt.Println(errors.New("转换notice.email_port类型失败 error:" + err.Error()))
		return
	}
	isEncrypt, err := strconv.Atoi(AdminConfig("notice.email_is_encrypt"))
	if err != nil {
		fmt.Println(errors.New("转换notice.email_is_encrypt类型失败 error:" + err.Error()))
		return
	}
	conn := gomail.NewDialer(
		AdminConfig("notice.email_server_host"),
		port,
		AdminConfig("notice.email_username"),
		AdminConfig("notice.email_password"),
	)
	conn.TLSConfig = &tls.Config{InsecureSkipVerify: isEncrypt == 1}
	if err = conn.DialAndSend(email); err != nil {
		fmt.Println(errors.New("邮件发送失败 error:" + err.Error()))
	}
}

type NoticeGotifyError struct {
	Error            string `json:"error"`
	ErrorCode        int    `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

func NoticeGotify(title, content string) {
	url := fmt.Sprintf("%s/message?token=%s", AdminConfig("notice.gotify_server_url"), AdminConfig("notice.gotify_server_token"))
	resp, err := nic.Post(url, nic.H{
		JSON: nic.KV{
			"title":   title,
			"message": content,
		},
	})
	if err != nil {
		fmt.Println(errors.New("发送消息,请求失败 error:" + err.Error()))
	}
	noticeGotifyError := &NoticeGotifyError{}
	if err := json.Unmarshal(resp.Bytes, noticeGotifyError); err != nil {
		fmt.Println(errors.New("json反序列化 error:" + err.Error()))
	}
	if noticeGotifyError.ErrorCode != 0 {
		fmt.Println("发送消息错误 error:" + noticeGotifyError.ErrorDescription)
	}
}
