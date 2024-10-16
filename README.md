# ok web develop tool library

[//]: # ([![Build Status]&#40;https://github.com//qm012/ok/workflows/Run%20Tests/badge.svg?branch=main&#41;]&#40;https://github.com/qm012/ok/actions?query=branch%3Amian&#41;)

[//]: # ([![codecov]&#40;https://codecov.io/gh//qm012/ok/branch/main/graph/badge.svg&#41;]&#40;https://codecov.io/gh/qm012/ok&#41;)
[![GoDoc](https://pkg.go.dev/badge/github.com/qm012/ok?status.svg)](https://pkg.go.dev/github.com/qm012/ok?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/qm012/ok/-/badge.svg)](https://sourcegraph.com/github.com/qm012/ok?badge)
[![Release](https://img.shields.io/github/release/qm012/ok.svg?style=flat-square)](https://github.com/qm012/ok/releases)

go gin web develop tool library,includes pagination, middleware, pageSearch, response and other functions

## Getting started

### Getting ok

```sh
$ go get -u github.com/qm012/ok
```

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```go
import "github.com/qm012/ok"
```

### Examples reference

#### Base info

```go
type SearchUserinfoReq struct {
	Query string `json:"query" binding:"required,max=1000"`
	ok.PageSearch
}
```

#### PageInfo

```go
// for gorm paging
// Part of the code is omitted
func gormPaging(req *SerchUserinfoReq) (*ok.PageInfo, err){
    var (
        count int64
        userinfo Userinfo
    )
    err := *gorm.DB.Model(&userinfo).Select("id").Where(cmd, values...).Count(&count).Error
    if err != nil {
        return nil, err
    }
	
    userinfoList := make([]*Userinfo, 0, req.PageSize)
    err = *gorm.DB.Where(cmd, values...).Order(req.SortByMysql(nil, "id")).Limit(req.PageSize).Offset(req.Offset(count)).Find(&userinfos).Error
    if err !=nil {
        return nil, err
    }
    
    info := ok.NewPageInfo(count, userinfoList).SetPageSize(req.PageNum, req.PageSize)
    // info object can be used by the frontend
    return info, nil
}

// for mongo paging 
// Part of the code is omitted
func mongoPaging(req *SerchUserinfoReq) (*ok.PageInfo, err){
    count, err := *mongo.Collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return nil, err
    }
	
    opt := options.Find().
        SetLimit(int64(req.PageSize)).
        SetSkip(int64(req.Offset(count)))
    cursor,err:=*mongo.Collection.Find(ctx, filter, opt)
    if err != nil {
        return nil, err
    }
    userinfoList := make([]*Userinfo, 0, req.PageSize) // cursor.All(ctx, &userinfos) 
    info := ok.NewPageInfo(count, userinfoList).SetPageSize(req.PageNum, req.PageSize)
    // info object can be used by the frontend
    return info, nil
} 

```
#### PageSearch

##### P.S.：The value of `offset` depends on whether the current page exceeds the maximum number of pages. By default, the maximum number of pages is the main number. You can also call `ok.DisableCalcPageNum()` to cancel the calculation，`ok.DisableCalcPageNum()` global valid
```go
func GetUserinfoService(req *SearchUserinfoReq)  {
    // get request object data
	sortFieldMap := map[string]struct{}{
	    "id":  {},
	    "name":{},
		"sort":{},
    }
    // Prevents sql injection. Properties are valid when they exist in sortFieldMap 
	`example1：`*gorm.DB.Where(cmd, values...).Order(req.SortByMysql(sortFieldMap))
    // req.SortField level gt customField
	`example2：`*gorm.DB.Where(cmd, values...).Order(req.SortByMysql(sortFieldMap, "id"))

    totalCount := 290 // total records，data source: Select/Find mysql/mongo data
    offset:=req.Offset(totalCount)
    `example mysql：`*gorm.DB.Where(cmd, values...).Limit(req.PageSize).Offset(offset)
    `example mongo：`opt := options.Find().
                            SetLimit(int64(req.PageSize)).
                            SetSkip(int64(offset))
}

```

## License

The ok web tool is open-sourced software licensed under the [Apache license](./LICENSE).

## Acknowledgments

The following project had particular influence on ok's design.

- [pagehelper/Mybatis-PageHelper](https://github.com/pagehelper/Mybatis-PageHelper) Mybatis通用分页插件
