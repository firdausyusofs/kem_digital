package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	defaultLimit = 10
)

type PaginationQuery struct {
	Limit   int    `query:"limit,omitempty"`
	Page    int    `query:"page,omitempty"`
	OrderBy string `query:"order,omitempty"`
}

func (p *PaginationQuery) setLimit(query string) error {
	if query == "" {
		p.Limit = defaultLimit
		return nil
	}

	n, err := strconv.Atoi(query)
	if err != nil {
		return err
	}

	p.Limit = n

	return nil
}

func (p *PaginationQuery) setPage(query string) error {
	if query == "" {
		p.Page = 1
		return nil
	}

	n, err := strconv.Atoi(query)
	if err != nil {
		return err
	}

	p.Page = n

	return nil
}

func (p *PaginationQuery) setOrder(query string) {
	if query == "" {
		p.OrderBy = "desc"
		return
	}

	p.OrderBy = query
}

func (p *PaginationQuery) GetLimit() int {
	return p.Limit
}

func (p *PaginationQuery) GetPage() int {
	return p.Page
}

func (p *PaginationQuery) GetOffset() int {
	if p.Page == 0 {
		return 0
	}

	return (p.Page - 1) * p.Limit
}

func (p *PaginationQuery) GetOrderBy() string {
	return p.OrderBy
}

func NewPaginationQuery(c echo.Context) (*PaginationQuery, error) {
	p := &PaginationQuery{}

	if err := p.setPage(c.QueryParam("page")); err != nil {
		return nil, err
	}

	if err := p.setLimit(c.QueryParam("limit")); err != nil {
		return nil, err
	}

	p.setOrder(c.QueryParam("order"))

	return p, nil
}

func GetTotalPage(totalData int, limit int) int {
	if totalData%limit == 0 {
		return totalData / limit
	}

	return (totalData / limit) + 1
}

func CheckHasMore(totalData int, limit int, page int) bool {
	totalPage := GetTotalPage(totalData, limit)

	return page < totalPage
}
