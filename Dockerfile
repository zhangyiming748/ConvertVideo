FROM alpine:3.19.1

LABEL authors="zen"
RUN cp /etc/apk/repositories /etc/apk/repositories.bak
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add ffmpeg nano mediainfo go
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/root/go/bin
RUN mkdir -p /root/go/src
WORKDIR /root/go/src
COPY . .
RUN go build -o /usr/local/bin/conv main.go
RUN chmod +x /usr/local/bin/conv
WORKDIR /usr/local/bin
CMD ["conv"]
# docker build -t test:2