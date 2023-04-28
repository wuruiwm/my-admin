package request

import "errors"

type YoutubeList struct {
	PageBase
}

type YoutubeCreate struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PasswordDelete struct {
	RequiredId
}

func (param *YoutubeCreate) Check() error {
	if param.Name == "" {
		return errors.New("请输入音乐名")
	}
	if param.Url == "" {
		return errors.New("请输入youtube地址")
	}
	return nil
}
