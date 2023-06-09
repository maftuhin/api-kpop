package paging

import (
	"math"

	"gorm.io/gorm"
)

type Paginator struct {
	TotalRecord int64       `json:"totalRecord"`
	TotalPage   int         `json:"totalPage"`
	Records     interface{} `json:"records"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prevPage"`
	NextPage    int         `json:"nextPage"`
}
type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	ShowSQL bool
}

func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB
	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int64
	var offset int

	db.Model(result).Count(&count)
	done <- true

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(&result)
	<-done

	paginator.TotalRecord = count
	paginator.Records = result
	paginator.Page = p.Page
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator
}
