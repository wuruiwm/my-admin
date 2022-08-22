package request

import (
	"app/util"
	"fmt"
	"strings"
)

type PageBase struct {
	Page     int `form:"page" binding:"required,gte=1"`
	PageSize int `form:"page_size" binding:"required,gte=1"`
	Search
}

type Search struct {
	Keyword   string `form:"keyword"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}

func (p *Search) GetKeywordWhere(field ...string) string {
	var where string
	if p.Keyword == "" {
		return where
	}
	like := "'%" + p.Keyword + "%'"
	for _, v := range field {
		where = fmt.Sprintf("%s%s like %s or ", where, v, like)
	}
	where = strings.TrimRight(where, " or ")
	return where
}

func (p *Search) GetTimeWhere(field ...string) string {
	var where string
	_, startTimeErr := util.DateToUnix(p.StartTime)
	_, endTimeErr := util.DateToUnix(p.EndTime)
	if startTimeErr != nil || endTimeErr != nil {
		return where
	}
	if len(field) == 1 {
		where = fmt.Sprintf("%s >= '%s' and %s <= '%s'", field[0], p.StartTime, field[0], p.EndTime)
	} else if len(field) == 2 {
		where = fmt.Sprintf("(%s >= '%s' and %s <= '%s') or (%s >= '%s' and %s <= '%s') or (%s <= '%s' and %s >= '%s')", field[0], p.StartTime, field[0], p.EndTime, field[1], p.StartTime, field[1], p.EndTime, field[0], p.StartTime, field[1], p.EndTime)
	}
	return where
}

type Import struct {
	ExcelPath string `json:"excel_path" binding:"required"`
}

type RequiredId struct {
	Id string `form:"id" json:"id" binding:"required"`
}
