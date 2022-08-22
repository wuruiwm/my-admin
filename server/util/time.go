package util

import (
	"errors"
	"time"
)

//GetUnix 获取当前时间戳(秒)
func GetUnix() int64 {
	return time.Now().Unix()
}

//UnixToDate 格式化时间戳为Y-m-d H:i:s格式
func UnixToDate(unix int64) string {
	return time.Unix(unix, 0).Format("2006-01-02 15:04:05")
}

//GetDate 获取当前Y-m-d H:i:s格式的日期字符串
func GetDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//DateToUnix 日期字符串转时间戳
//参数1 必传 需要格式化的日期
//参数2 非必传 需要格式化的格式 默认2006-01-02 15:04:05
func DateToUnix(dateArr ...string) (int64, error) {
	t, err := DateToTime(dateArr...)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

//DateToTime 日期字符串转time类型
//参数1 必传 需要格式化的日期
//参数2 非必传 需要格式化的格式 默认2006-01-02 15:04:05
func DateToTime(dateArr ...string) (time.Time, error) {
	var (
		timeStr string
		format  string
	)
	if len(dateArr) == 1 {
		timeStr = dateArr[0]
		format = "2006-01-02 15:04:05"
	} else if len(dateArr) == 2 {
		timeStr = dateArr[0]
		format = dateArr[1]
	} else {
		return time.Now(), errors.New("传入的参数过多")
	}
	res, err := time.ParseInLocation(format, timeStr, time.Local)
	if err != nil {
		return time.Now(), errors.New("日期格式化失败")
	}
	return res, nil
}

func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
