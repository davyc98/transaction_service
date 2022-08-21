package commons

import (
	"math"

	"github.com/jinzhu/gorm"
)

type Pagination struct {
	TotalRecord         int64       `json:"total_record"`
	TotalRecordFiltered int64       `json:"total_record_filtered"`
	TotalPage           int         `json:"total_page"`
	Records             interface{} `json:"records"`
	Offset              int         `json:"offset"`
	Limit               int         `json:"limit"`
	Page                int         `json:"page"`
	PrevPage            int         `json:"prev_page"`
	NextPage            int         `json:"next_page"`
}

func NewPagination() *Pagination {
	return &Pagination{}
}

func (p *Pagination) Paging(db *gorm.DB, page int, limit int, orderBy []string, result interface{}) (*Pagination, error) {
	var offset int
	if page == 0 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return nil, err
	}

	if len(orderBy) > 0 {
		for _, v := range orderBy {
			db = db.Order(v)
		}
	}
	err = db.Limit(limit).Offset(offset).Find(result).Error
	if err != nil {
		return nil, err
	}

	p.TotalRecord = count
	p.Records = result
	p.Limit = limit
	p.Offset = offset
	p.Page = page
	p.TotalPage = int(math.Ceil(float64(count) / float64(limit)))

	if page > 1 {
		p.PrevPage = page - 1
	} else {
		p.PrevPage = page
	}

	if page == p.TotalPage {
		p.NextPage = page
	} else {
		p.NextPage = page + 1
	}

	return p, err
}
