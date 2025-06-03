# This dockerfile uses extends image https://hub.docker.com/sinlov/gh-conventional-kit
# VERSION 1
# Author: sinlov
# dockerfile official document https://docs.docker.com/engine/reference/builder/
# https://hub.docker.com/_/golang
FROM golang:1.23.8 AS golang-builder

ARG GO_ENV_PACKAGE_NAME=github.com/sinlov/gh-conventional-kit
ARG GO_ENV_ROOT_BUILD_BIN_NAME=gh-conventional-kit
ARG GO_ENV_ROOT_BUILD_BIN_PATH=build/${GO_ENV_ROOT_BUILD_BIN_NAME}
ARG GO_ENV_ROOT_BUILD_ENTRANCE="cmd/gh-conventional-kit/main.go"

ARG GO_PATH_SOURCE_DIR=/go/src
WORKDIR ${GO_PATH_SOURCE_DIR}

RUN mkdir -p ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME}
COPY . ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME}

# proxy golang
RUN go env -w "GOPROXY=https://goproxy.cn,direct"
RUN go env -w "GOPRIVATE='*.gitlab.com,*.gitee.com,*.sinlov.cn"

RUN cd ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME} && \
    go mod download -x

RUN  cd ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME} && \
  CGO_ENABLED=0 \
  go build \
  -a \
  -installsuffix cgo \
  -ldflags '-w -s --extldflags "-static -fpic"' \
  -tags netgo \
  -o ${GO_ENV_ROOT_BUILD_BIN_PATH} \
  ${GO_ENV_ROOT_BUILD_ENTRANCE}

# https://hub.docker.com/_/alpine
FROM alpine:3.17

# ARG DOCKER_CLI_VERSION=${DOCKER_CLI_VERSION}
ARG GO_ENV_PACKAGE_NAME=github.com/sinlov/gh-conventional-kit
ARG GO_ENV_ROOT_BUILD_BIN_NAME=gh-conventional-kit
ARG GO_ENV_ROOT_BUILD_BIN_PATH=build/${GO_ENV_ROOT_BUILD_BIN_NAME}

ARG GO_PATH_SOURCE_DIR=/go/src

#RUN apk --no-cache add \
#  ca-certificates mailcap curl \
#  && rm -rf /var/cache/apk/* /tmp/*

RUN mkdir /app
WORKDIR /app

COPY --from=golang-builder ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME}/${GO_ENV_ROOT_BUILD_BIN_PATH} .
ENTRYPOINT [ "/app/gh-conventional-kit" ]
# CMD ["/app/gh-conventional-kit", "--help"]