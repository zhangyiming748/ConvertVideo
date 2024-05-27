# ConvertVideo
视频相关转码工具 尽量不依赖第三方库
# Docker 参数


docker run -dit --rm --name merge -e to=merge -e level=Debug -v /e/video/joi:/data video:1
docker run -dit --rm --name rotate -e to=rotate -e level=Debug -e direction=ToRight -v /e/telegram/dance/ToRight:/data videos:1
/e/telegram/dance/ToRight
docker run -e direction=ToRight|ToLeft -e root=/path/to/conv -e to=转换的编码 -e level=Debug -v /path/to/src:/data

docker run -dit --rm --name vp9 --cpus=8 --memory=8192M -e to=vp9 -e level=Debug -v /e/pikpak/avs-museum:/data videos:latest
docker run -dit --rm --name vp9 -e to=vp9 -e level=Debug -v /e/pikpak/avs-museum:/data videos:latest
docker run -dit --name vp9 --cpus=8 --memory=8192M -e to=vp9 -e level=Debug -v /e/pikpak/avs-museum:/data videos:latest bash
/e/telegram/dance/ToRight
# 环境变量
+ `$root` 设置一个不同的视频存放目录
+ `$level` 日志的输出等级 可选参数 `Debug` `Info` `Warn` `Err`
+ `$to` 指定执行的任务种类 可选参数 `vp9` `rotate` `merge`
+ `$direction` 当任务选择为`rotate`时用来指定旋转方向 可选参数 `ToRight``ToLeft`
# 音频编码器
ogg -> libvorbis
opus -> libopus

