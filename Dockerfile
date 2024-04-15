FROM golang:1.22.2-bookworm
# 已经测试过alpine
LABEL authors="zen"
RUN cp /etc/apt/sources.list.d/debian.sources /etc/apt/sources.list.d/debian.sources.bak
RUN sed -i 's/deb.debian.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/debian.sources
RUN apt update
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/root/go/bin
RUN mkdir -p /root/app
WORKDIR /root/app
COPY . .
RUN chmod +x /root/app/install-retry.sh
RUN /root/app/install-retry.sh ffmpeg nano mediainfo build-essential
RUN go build -o /usr/local/bin/conv main.go
RUN chmod +x /usr/local/bin/conv
WORKDIR /usr/local/bin
CMD ["conv"]
# docker build -t test:2
# docker run -d -v /f/Telegram/data test:1
# docker run -itd --name test -v /c/Users/zen/Github/ConvertVideo:/data test:2 bash