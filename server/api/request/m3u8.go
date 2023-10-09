package request

import "errors"

type M3u8List struct {
	PageBase
}

type M3u8Create struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type M3u8Retry struct {
	RequiredId
}

type M3u8Delete struct {
	RequiredId
}

func (param *M3u8Create) Check() error {
	if param.Name == "" {
		return errors.New("请输入视频名")
	}
	if param.Url == "" {
		return errors.New("请输入m3u8地址")
	}
	return nil
}
