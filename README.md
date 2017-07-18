# go-mysql-api

[![Build Status](https://ci.fornever.org/buildStatus/icon?job=go-mysql-api)](https://ci.fornever.org/job/go-mysql-api)

apify mysql database. based on [Echo](https://github.com/labstack/echo), [goqu](https://github.com/doug-martin/goqu), [cli](https://github.com/mkideal/cli) and [go-mysql-driver](https://github.com/go-sql-driver/mysql)

## install

```bash
go get -u -v https://github.com/Soontao/go-mysql-api
```

or download binary from [release page](https://github.com/Soontao/go-mysql-api/releases/tag/v1.0.0) !

or with [docker container](https://hub.docker.com/r/theosun/go-mysql-api/)

## command args

you could run go-mysql-api from cli directly

```bash
go-mysql-api --help
Options:

  -h, --help                     display help information
  -c, --*conn[=$API_CONN_STR]   *mysql connection str
  -l, --*listen[=$API_HOST_LS]  *listen host and port
```

## start

you could start will cli args, but env var also works

```bash
go-mysql-api -c "user:pass@tcp(domain:port)/db" -l "0.0.0.0:1323"
[INFO] 2017-07-07T10:28:42.431074+08:00 server start at 0.0.0.0:1323
```

more information about connection str, you could see [here](https://github.com/go-sql-driver/mysql#examples)


## docker

if you use docker, set environment vars to setup your server

```bash
docker run -d --restart=always -p 1323:1323 -e API_CONN_STR='user:pass@tcp(domain:port)/db' -e API_HOST_LS=':1323' theosun/go-mysql-api:v1
```

## apis

if you have any web dev experience, apis will easy to understand

```golang
server.e.GET("/api/metadata", server.endpointMetadata)             // metadata
server.e.POST("/api/echo", server.endpointEcho)                    // echo api
server.e.Any("/api/updatemetadata", server.endpointUpdateMetadata) // update metadata

server.e.GET("/api/:table", server.endpointTableGet)       // Retrive
server.e.PUT("/api/:table", server.endpointTableCreate)    // Create
server.e.DELETE("/api/:table", server.endpointTableDelete) // Remove

server.e.GET("/api/:table/:id", server.endpointTableGetSpecific)       // Retrive
server.e.DELETE("/api/:table/:id", server.endpointTableDeleteSpecific) // Delete
server.e.POST("/api/:table/:id", server.endpointTableUpdateSpecific)   // Update
```

pls use `application/json` MIME and json format in client request.

pls use json object(`{object}`) in Create, Update, Delete method (if need payload), and there is no support for batch process now.

## Get DB Metadata

You could use `/api/metadata` get database metadata, or with simple query param get simple metadata

```bash

# GET /api/metadata?simple=true

```

```json

{
    "monitor": {
        "create_at": "datetime",
        "mid": "int(11)",
        "target": "varchar(255)",
        "type": "enum('TCP','HTTP')",
        "uid": "int(11)"
    },
    "monitor_log": {
        "create_at": "datetime",
        "duration": "int(5)",
        "lid": "int(11)",
        "mid": "int(11)",
        "success": "tinyint(1)"
    },
    "sessions": {
        "data": "text",
        "expires": "int(11) unsigned",
        "session_id": "varchar(128)"
    },
    "user": {
        "create_at": "datetime",
        "uid": "int(11)",
        "uname": "varchar(128)",
        "utoken": "varchar(32)"
    }
}

```

## operate record

use **PUT** method to create a record

```bash

# POST /api/user

```

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

and use **GET `/api/user/31`** to get our created record

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

and use **DELETE `/api/user/31`** to delete the record, (body is not needed)

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
