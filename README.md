# go-mysql-api

[![Build Status](https://ci.fornever.org/buildStatus/icon?job=go-mysql-api)](https://ci.fornever.org/job/go-mysql-api)

apify mysql database. provide restful api for mysql/mariadb database

Based on [Echo](https://github.com/labstack/echo), [goqu](https://github.com/doug-martin/goqu), [cli](https://github.com/mkideal/cli) and [go-mysql-driver](https://github.com/go-sql-driver/mysql)

## install

get go-mysql-api with go env

```bash
go get -u -v https://github.com/Soontao/go-mysql-api
```

or download binary from [release page](https://github.com/Soontao/go-mysql-api/releases) !

or see [docker image](https://hub.docker.com/r/theosun/go-mysql-api/) instance

or download latest build from [here](https://download.fornever.org/go-mysql-api/latest/)

## start server

You could run go-mysql-api binary from cli directly

```bash
go-mysql-api --help
Options:

  -h, --help                        display help information
  -c, --*conn[=$API_CONN_STR]      *mysql connection str
  -l, --*listen[=$API_HOST_LS]     *listen host and port
  -n, --noinfo[=$API_NO_USE_INFO]   dont use mysql information shcema

```

defaultly, server will retrive metadata from mysql information schema, if there are any problem, pls use `-n` option

## start

you could start will cli args, but env var also works

```bash
go-mysql-api -c "monitor:pass@tcp(mysql:3306)/monitor" -l "0.0.0.0:1323"

[INFO] 2017-07-26T15:09:48.4086821+08:00 connected to mysql with conn_str: monitor:pass@tcp(mysql:3306)/monitor
[INFO] 2017-07-26T15:09:49.7367783+08:00 retrived metadata from mysql database: monitor
[INFO] 2017-07-26T15:09:49.7367783+08:00 server start at :1323
```

more information about connection str, please see [here](https://github.com/go-sql-driver/mysql#examples)

## docker

if you use docker, set environment vars to setup your server

```bash
docker run -d --restart=always --link mariadb:mysql -p 1323:1323 -e API_CONN_STR='user:pass@tcp(domain:port)/db' -e API_HOST_LS=':1323' theosun/go-mysql-api:latest
```

please use correct connection string, or connectwith with public mysql database

## apis

if you have any web dev experience, apis will easy to understand

```golang
s.GET("/api/metadata", endpointMetadata(api)).Name = "Database Metadata"
s.POST("/api/echo", endpointEcho).Name = "Echo API"
s.GET("/api/endpoints", endpointServerEndpoints(s)).Name = "Server Endpoints"
s.GET("/api/updatemetadata", endpointUpdateMetadata(api)).Name = "Update DB Metadata"
s.GET("/api/swagger.json", endpointSwaggerJSON(api)).Name = "Swagger Infomation"
s.GET("/api/swagger-ui.html", endpointSwaggerUI).Name = "Swagger UI"

s.GET("/api/:table", endpointTableGet(api)).Name = "Retrive Some Records"
s.POST("/api/:table", endpointTableCreate(api)).Name = "Create Single Record"
s.DELETE("/api/:table", endpointTableDelete(api)).Name = "Remove Some Records"

s.GET("/api/:table/:id", endpointTableGetSpecific(api)).Name = "Retrive Record By ID"
s.DELETE("/api/:table/:id", endpointTableDeleteSpecific(api)).Name = "Delete Record By ID"
s.PATCH("/api/:table/:id", endpointTableUpdateSpecific(api)).Name = "Update Record By ID"

s.POST("/api/batch/:table", endpointBatchCreate(api)).Name = "Batch Create Records"
```

## Swagger UI Support

The go-mysql-api support swagger.json and provide swagger.html page

Open **/api/swagger-ui.html** to see swagger documents, the interactive documention will be helpful.

And **go-mysql-api** provide the *swagger.json* at path **/api/swagger.json**

## Get DB Metadata

You could use **GET** `/api/metadata` get database metadata, or with `?simple=true` param to get simple metadata

## Operate record

* use **POST `/api/user`** method to create new user record

body

```json

{
    "uname":"fjdasl@fjdksalf",
    "utoken":"atoken"
}

```

* use **GET `/api/user/31`** to get our created record

```json

{
    "status": 200,
    "message": "get table by id",
    "data": [
        {
            "create_at": "2017-07-18 03:21:16",
            "uid": "31",
            "uname": "fjdasl@fjdksalf",
            "utoken": "atoken"
        }
    ]
}
```

* use **DELETE `/api/user/31`** to delete the record, (body is not needed)

## Advance query

query apis could use **_limit**, **_skip**, **_field**, **_fields**, **_where**, **_link** and **_search** in query param

### Whole table search, with **low query performance**

`GET /api/user?_search=outlook`

```sql
-- SQL mapping
SELECT * FROM `user`
  WHERE
    (
      (`create_at` LIKE BINARY '%outlook%')
        OR
      (`uid` LIKE BINARY '%outlook%')
        OR
      (`uname` LIKE BINARY '%outlook%')
        OR
      (`utoken` LIKE BINARY '%outlook%')
    )

```

### Auto join and custome query

You could use `in`, `notIn`, `like`, `is`, `neq`, `isNot` and `eq` in `_where` param

`GET /api/monitor?_link=user&_link=monitor_log&_limit=100&_where='user.uid'.in(11,22)&_where='monitor_log.success'.eq(false)`

```sql
-- SQL mapping
SELECT * FROM `monitor`
  INNER JOIN `user`
    ON (`user`.`uid` = `monitor`.`uid`)
  INNER JOIN `monitor_log`
    ON (`monitor_log`.`mid` = `monitor`.`mid`)
  WHERE
    (
      (`user`.`uid` IN ('11', '22'))
    AND
      (`monitor_log`.`success` = 'false')
    )
  LIMIT 100

```

**Even if go-mysql-api has already supported simple association query, we still recommend using views for complex queries**

## Some tests

yeah, there are some in-package tests, but not work for out-package, and based on env var

I test this project by my existed mysql schema, and it works correctly
