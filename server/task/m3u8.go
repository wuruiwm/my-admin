package task

import (
	"app/global"
	"app/model"
	"app/util"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"io"
	"os"
	"sync"
)

var m3u8Lock sync.Mutex

type m3u8TasK struct{}

func M3u8() {
	mq, err := util.NewRabbitmq()
	if err != nil {
		panic(err)
	}

	m := &m3u8TasK{}
	mq.Consume("m3u8", "direct", m.handle)
}

func (m *m3u8TasK) handle(delivery amqp091.Delivery, rabbitmq *util.Rabbitmq) {
	m3u8Lock.Lock()
	defer m3u8Lock.Unlock()
	//将消息反序列化
	m3u8 := &model.M3u8{}
	err := json.Unmarshal(delivery.Body, m3u8)
	if err != nil {
		util.Logger.Error("m3u8", util.Map{
			"name": m3u8.Name,
			"url":  m3u8.Url,
			"msg":  "json error: " + err.Error(),
		})
		_ = delivery.Ack(true)
		return
	}
	if err := global.Db.Where("id", m3u8.Id).Take(m3u8).Error; err != nil {
		util.Logger.Error("m3u8", util.Map{
			"name": m3u8.Name,
			"url":  m3u8.Url,
			"msg":  "query error: " + err.Error(),
		})
		_ = delivery.Ack(true)
		return
	}
	if m3u8.Status != 0 {
		util.Logger.Error("m3u8", util.Map{
			"name": m3u8.Name,
			"url":  m3u8.Url,
			"msg":  "该任务已执行完成",
		})
		_ = delivery.Ack(true)
		return
	}
	//调用ffmpeg下载视频 并将结果写入到db
	savePath := "tmp.mp4"
	err, cmd, content := m.download(m3u8.Url, savePath)
	m3u8.Command = cmd
	m3u8.Content = content
	if err != nil {
		util.Logger.Error("m3u8", util.Map{
			"name": m3u8.Name,
			"url":  m3u8.Url,
			"msg":  "download error:" + err.Error(),
		})
		m3u8.Status = 2
		err = global.Db.Where("id", m3u8.Id).
			Select("status", "command", "content").
			Updates(m3u8).Error
		if err != nil {
			util.Logger.Error("m3u8", util.Map{
				"name": m3u8.Name,
				"url":  m3u8.Url,
				"msg":  "db error:" + err.Error(),
			})
		}
		_ = delivery.Ack(true)
		return
	}
	//将文件上传到nas
	uploadPath := fmt.Sprintf("%s/%s.mp4", global.Config.AdminConfig.Script.M3u8SaveDir, m3u8.Name)
	err = util.SmbUpload(savePath, uploadPath)
	if err != nil {
		util.Logger.Error("m3u8", util.Map{
			"name": m3u8.Name,
			"url":  m3u8.Url,
			"msg":  "upload error:" + err.Error(),
		})
		m3u8.Status = 2
		m3u8.Content = m3u8.Content + "\n\n" + "upload error:" + err.Error()
	} else {
		m3u8.Status = 1
	}
	err = global.Db.Where("id", m3u8.Id).
		Select("status", "command", "content").
		Updates(m3u8).Error
	if err != nil {
		util.Logger.Error("m3u8", util.Map{
			"name": m3u8.Name,
			"url":  m3u8.Url,
			"msg":  "db error:" + err.Error(),
		})
	}
	_ = delivery.Ack(true)
}

func (m *m3u8TasK) download(m3u8Url string, savePath string) (err error, cmd string, content string) {
	if _, err = os.Stat(savePath); err == nil {
		if err = os.Remove(savePath); err != nil {
			return err, cmd, content
		}
	}
	buf := bytes.NewBuffer([]byte{})
	cmd = fmt.Sprintf(`ffmpeg -threads 8 -i %s -c copy %s`, m3u8Url, savePath)
	err = util.Command(cmd, buf)
	bufByt, err := io.ReadAll(buf)
	if err != nil {
		return err, cmd, content
	}
	content = string(bufByt)
	if err != nil {
		return err, cmd, content
	}
	return nil, cmd, content
}
