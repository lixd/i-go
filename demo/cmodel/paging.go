package cmodel

type Page struct {
	Page  int   `json:"page" form:"page"`
	Size  int   `json:"size" form:"size"`
	Total int64 `json:"total"`
}

const MaxPageSize = 200

func NewPaging(page, size int) *Page {

	if page < 0 {
		page = 0
	}

	if size > MaxPageSize {
		size = MaxPageSize
	}

	return &Page{
		Page:  page,
		Size:  size,
		Total: 0,
	}
}

func (p *Page) SetTotal(total int64) *Page {
	p.Total = total
	return p
}

// Skip page 从1开始
func (p *Page) Skip() int64 {
	return int64((p.Page - 1) * p.Size)
}

func (p *Page) Limit() int64 {
	return int64(p.Size)
}

func (p *Page) GetTotal() int64 {
	return p.Total
}
