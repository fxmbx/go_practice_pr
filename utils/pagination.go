package utils

type Pagination struct {
	Limit      int32  `json:"limit"`
	Page       int32  `json:"page"`
	Sort       string `json:"sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int32  `json:"total_pages"`
	Data       any    `json:"data"`
}

func (p *Pagination) GetOffset() int32 {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetPage() int32 {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}
func (p *Pagination) GetLimit() int32 {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetSort() string {
	if len(p.Sort) < 1 {
		p.Sort = "Id desc"
	}
	return p.Sort
}
