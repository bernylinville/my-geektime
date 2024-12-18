FROM golang:1.23 AS builder

ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /app
COPY . /app

RUN make all

FROM jrottenberg/ffmpeg

COPY --from=builder /app/mygeektime /usr/bin/mygeektime

EXPOSE 8090

ENTRYPOINT ["mygeektime"]
