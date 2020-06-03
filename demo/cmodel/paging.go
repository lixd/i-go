package cmodel

type Paging struct {
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

const MaxPageSize = 200

func NewPaging(page, size int) *Paging {

	if page < 0 {
		page = 0
	}

	if size > MaxPageSize {
		size = MaxPageSize
	}

	return &Paging{
		Page:  page,
		Size:  size,
		Total: 0,
	}
}

func (p *Paging) SetTotal(total int64) *Paging {
	p.Total = total
	return p
}

// Skip page 从1开始
func (p *Paging) Skip() int64 {
	return int64((p.Page - 1) * p.Size)
}

func (p *Paging) Limit() int64 {
	return int64(p.Size)
}

func (p *Paging) GetTotal() int64 {
	return p.Total
}
