# go-mysql-api

[![Build Status](https://ci.fornever.org/buildStatus/icon?job=go-mysql-api)](https://ci.fornever.org/job/go-mysql-api)

apify mysql database. based on [Echo](https://github.com/labstack/echo), [goqu](https://github.com/doug-martin/goqu), [cli](https://github.com/mkideal/cli) and [go-mysql-driver](https://github.com/go-sql-driver/mysql)

## install

```bash
go get -u -v https://github.com/Soontao/go-mysql-api
```

or download binary from [release page](https://github.com/Soontao/go-mysql-api/releases) !

or with [docker container](https://hub.docker.com/r/theosun/go-mysql-api/)

## command args

you could run go-mysql-api from cli directly

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

more information about connection str, you could see [here](https://github.com/go-sql-driver/mysql#examples)

## docker

if you use docker, set environment vars to setup your server

```bash
docker run -d --restart=always --link mariadb:mysql -p 1323:1323 -e API_CONN_STR='user:pass@tcp(domain:port)/db' -e API_HOST_LS=':1323' theosun/go-mysql-api:latest
```

use correct link, or config with public mysql database

## apis

if you have any web dev experience, apis will easy to understand

```golang
server.e.GET("/api/metadata", server.endpointMetadata).Name = "Database Metadata"
server.e.POST("/api/echo", server.endpointEcho).Name = "Echo API"
server.e.GET("/api/endpoints", server.endpointServerEndpoints).Name = "Server Endpoints"
server.e.GET("/api/updatemetadata", server.endpointUpdateMetadata).Name = "Update DB Metadata"

server.e.GET("/api/:table", server.endpointTableGet).Name = "Retrive Some Records"
server.e.PUT("/api/:table", server.endpointTableCreate).Name = "Create Single Record"
server.e.DELETE("/api/:table", server.endpointTableDelete).Name = "Remove Some Records"

server.e.GET("/api/:table/:id", server.endpointTableGetSpecific).Name = "Retrive Record By ID"
server.e.DELETE("/api/:table/:id", server.endpointTableDeleteSpecific).Name = "Delete Record By ID"
server.e.POST("/api/:table/:id", server.endpointTableUpdateSpecific).Name = "Update Record By ID"

server.e.PUT("/api/batch/:table", server.endpointBatchCreate).Name = "Batch Create Record"
```

pls use `application/json` MIME and json format in client request.

pls use json object(`{object}`) in Create, Update, Delete method (if need payload), and there is no support for batch process now.

## Get DB Metadata

You could use **GET** `/api/metadata` get database metadata, or with `?simple=true` get simple metadata

```json

{
    "[BASE TABLE] (1 rows) sessions": [
        "session_id varchar(128)  NullAble(NO) ''",
        "expires int(11) unsigned  NullAble(NO) ''",
        "data text  NullAble(YES) ''"
    ],
    "[BASE TABLE] (111802 rows) monitor_log": [
        "lid int(11)  NullAble(NO) 'Log ID'",
        "mid int(11)  NullAble(NO) 'Monitor ID'",
        "success tinyint(1)  NullAble(NO) 'Is Success'",
        "duration int(5)  NullAble(NO) 'Request duration'",
        "create_at datetime current_timestamp() NullAble(NO) ''"
    ],
    "[BASE TABLE] (2 rows) user": [
        "uid int(11)  NullAble(NO) 'User ID'",
        "uname varchar(128)  NullAble(NO) 'User Name/Email'",
        "utoken varchar(32)  NullAble(NO) 'User Token'",
        "create_at datetime current_timestamp() NullAble(NO) ''"
    ],
    "[BASE TABLE] (3 rows) monitor": [
        "mid int(11)  NullAble(NO) 'Monitor ID'",
        "uid int(11)  NullAble(NO) 'User ID'",
        "type enum('TCP','HTTP')  NullAble(NO) 'Monitor Type'",
        "target varchar(255)  NullAble(NO) 'Monitor check target'",
        "create_at datetime current_timestamp() NullAble(YES) ''"
    ]
}

```

## operate record

* use **PUT `/api/user`** method to create new user record

body

```json

{
    "uname":"fjdasl@fjdksalf",
    "utoken":"atoken"
}

```

response

```json

{
    "status": 200,
    "message": "create record",
    "data": {
        "lastInsertID": 31,
        "rowesAffected": 1
    }
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

```json


{
    "status": 200,
    "message": "delete record by id",
    "data": {
        "lastInsertID": 0,
        "rowesAffected": 1
    }
}

```

## Advance query

query apis could use **_limit**, **_skip**, **_field**, **_where**, **_link** query param

* filter fields

you could use `_field` choose which fields you need

```bash

http :1323/api/monitor_log _limit==10 _field==lid _field==mid _field==success _skip==10 -v

# GET /api/monitor_log?_limit=10&_field=lid&_field=mid&_field=success&_skip=10 HTTP/1.1

```

```sql

SELECT `lid`, `mid`, `success`
  FROM `monitor_log`
  LIMIT 10
  OFFSET 10

```

* auto join and powerful query

You could use `in`, `notIn`, `like`, `is`, `neq`, `isNot` and `eq` in `_where` param

```bash

# GET /api/monitor?_link=user&_link=monitor_log&_limit=100&_where='user.uid'.in(11,22)&_where='monitor_log.success'.eq(false)

```

```sql

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

## any tests ?

yeah, there are some in-package tests, but not work for out-package, and based on env var

I test this project by my existed mysql schema, and it works correctly
