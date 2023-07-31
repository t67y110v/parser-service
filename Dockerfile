FROM golang:latest

WORKDIR /go/src/api-app


# Copy dependency locks so we can cache.
COPY ./ ./

RUN go build -o server cmd/server/main.go   


CMD [ "./server" ]