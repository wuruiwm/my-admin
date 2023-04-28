package task

import (
	"app/global"
	"app/model"
	"app/util"
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/hirochachacha/go-smb2"
	"github.com/rabbitmq/amqp091-go"
	"io"
	"net"
	"os"
	"sync"
)

var lock sync.Mutex
var tmpFileName = "tmp.mp3"

func Youtube() {
	mq, err := util.NewRabbitmq()
	if err != nil {
		panic(err)
	}
	mq.Consume("youtube", "direct", handle)
}

func handle(delivery amqp091.Delivery, rabbitmq *util.Rabbitmq) {
	lock.Lock()
	defer lock.Unlock()
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
	//调用yt-dlp下载视频 并将结果写入到db
	err, cmd, content := download(youtube)
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
	err = upload(youtube)
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

func download(youtube *model.Youtube) (err error, cmd string, content string) {
	if _, err = os.Stat(tmpFileName); err == nil {
		if err = os.Remove(tmpFileName); err != nil {
			return err, cmd, content
		}
	}
	buf := bytes.NewBuffer([]byte{})
	cmd = fmt.Sprintf(`yt-dlp -x --audio-format mp3 -o %s "%s"`, tmpFileName, youtube.Url)
	err = util.Command(cmd, buf)
	bufByt, _ := io.ReadAll(buf)
	content = string(bufByt)
	if err != nil {
		return err, cmd, content
	}
	return nil, cmd, content
}

func upload(youtube *model.Youtube) error {
	//连接smb
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:445", global.Config.AdminConfig.Youtube.Host))
	if err != nil {
		return err
	}
	defer conn.Close()
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     global.Config.AdminConfig.Youtube.Username,
			Password: global.Config.AdminConfig.Youtube.Password,
		},
	}
	s, err := d.Dial(conn)
	if err != nil {
		return err
	}
	defer s.Logoff()
	fs, err := s.Mount(global.Config.AdminConfig.Youtube.MountDir)
	if err != nil {
		return err
	}
	defer fs.Umount()
	//读取本地文件
	fileByt, err := os.ReadFile(tmpFileName)
	if err != nil {
		return err
	}
	//远程文件如果存在 则删除
	filename := fmt.Sprintf("%s\\%s.mp3", global.Config.AdminConfig.Youtube.MusicDir, youtube.Name)
	if _, err = fs.Stat(filename); err == nil {
		if err = fs.Remove(filename); err != nil {
			return err
		}
	}
	//创建远程文件 并写入
	f, err := fs.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(fileByt)
	if err != nil {
		return err
	}
	return nil
}
