FROM golang:1.16.4-alpine3.13 AS build

RUN apk --update upgrade && \
    apk --no-cache add ca-certificates curl bash git openssh gcc curl musl-dev tzdata make && \
    update-ca-certificates && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del --purge tzdata

WORKDIR /app
ADD . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/linux/out ./cmd/server/main.go

FROM alpine:latest
RUN apk --update upgrade && \
    apk --no-cache add ca-certificates tzdata && \
    update-ca-certificates && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del --purge tzdata
WORKDIR /app
COPY --from=build /go/src/app/out/linux/out /app
CMD ["/root/app", "--port=8888", "--host=0.0.0.0"]
