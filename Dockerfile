FROM golang:1.18-alpine3.15 AS builder

COPY . /github.com/t67y110v/parser-service/
WORKDIR /github.com/t67y110v/parser-service/

RUN go mod download

RUN go build -o ./bin/server cmd/apiserver/main.go


FROM alpine:latest 

WORKDIR /root/

COPY --from=0 /github.com/t67y110v/parser-service/bin/server .
COPY --from=0 /github.com/t67y110v/parser-service/configs configs/ 

EXPOSE 80

CMD ["./server"]


#docker build -t server-api:v0.1 .
#docker run --name server -p 80:80 --env-file configs/apiserver.toml server-api:v0.1