package dun

const (
	DefaultPageNum  = 1
	DefaultPageSize = 10
)

const (
	spaceConnection = " "
)

var (
	// Whether to calculate page numbers
	// Default calculation
	calcPageNum = true
)

// DisableCalcPageNum Close the calculation page number
func DisableCalcPageNum() {
	calcPageNum = false
}

type PageSearch struct {
	PageSize  int    `json:"pageSize,string" xml:"pageSize" param:"pageSize" query:"pageSize" form:"pageSize" binding:"omitempty,min=1,max=10000"` // How many on each page
	PageNum   int    `json:"pageNum,string" xml:"pageNum" param:"pageNum" query:"pageNum" form:"pageNum" binding:"omitempty,min=1"`                // What page
	Sort      Sort   `json:"sort,string" xml:"sort" param:"sort" query:"sort" form:"sort" binding:"omitempty,oneof=1 2"`                           // sort level
	SortField string `json:"sortField" xml:"sortField" param:"sortField" query:"sortField" form:"sortField" binding:"omitempty"`                   // sort field
}

// SortByMysql mysql based sorting
// Field priority：p.SortField > customField[0]
func (p *PageSearch) SortByMysql(sortFieldMap map[string]struct{}, customFields ...string) string {
	if p.Sort == 0 {
		return ""
	}

	var customField string
	if len(customFields) > 0 {
		customField = customFields[0]
	}

	if p.SortField == "" && customField == "" {
		return ""
	}

	if p.SortField != "" {

		if sortFieldMap != nil {
			if _, ok := sortFieldMap[p.SortField]; !ok {
				return ""
			}
		}

		return p.SortField + spaceConnection + p.Sort.Mysql()
	}

	if customField != "" {

		if sortFieldMap != nil {
			if _, ok := sortFieldMap[customField]; !ok {
				return ""
			}
		}

		return customField + spaceConnection + customField
	}

	return ""
}

// Offset calculate mysql(offset)/mongo(skip)
func (p *PageSearch) Offset(totalCount int64) int {
	if totalCount <= 0 {
		return 0
	}
	// Whether to assign a default page number
	if calcPageNum && p.PageNum <= 0 {
		p.PageNum = DefaultPageNum
	}
	if p.PageSize <= 0 {
		p.PageSize = DefaultPageSize
	}
	var (
		maxPage  int64
		pageSize = int64(p.PageSize)
	)

	if totalCount%pageSize == 0 {
		maxPage = totalCount / pageSize
	} else {
		maxPage = (totalCount / pageSize) + 1
	}
	// The current page number exceeds the maximum page number assigned
	if calcPageNum && maxPage < int64(p.PageNum) {
		p.PageNum = int(maxPage)
	}

	if p.PageNum == 0 {
		return 0
	}
	return (p.PageNum - 1) * int(pageSize)
}
