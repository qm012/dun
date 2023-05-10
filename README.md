# Dun web develop tool library

[//]: # ([![Build Status]&#40;https://github.com//qm012/dun/workflows/Run%20Tests/badge.svg?branch=main&#41;]&#40;https://github.com/qm012/dun/actions?query=branch%3Amian&#41;)

[//]: # ([![codecov]&#40;https://codecov.io/gh//qm012/dun/branch/main/graph/badge.svg&#41;]&#40;https://codecov.io/gh/qm012/dun&#41;)
[![GoDoc](https://pkg.go.dev/badge/github.com/qm012/dun?status.svg)](https://pkg.go.dev/github.com/qm012/dun?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/qm012/dun/-/badge.svg)](https://sourcegraph.com/github.com/qm012/dun?badge)
[![Release](https://img.shields.io/github/release/qm012/dun.svg?style=flat-square)](https://github.com/qm012/dun/releases)

go web develop tool library,includes pagination, middleware, pageSearch, response and other functions

## Getting started

### Getting dun

```sh
$ go get -u github.com/qm012/dun
```

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```go
import "github.com/qm012/dun"
```

### Examples reference

#### PageInfo

```go
// for gorm paging
// Part of the code is omitted
func gormPaging() (*dun.PageInfo, err){
    var (
        count int64
        userinfo Userinfo
    )
    err := *gorm.DB.Model(&userinfo).Select("id").Where(cmd, values...).Count(&count).Error
    if err != nil {
        return nil, err
    }
	
    userinfoList := make([]*Userinfo, 0, req.PageSize)
    err = *gorm.DB.Where(cmd, values...).Order(order).Limit(req.PageSize).Offset(req.RecordNo).Find(&userinfos).Error
    if err !=nil {
        return nil, err
    }
    
    info := dun.NewPageInfo(count, userinfoList).SetPageSize(req.PageNum, req.PageSize)
    // info object can be used by the frontend
    return info, nil
}

// for mongo paging 
// Part of the code is omitted
func mongoPaging() (*dun.PageInfo, err){
    count, err := *mongo.Collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return nil, err
    }
	
    opt := options.Find().
        SetLimit(int64(req.PageSize)).
        SetSkip(int64(req.RecordNo))
    cursor,err:=*mongo.Collection.Find(ctx, filter, opt)
    if err != nil {
        return nil, err
    }
    userinfoList := make([]*Userinfo, 0, req.PageSize) // cursor.All(ctx, &userinfos) 
    info := dun.NewPageInfo(count, userinfoList).SetPageSize(req.PageNum, req.PageSize)
    // info object can be used by the frontend
    return info, nil
} 

```


## License

The dun web tool is open-sourced software licensed under the [Apache license](./LICENSE).

## Acknowledgments

The following project had particular influence on dun's design.

- [pagehelper/Mybatis-PageHelper](https://github.com/pagehelper/Mybatis-PageHelper) Mybatis通用分页插件