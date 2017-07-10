# go-mysql-api

[![Build Status](https://ci.fornever.org/buildStatus/icon?job=go-mysql-api)](https://ci.fornever.org/job/go-mysql-api)

apify mysql database. based on [Echo](https://github.com/labstack/echo), [goqu](https://github.com/doug-martin/goqu), [cli](https://github.com/mkideal/cli) and [go-mysql-driver](https://github.com/go-sql-driver/mysql)

## install

```bash
go get -u -v https://github.com/Soontao/go-mysql-api
```

or download binary from [release page](https://github.com/Soontao/go-mysql-api/releases/tag/v1.0.0) !

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

## apis

if you have any web dev experience, apis will easy to understand

```golang
server.e.GET("/api/metadata", server.endpointMetadata) // metadata
server.e.POST("/api/echo", server.endpointEcho)        // echo api

server.e.GET("/api/:table", server.endpointTableGet)       // Retrive
server.e.PUT("/api/:table", server.endpointTableCreate)    // Create
server.e.DELETE("/api/:table", server.endpointTableDelete) // Remove

server.e.GET("/api/:table/:id", server.endpointTableGetSpecific)       // Retrive
server.e.DELETE("/api/:table/:id", server.endpointTableDeleteSpecific) // Delete
server.e.POST("/api/:table/:id", server.endpointTableUpdateSpecific)   // Update
```

pls use `application/json` MIME and json format in client request.

pls use json object in Create, Update, Delete method (if need payload), and there is no support for batch process now.

follow api could use **_limit**, **_skip** and **_field** query param

* GET /api/:table
* GET /api/:table/:id

follow is an example

```bash
http :1323/api/monitor_log _limit==10 _field==lid _field==mid _field==success _skip==10 -v

# GET /api/monitor_log?_limit=10&_field=lid&_field=mid&_field=success&_skip=10 HTTP/1.1

# SELECT `lid`, `mid`, `success` FROM `monitor_log` LIMIT 10 OFFSET 10
```

## any tests ?

yeah, there are some in-package tests, but not work for out-package, and based on env var

i test this project by my existed mysql schema, and it works correctly
