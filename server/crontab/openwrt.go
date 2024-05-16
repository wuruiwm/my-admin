package crontab

import (
	"app/global"
	"app/logic"
	"app/util"
	"context"
	"fmt"
	"net"
)

func Openwrt() {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", fmt.Sprintf("%s:53", global.Config.AdminConfig.Openwrt.Host))
		},
	}

	list, err := resolver.LookupIPAddr(context.Background(), global.Config.AdminConfig.Openwrt.Domain)
	if err != nil {
		util.Logger.Error("Openwrt", util.Map{
			"err": "dns查询错误 error:" + err.Error(),
		})
		return
	}

	ok := false
	for _, ip := range list {
		if ip.String() == global.Config.AdminConfig.Openwrt.DomainIp {
			ok = true
			break
		}
	}
	if !ok {
		util.Logger.Info("Openwrt", util.Map{
			"msg": "检测到ip不一致，重启防火墙",
		})
		err := logic.OpenwrtRestartFirewall()
		if err != nil {
			util.Logger.Error("Openwrt", util.Map{
				"err": "重启失败 error:" + err.Error(),
			})
		} else {
			util.Logger.Info("Openwrt", util.Map{
				"msg": "重启成功",
			})
		}
	}
}
