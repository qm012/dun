package dun

import (
	"errors"
	"reflect"
)

var (
	ErrDataTypeMismatch = errors.New("the list data type must be slice")
)

type PageInfo struct {
	Total    int64 `json:"total"`    // total records: gorm,mongo result dataType is int64
	List     any   `json:"list"`     // Result set
	PageSize int   `json:"pageSize"` // The number per page
	PageNum  int   `json:"pageNum"`  // Current page
	Size     int   `json:"size"`     // The number of current pages

	// Since startRow and endRow are not commonly used, here's how to use them
	// "Show the total size of startRow to endRow" in the page.

	StartRow int `json:"startRow"` // The row number of the first element of the current page in the database
	EndRow   int `json:"endRow"`   // The row number of the last element of the current page in the database
	Pages    int `json:"pages"`    // total page count

	PrePage  int `json:"prePage"`  // Previous page
	NextPage int `json:"nextPage"` // Next page

	IsFirstPage       bool  `json:"isFirstPage"`       // Whether it is the first page
	IsLastPage        bool  `json:"isLastPage"`        // Whether it is the last page
	HasPreviousPage   bool  `json:"hasPreviousPage"`   // Whether there is a previous page
	HasNextPage       bool  `json:"hasNextPage"`       // Whether there is a next page
	NavigatePages     int   `json:"navigatePages"`     // Navigation page number
	NavigatePageNums  []int `json:"navigatePageNums"`  // All navigation page numbers
	NavigateFirstPage int   `json:"navigateFirstPage"` // The first page on the navigation bar
	NavigateLastPage  int   `json:"navigateLastPage"`  // The last page on the navigation bar
}

// NewPageInfo Instantiate a paging object
func NewPageInfo(count int64, list any) *PageInfo {
	info := &PageInfo{}
	info.pageInfo(count, list, 10)
	return info
}

// SetPageSize Set the current number of items on a page
func (p *PageInfo) SetPageSize(pageNum, pageSize int) *PageInfo {
	p.PageNum = pageNum
	p.PageSize = pageSize
	if p.PageSize > 0 {
		p.Pages = int(p.Total/int64(pageSize)) + If(p.Total%int64(pageSize) == 0, 0, 1)
	} else {
		p.Pages = 0
	}

	if pageSize == 0 {
		p.StartRow = 0
		p.EndRow = 0
	} else {
		// Since the result is greater than startRow, the actual requirement is +1
		p.StartRow = If(pageNum > 0, (pageNum-1)*pageSize, 0) + 1
		tempTotal := int64(pageSize * pageNum)
		p.EndRow = int(If(tempTotal > p.Total, p.Total, tempTotal))
	}

	p.calcNavigatePageNums()
	p.calcPage()
	p.judgePageBoundary()
	return p
}

func (p *PageInfo) pageInfo(count int64, list any, navigatePages int) {
	p.Total = count
	p.List = list

	t := reflect.TypeOf(p.List)
	if t.Kind() != reflect.Slice {
		panic("the list data type must be slice")
	}
	v := reflect.ValueOf(p.List)
	p.Size = v.Len()

	p.NavigatePages = navigatePages
}

// calcNavigatePageNums Computational navigation page
func (p *PageInfo) calcNavigatePageNums() {
	if p.Pages <= p.NavigatePages {

		p.NavigatePageNums = make([]int, 0, p.Pages)
		for i := 0; i < p.Pages; i++ {
			p.NavigatePageNums = append(p.NavigatePageNums, i+1)
		}

	} else {

		p.NavigatePageNums = make([]int, 0, p.NavigatePages)
		startNum := p.PageNum - p.NavigatePages/2
		endNum := p.PageNum + p.NavigatePages/2

		if startNum < 1 {
			startNum = 1
			for i := 0; i < p.NavigatePages; i++ {
				startNum++
				p.NavigatePageNums = append(p.NavigatePageNums, startNum)
			}
		} else if endNum > p.Pages {
			endNum = p.Pages
			for i := p.NavigatePages - 1; i >= 0; i-- {
				endNum--
				p.NavigatePageNums = append(p.NavigatePageNums, endNum)
			}
		} else {
			for i := 0; i < p.NavigatePages; i++ {
				startNum++
				p.NavigatePageNums = append(p.NavigatePageNums, startNum)
			}
		}
	}
}

// calcPage Count pages
func (p *PageInfo) calcPage() {
	length := len(p.NavigatePageNums)
	if length > 0 {
		p.NavigateFirstPage = p.NavigatePageNums[0]
		p.NavigateLastPage = p.NavigatePageNums[length-1]
		if p.PageNum > 1 {
			p.PrePage = p.PageNum - 1
		}
		if p.PageNum < p.Pages {
			p.NextPage = p.PageNum + 1
		}
	}
}

// judgePageBoundary Determine page boundaries
func (p *PageInfo) judgePageBoundary() {
	p.IsFirstPage = p.PageNum == 1
	p.IsLastPage = p.PageNum == p.Pages || p.Pages == 0
	p.HasPreviousPage = p.PageNum > 1
	p.HasNextPage = p.PageNum < p.Pages
}
