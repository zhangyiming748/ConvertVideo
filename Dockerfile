FROM golang:1.22.3-alpine3.20
# 已经测试过alpine
LABEL authors="zen"
# docker exec -it test ash
# docker run -dit --name test --rm -v '/media/zen/Windows 11/Users/zen/Github/ConvertVideo:/app' golang:1.22.3-alpine3.19 ash
RUN cp /etc/apk/repositories /etc/apk/repositories.bak
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk update
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/root/go/bin
RUN mkdir -p /root/app
WORKDIR /root/app
COPY . .
RUN apk add ffmpeg opencore-amr-dev libvorbis-dev mediainfo build-base xz wget ca-certificates dialog make cmake alpine-sdk gcc nasm yasm aom-dev libvpx-dev libwebp-dev x264-dev x265-dev dav1d-dev xvidcore-dev fdk-aac-dev opencore-amr-dev libvorbis-dev
RUN go build -o /usr/local/bin/conv main.go
RUN chmod +x /usr/local/bin/conv
WORKDIR /root/app
#RUN rm -rf ffmpeg-7.0.1.tar.xz ffmpeg-7.0.1
CMD ["conv"]
# docker build -t videos:latest .
# docker run -dit --rm --name h265 -e root=/data -e to=h265 -v /home/zen/Videos:/data zhangyiming748/convertvideo:latest bash