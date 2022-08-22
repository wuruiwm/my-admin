package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

//ExcelExport 导出excel表格
func ExcelExport(title []string, data interface{}) (*excelize.File, error) {
	//断言data类型
	val, ok := data.([][]interface{})
	if !ok {
		return nil, errors.New("type error")
	}
	//new excel对象实例
	excel := excelize.NewFile()
	//写入标题
	err := excel.SetSheetRow("Sheet1", "A1", &title)
	if err != nil {
		return nil, errors.New("写入excel表格标题失败 error: " + err.Error())
	}
	//写入数据
	for k, v := range val {
		err := excel.SetSheetRow("Sheet1", fmt.Sprintf("A%d", k+2), &v)
		if err != nil {
			return nil, errors.New("写入excel表格数据失败 error: " + err.Error())
		}
	}
	//保存excel对象实例中的数据到文件
	//如果不需要保留导出的excel文件做记录 可以不需要保存文件 直接写入到response即可
	err = excel.SaveAs("./resource/excel/" + Uuid() + ".xlsx")
	if err != nil {
		return nil, errors.New("保存excel表格数据失败 error: " + err.Error())
	}

	return excel, nil
}

//ExcelImport 导入excel表格
func ExcelImport(filePath string) ([][]string, error) {
	data := make([][]string, 0)
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return data, errors.New("打开excel表格失败 error: " + err.Error())
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return data, errors.New("读取excel表格Sheet1失败 error: " + err.Error())
	}

	if len(rows) > 1 {
		data = rows[1:]
	}
	return data, nil
}

func ExcelDownload(c *gin.Context, excel *excelize.File, fileName string) error {
	//设置header头
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName+".xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	//写入到response
	err := excel.Write(c.Writer)
	if err != nil {
		return errors.New("写入文件流到response失败 error: " + err.Error())
	}
	return nil
}
