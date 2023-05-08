package dun

import (
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	type testCase struct {
		expr                    bool
		firstValue, secondValue any
		want                    any
	}
	testCases := []testCase{
		{expr: false, firstValue: 3, secondValue: true, want: true},
		{expr: true, firstValue: "22", secondValue: 12, want: "22"},
		{expr: true, firstValue: 3232.3333, secondValue: false, want: 3232.3333},
	}

	for _, tc := range testCases {
		got := If(tc.expr, tc.firstValue, tc.secondValue)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected:%v, got:%v", tc.want, got)
		}
	}
}
