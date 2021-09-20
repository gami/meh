FROM golang:1.16.4-alpine3.13

ENV GOOS linux
ENV GOARCH amd64
ENV PATH $PATH:/go/bin/linux_amd64
RUN \
  echo http://dl-cdn.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories && \
  apk update --no-cache && \
  apk add --virtual .build \
  bash \
  git \
  make \
  build-base \
  mysql-client \
  tzdata && \
  ln -f -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.41.1
RUN go install github.com/cosmtrek/air@latest
RUN	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.8.1
RUN	go install github.com/volatiletech/sqlboiler/v4@v4.6.0
RUN	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.6.0
RUN go get -u github.com/kyoh86/richgo

WORKDIR /app

COPY Makefile .
RUN make setup

COPY go.mod .
COPY go.sum .
RUN go mod download
