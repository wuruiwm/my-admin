package util

import (
	"fmt"
	"math"
	"strconv"
)

// StringToFloat64 字符串转浮点
func StringToFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// Float64Round 浮点型保留两位小数 四舍五入
func Float64Round(number float64) float64 {
	number = math.Trunc(number*100+0.5) / 100
	number, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", number), 64)
	return number
}

// StringToFloat64Round 字符串转浮点 并保留两位小数 四舍五入
func StringToFloat64Round(str string) (float64, error) {
	number, err := StringToFloat64(str)
	if err != nil {
		return number, err
	}
	return Float64Round(number), nil
}
