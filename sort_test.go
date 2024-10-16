package ok

import "testing"

func TestSort(t *testing.T) {

	type testCase struct {
		input     Sort
		wantMysql string
		wantMongo int
	}

	testCases := []testCase{
		{input: 1, wantMysql: "ASC", wantMongo: 1},
		{input: 2, wantMysql: "DESC", wantMongo: -1},
		{input: 100, wantMysql: "ASC", wantMongo: 1},
	}

	for _, tc := range testCases {
		gotMysql := tc.input.Mysql()
		gotMongo := tc.input.Mongo()

		if gotMysql != tc.wantMysql {
			t.Errorf("Mysql sort：expected:%v, got:%v", tc.wantMysql, gotMysql)
		}
		if gotMongo != tc.wantMongo {
			t.Errorf("Mongo sort：expected:%v, got:%v", tc.wantMongo, gotMongo)
		}
	}
}
