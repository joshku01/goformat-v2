# 第一層基底
FROM golang:1.11.2-alpine AS build

# 複製原始碼
COPY . /go/src/Melon
WORKDIR /go/src/Melon

# 進行編譯(名稱為：melon)
RUN go build -o melon

# 最終運行golang 的基底
FROM alpine

# 上傳檔案至 s3 需要此套件
RUN apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

COPY --from=build /go/src/Melon/melon /app/melon
COPY ./env /app/env
WORKDIR /app

RUN mkdir -p /home/log/
RUN ln -sf /dev/stdout /home/log/melon_access.log \
    && ln -sf /dev/stdout /home/log/melon_error.log

ENTRYPOINT [ "./melon" ]