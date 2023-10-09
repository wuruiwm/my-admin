package util

import (
	"app/global"
	"fmt"
	"github.com/hirochachacha/go-smb2"
	"net"
	"os"
	"strings"
)

func SmbConn() (net.Conn, error) {
	t := global.Config.AdminConfig
	fmt.Println(t)
	return net.Dial("tcp", fmt.Sprintf("%s:445", global.Config.AdminConfig.Nas.Host))
}

func SmbAuth(conn net.Conn) (*smb2.Session, error) {
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     global.Config.AdminConfig.Nas.Username,
			Password: global.Config.AdminConfig.Nas.Password,
		},
	}
	return d.Dial(conn)
}

func SmbMount(session *smb2.Session, mountDir string) (*smb2.Share, error) {
	return session.Mount(mountDir)
}

func SmbUploadPathHandle(uploadPath string) (mountDir string, savePath string) {
	uploadPath = strings.Trim(strings.ReplaceAll(uploadPath, "/", "\\"), "\\")
	arr := strings.Split(uploadPath, "\\")

	for k, v := range arr {
		if k == 0 {
			mountDir = v
		} else {
			savePath = savePath + v
			if k != (len(arr) - 1) {
				savePath = savePath + "\\"
			}
		}
	}
	return
}

func SmbUpload(filePath string, uploadPath string) error {
	mountDir, savePath := SmbUploadPathHandle(uploadPath)
	//连接smb
	conn, err := SmbConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	session, err := SmbAuth(conn)
	if err != nil {
		return err
	}
	defer session.Logoff()
	fs, err := SmbMount(session, mountDir)
	if err != nil {
		return err
	}
	defer fs.Umount()
	//读取本地文件
	fileByt, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	//远程文件如果存在 则删除
	if _, err = fs.Stat(savePath); err == nil {
		if err = fs.Remove(savePath); err != nil {
			return err
		}
	}
	//创建远程文件 并写入
	f, err := fs.Create(savePath)
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
