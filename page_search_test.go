package dun

import "testing"

func TestPageSearch(t *testing.T) {

	type testCase struct {
		input         *PageSearch
		isCalcPageNum bool
		totalCount    int64
		sortFieldMap  map[string]struct{}
		customFields  string
		wantOffset    int
		wantMysqlSort string
	}

	testCases := []*testCase{
		{input: &PageSearch{PageSize: 100, PageNum: 1, Sort: 1, SortField: ""}, isCalcPageNum: true, totalCount: 323, sortFieldMap: map[string]struct{}{"_id": {}}, wantOffset: 0, wantMysqlSort: ""},
		{input: &PageSearch{PageSize: 100, PageNum: 2, Sort: 3, SortField: "abc"}, isCalcPageNum: true, totalCount: 323, sortFieldMap: map[string]struct{}{"_id": {}}, wantOffset: 100, wantMysqlSort: ""},
		{input: &PageSearch{PageSize: 100, PageNum: 2, Sort: 3, SortField: "abc"}, isCalcPageNum: true, totalCount: 323, sortFieldMap: map[string]struct{}{"_id": {}, "abc": {}}, wantOffset: 100, wantMysqlSort: "abc ASC"},
		{input: &PageSearch{PageSize: 100, PageNum: 5555, Sort: 3, SortField: "abc"}, isCalcPageNum: false, totalCount: 1000, sortFieldMap: map[string]struct{}{"_id": {}, "abc": {}}, wantOffset: 555400, wantMysqlSort: "abc ASC"},
		{input: &PageSearch{PageSize: 22, PageNum: 88, Sort: 1, SortField: "number"}, isCalcPageNum: true, totalCount: 98989, sortFieldMap: map[string]struct{}{"_id": {}, "abc": {}}, wantOffset: 1914, wantMysqlSort: ""},
		{input: &PageSearch{PageSize: 22, PageNum: 88, Sort: 1, SortField: "number"}, isCalcPageNum: true, totalCount: 100, sortFieldMap: map[string]struct{}{"_id": {}, "number": {}}, wantOffset: 88, wantMysqlSort: "number ASC"},
	}

	for _, tc := range testCases {
		if !tc.isCalcPageNum {
			DisableCalcPageNum()
		} else {
			calcPageNum = true
		}
		gotOffset := tc.input.Offset(tc.totalCount)
		if gotOffset != tc.wantOffset {
			t.Errorf("offset ：expected:%v, got:%v", tc.wantOffset, gotOffset)
		}
		gotMysqlSort := tc.input.SortByMysql(tc.sortFieldMap, tc.customFields)
		if gotMysqlSort != tc.wantMysqlSort {
			t.Errorf("mysql sort ：expected:%v, got:%v", tc.wantMysqlSort, gotMysqlSort)
		}
	}

}
