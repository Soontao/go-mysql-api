FROM ubuntu

ENV API_HOST_LS :1323

COPY go-mysql-api-linux-amd64 /go-mysql-api

EXPOSE 1323

CMD ["/go-mysql-api"]
