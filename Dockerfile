FROM golang:1.22.3-alpine3.20
# 已经测试过alpine
LABEL authors="zen"
# docker exec -it test ash
# docker run -dit --name test --rm -v '/media/zen/Windows 11/Users/zen/Github/ConvertVideo:/app' golang:1.22.3-alpine3.19 ash
RUN cp /etc/apk/repositories /etc/apk/repositories.bak
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk update
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/root/go/bin
RUN mkdir -p /root/app
WORKDIR /root/app
COPY . .
RUN apk add ffmpeg mediainfo build-base
RUN go build -o /usr/local/bin/conv main.go
RUN chmod +x /usr/local/bin/conv
WORKDIR /usr/local/bin
CMD ["conv"]
# docker build -t videos:latest .
# docker run -dit --rm --name vp9 -e root=/data -e to=vp9 -e level=Debug -v /media/zen/swap/pikpak/telegram:/data videos:latest