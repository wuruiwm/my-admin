package logic

import (
	"app/api/response"
	"app/global"
	"app/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func UploadImage(c *gin.Context, fileField string) (*response.Upload, error) {
	_, fileHeader, err := c.Request.FormFile(fileField)
	if err != nil {
		return nil, errors.New("请上传文件")
	}
	ext := util.GetFileExt(fileHeader.Filename)
	if !util.InArray(ext, []string{".jpg", ".jpeg", ".png", ".gif"}) {
		return nil, errors.New("请上传正确的图片类型 jpg,jpeg,png,gif,bmp")
	}

	//打开图片
	file, err := fileHeader.Open()
	if err != nil {
		return nil, errors.New("打开图片失败 error: " + err.Error())
	}
	defer file.Close()

	//读取图片
	var img image.Image
	if util.InArray(ext, []string{".jpeg", ".jpg"}) {
		img, err = jpeg.Decode(file)
	} else if ext == ".png" {
		img, err = png.Decode(file)
	} else if ext == ".gif" {
		img, err = gif.Decode(file)
	}
	if err != nil {
		return nil, errors.New("读取图片失败 error: " + err.Error())
	}

	//图片宽度调整为400 高度同比例降低
	m := resize.Resize(400, 0, img, resize.Lanczos3)
	filePath := "./resource/upload/" + util.Uuid() + ext
	out, err := os.Create(filePath)
	if err != nil {
		return nil, errors.New("创建图片失败 error: " + err.Error())
	}
	defer out.Close()

	//保存图片
	if util.InArray(ext, []string{".jpeg", ".jpg"}) {
		err = jpeg.Encode(out, m, nil)
	} else if ext == ".png" {
		err = png.Encode(out, m)
	} else if ext == ".gif" {
		err = gif.Encode(out, m, nil)
	}
	if err != nil {
		return nil, errors.New("保存图片失败 error: " + err.Error())
	}

	urlPath := strings.Replace(filePath, "./resource", "", 1)
	return &response.Upload{
		Path: urlPath,
		Url:  global.Config.DomainName + urlPath,
	}, nil
}

func UploadFile(c *gin.Context, fileField string) (*response.Upload, error) {
	_, fileHeader, err := c.Request.FormFile(fileField)
	if err != nil {
		return nil, errors.New("请上传文件")
	}
	filePath := "./resource/upload/" + util.Uuid() + util.GetFileExt(fileHeader.Filename)
	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		return nil, errors.New("保存文件失败 error: " + err.Error())
	}
	urlPath := strings.Replace(filePath, "./resource", "", 1)
	return &response.Upload{
		Path: urlPath,
		Url:  global.Config.DomainName + urlPath,
	}, nil
}
