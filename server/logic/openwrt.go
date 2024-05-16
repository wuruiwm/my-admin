package logic

import (
	"app/global"
	"errors"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"net"
)

func OpenwrtRestartFirewall() error {
	config := &goph.Config{
		User:    global.Config.AdminConfig.Openwrt.Username,
		Addr:    global.Config.AdminConfig.Openwrt.Host,
		Port:    uint(global.Config.AdminConfig.Openwrt.Port),
		Auth:    goph.Password(global.Config.AdminConfig.Openwrt.Password),
		Timeout: goph.DefaultTimeout,
		Callback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	client, err := goph.NewConn(config)
	if err != nil {
		return errors.New("ssh连接失败 error:" + err.Error())
	}
	_, err = client.Run("service firewall restart")
	if err != nil {
		return errors.New("ssh命令执行失败 error:" + err.Error())
	}
	return nil
}
