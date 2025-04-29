package crontab

import (
	"app/global"
	"app/logic"
	"app/util"
)

func AliyunCdt() {
	title := "阿里云CDT流量监控"
	flow, err := logic.AliyunListCdtInternetTraffic()
	if err != nil {
		util.Notice(title, "获取流量失败 error:"+err.Error())
		return
	}
	util.CacheSet("AliyunCdtFlow", util.Map{
		"flow": flow,
	}, 86400)
	isEnable, err := logic.AliyunDescribeSecurityGroupAttribute()
	if err != nil {
		util.Notice(title, "获取安全组规则状态失败 error:"+err.Error())
		return
	}
	if flow > global.Config.AdminConfig.Aliyun.Limit {
		if isEnable {
			//移除
			ok, err := logic.AliyunRevokeSecurityGroup()
			if ok {
				util.Notice(title, "流量耗尽，移除安全组规则成功")
			} else {
				util.Notice(title, "流量耗尽，移除安全组规则状态失败 error:"+err.Error())
			}
		}
	} else {
		if !isEnable {
			//添加
			ok, err := logic.AliyunAuthorizeSecurityGroup()
			if ok {
				util.Notice(title, "流量恢复，添加安全组规则成功")
			} else {
				util.Notice(title, "流量恢复，添加安全组规则状态失败 error:"+err.Error())
			}
		}
	}
}
