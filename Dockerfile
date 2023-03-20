FROM golang:1.18 as builder


WORKDIR /go/bin

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/sample  cmd/sample/main.go


FROM alpine:3.13

LABEL MAINTAINER="Colynn Liu <colynn.liu@gmail.com>"

WORKDIR /go/bin

COPY --from=builder /go/bin/sample  /go/bin/

EXPOSE 5080

ENTRYPOINT ["./sample"]
