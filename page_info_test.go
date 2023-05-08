package dun

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type user struct {
	id   int
	name string
}

func getUsers(count int) []*user {
	users := make([]*user, 0, count)
	for i := 0; i < count; i++ {
		users = append(users, &user{
			id:   i + 1,
			name: "",
		})
	}
	return users
}

func TestNavigatePages(t *testing.T) {

	var (
		count = int64(183)
		users = getUsers(int(count))
	)

	pageInfo := NewPageInfo(count, users[:10]).SetPageSize(1, 10)
	assert.Equal(t, 1, pageInfo.PageNum)
	assert.Equal(t, 10, pageInfo.PageSize)
	assert.Equal(t, 1, pageInfo.StartRow)
	assert.Equal(t, 10, pageInfo.EndRow)
	assert.Equal(t, count, pageInfo.Total)
	assert.Equal(t, 19, pageInfo.Pages)
	assert.Equal(t, true, pageInfo.IsFirstPage)
	assert.Equal(t, false, pageInfo.IsLastPage)
	assert.Equal(t, false, pageInfo.HasPreviousPage)
	assert.Equal(t, true, pageInfo.HasNextPage)

	pageInfo = NewPageInfo(count, users[51:100]).SetPageSize(2, 50)
	assert.Equal(t, 2, pageInfo.PageNum)
	assert.Equal(t, 50, pageInfo.PageSize)
	assert.Equal(t, 51, pageInfo.StartRow)
	assert.Equal(t, 100, pageInfo.EndRow)
	assert.Equal(t, count, pageInfo.Total)
	assert.Equal(t, 4, pageInfo.Pages)
	assert.Equal(t, false, pageInfo.IsFirstPage)
	assert.Equal(t, false, pageInfo.IsLastPage)
	assert.Equal(t, true, pageInfo.HasPreviousPage)
	assert.Equal(t, true, pageInfo.HasNextPage)
}

func TestPageInfoPanic(t *testing.T) {
	type testCase struct {
		count int64
		list  any
		want  error
	}
	testCases := []testCase{
		{count: 1, list: []int{1222}, want: nil},
		{count: 100, list: true, want: ErrDataTypeMismatch},
		{count: 1, list: 2.3, want: ErrDataTypeMismatch},
		{count: 1, list: struct{}{}, want: ErrDataTypeMismatch},
		{count: 1, list: -10, want: ErrDataTypeMismatch},
	}

	for _, tc := range testCases {
		got := detectionPageInfoPanic(t, tc.count, tc.list)
		if got == nil && tc.want == nil {
			continue
		}
		if got.Error() != tc.want.Error() {
			t.Errorf("expected:%v, got:%v", tc.want, got)
		}
	}
}

func detectionPageInfoPanic(t *testing.T, count int64, list any) (err error) {

	defer func() {
		if e := recover(); e != nil {
			switch value := e.(type) {
			case string:
				err = errors.New(value)
			case error:
				err = value
			default:
				err = nil
			}
		}
	}()

	NewPageInfo(count, list).SetPageSize(1, 1)
	return err
}
