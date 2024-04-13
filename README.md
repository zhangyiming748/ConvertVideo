# ConvertVideo
视频相关转码工具 尽量不依赖第三方库
# Docker 参数
docker run -e root=/path/to/conv -e to=转换的编码 -e level=Debug -v /path/to/src:/data
docker run -e direction=ToRight|ToLeft -e root=/path/to/conv -e to=转换的编码 -e level=Debug -v /path/to/src:/data