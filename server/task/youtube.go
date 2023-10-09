package task

import (
	"app/global"
	"app/model"
	"app/util"
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/rabbitmq/amqp091-go"
	"io"
	"os"
	"sync"
)

var youtubeLock sync.Mutex

type youtubeTasK struct{}

func Youtube() {
	mq, err := util.NewRabbitmq()
	if err != nil {
		panic(err)
	}

	y := &youtubeTasK{}
	mq.Consume("youtube", "direct", y.handle)
}

func (y *youtubeTasK) handle(delivery amqp091.Delivery, rabbitmq *util.Rabbitmq) {
	youtubeLock.Lock()
	defer youtubeLock.Unlock()
	//将消息反序列化
	youtube := &model.Youtube{}
	err := sonic.Unmarshal(delivery.Body, youtube)
	if err != nil {
		util.Logger.Error("youtube", util.Map{
			"name": youtube.Name,
			"url":  youtube.Url,
			"msg":  "json error: " + err.Error(),
		})
		_ = delivery.Ack(true)
		return
	}
	if err := global.Db.Where("id", youtube.Id).Take(youtube).Error; err != nil {
		util.Logger.Error("youtube", util.Map{
			"name": youtube.Name,
			"url":  youtube.Url,
			"msg":  "query error: " + err.Error(),
		})
		_ = delivery.Ack(true)
		return
	}
	if youtube.Status != 0 {
		util.Logger.Error("youtube", util.Map{
			"name": youtube.Name,
			"url":  youtube.Url,
			"msg":  "该任务已执行完成",
		})
		_ = delivery.Ack(true)
		return
	}
	//调用yt-dlp下载视频 并将结果写入到db
	savePath := "tmp.mp3"
	err, cmd, content := y.download(youtube.Url, savePath)
	youtube.Command = cmd
	youtube.Content = content
	if err != nil {
		util.Logger.Error("youtube", util.Map{
			"name": youtube.Name,
			"url":  youtube.Url,
			"msg":  "download error:" + err.Error(),
		})
		youtube.Status = 2
		err = global.Db.Where("id", youtube.Id).
			Select("status", "command", "content").
			Updates(youtube).Error
		if err != nil {
			util.Logger.Error("youtube", util.Map{
				"name": youtube.Name,
				"url":  youtube.Url,
				"msg":  "db error:" + err.Error(),
			})
		}
		_ = delivery.Ack(true)
		return
	}
	//将文件上传到nas
	uploadPath := fmt.Sprintf("%s/%s.mp3", global.Config.AdminConfig.Script.YoutubeSaveDir, youtube.Name)
	err = util.SmbUpload(savePath, uploadPath)
	if err != nil {
		util.Logger.Error("youtube", util.Map{
			"name": youtube.Name,
			"url":  youtube.Url,
			"msg":  "upload error:" + err.Error(),
		})
		youtube.Status = 2
		youtube.Content = youtube.Content + "\n\n" + "upload error:" + err.Error()
	} else {
		youtube.Status = 1
	}
	err = global.Db.Where("id", youtube.Id).
		Select("status", "command", "content").
		Updates(youtube).Error
	if err != nil {
		util.Logger.Error("youtube", util.Map{
			"name": youtube.Name,
			"url":  youtube.Url,
			"msg":  "db error:" + err.Error(),
		})
	}
	_ = delivery.Ack(true)
}

func (y *youtubeTasK) download(youtubeUrl string, savePath string) (err error, cmd string, content string) {
	if _, err = os.Stat(savePath); err == nil {
		if err = os.Remove(savePath); err != nil {
			return err, cmd, content
		}
	}
	buf := bytes.NewBuffer([]byte{})
	cmd = fmt.Sprintf(`yt-dlp -x --audio-format mp3 -o %s %s`, savePath, youtubeUrl)
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
