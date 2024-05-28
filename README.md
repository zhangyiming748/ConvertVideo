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
# ffmpeg运行时如何不限制cpu和内存的使用量
要在运行FFmpeg时不限制CPU和内存的使用量，可以通过设置一些参数来控制其行为。以下是一些常用的参数：

1. **-threads 参数**：可以使用 `-threads` 参数来指定要使用的线程数。例如，`-threads 0` 表示使用系统上所有可用的CPU核心。

2. **-cpu-used 参数**：可以使用 `-cpu-used` 参数来控制CPU的使用量。较低的值会减少CPU的使用，但可能会导致性能下降。例如，`-cpu-used 4` 表示使用较少的CPU资源。

3. **-preset 参数**：可以使用 `-preset` 参数来设置编码速度和质量之间的权衡。较慢的预设通常会使用更多的CPU和内存资源。例如，`-preset ultrafast` 表示使用最快的预设。

4. **-threads 参数**：可以使用 `-threads` 参数来指定要使用的线程数。例如，`-threads 0` 表示使用系统上所有可用的CPU核心。

请注意，更改这些参数可能会影响FFmpeg的性能和输出质量，因此建议根据您的需求进行调整。

# 我的电脑上有16个cpu核心 我想在运行的时候使用大约12个核心 ffmpeg的-cpu-used 参数应该设置为多少
如果您的电脑有16个CPU核心，并且希望在运行FFmpeg时使用大约12个核心，您可以尝试将`-cpu-used`参数设置为一个中等的值，以便充分利用可用的CPU资源。由于`-cpu-used`参数的值范围通常在0到16之间，您可以尝试将其设置为8或9，以便使用大约一半的CPU核心。

以下是一个示例命令，演示如何在FFmpeg中使用`-cpu-used`参数来使用大约12个CPU核心：

```bash
ffmpeg -i input.mp4 -c:v libx264 -cpu-used 8 output.mp4
```

您可以根据实际情况调整`-cpu-used`参数的值，以便在运行FFmpeg时充分利用您的电脑上的CPU资源。
# ffmpeg的-preset 参数可选什么值 分别代表什么
在FFmpeg中，`-preset`参数用于设置编码速度和质量之间的权衡。不同的预设值会影响编码速度和输出质量。以下是一些常用的`-preset`参数及其含义：

1. **ultrafast**：最快的预设，编码速度最快，但输出质量较低。

2. **superfast**：次快的预设，编码速度快，输出质量略有提升。

3. **veryfast**：快速预设，编码速度较快，输出质量适中。

4. **faster**：更快预设，编码速度更快，输出质量略有提升。

5. **fast**：快速预设，编码速度快，输出质量较高。

6. **medium**：中等预设，平衡编码速度和输出质量。

7. **slow**：慢速预设，编码速度较慢，输出质量高。

8. **slower**：更慢预设，编码速度更慢，输出质量更高。

9. **veryslow**：最慢预设，编码速度最慢，输出质量最高。

您可以根据需要选择适合您的情况的预设值，以在编码过程中平衡速度和质量。

# ffmpeg中libvps-vp9的-tile-columns  参数
在 FFmpeg 中，`-tile-columns` 参数用于设置 libvpx-vp9 编码器的 tile 列数。Tile 是一种并行处理单元，可以将视频帧分割成多个 tile 进行并行编码，从而提高编码效率和速度。

您可以使用 `-tile-columns` 参数来指定 libvpx-vp9 编码器使用的 tile 列数。例如，如果您希望设置 tile 列数为 4，您可以这样使用 `-tile-columns 4` 参数：

```bash
ffmpeg -i input.mp4 -c:v libvpx-vp9 -tile-columns 4 output.webm
```

通过设置合适的 `-tile-columns` 参数值，您可以根据您的系统配置和需求来优化 libvpx-vp9 编码器的性能。请注意，设置较大的 tile 列数可能会增加编码器的复杂性，但也有可能提高编码速度和效率。

如果您需要进一步的帮助或有其他问题，请提供更多关于您的需求和环境的信息，这样我可以为您提供更具体的指导。
# ffmpeg中libvps-vp9的-frame-parallel 参数
在 FFmpeg 中，`-frame-parallel` 参数用于设置 libvpx-vp9 编码器的帧并行模式。帧并行是一种并行处理技术，可以同时处理多个帧，从而提高编码效率和速度。

您可以使用 `-frame-parallel` 参数来启用或禁用 libvpx-vp9 编码器的帧并行模式。例如，如果您希望启用帧并行模式，您可以这样使用 `-frame-parallel 1` 参数：

```bash
ffmpeg -i input.mp4 -c:v libvpx-vp9 -frame-parallel 1 output.webm
```

通过设置 `-frame-parallel` 参数，您可以根据您的系统配置和需求来优化 libvpx-vp9 编码器的性能。启用帧并行模式可以提高编码速度，特别是在多核 CPU 系统上。

请注意，启用帧并行模式可能会增加编码器的复杂性，但通常会提高编码效率。如果您需要进一步的帮助或有其他问题，请提供更多关于您的需求和环境的信息，这样我可以为您提供更具体的指导。
# ffmpeg中libvps-vp9的-tile-columns  参数取值范围
在 FFmpeg 中，对于 libvpx-vp9 编码器的 `-tile-columns` 参数，其取值范围通常是 0 到 6。这意味着您可以将 `-tile-columns` 参数设置为 0 到 6 之间的整数值来指定 libvpx-vp9 编码器使用的 tile 列数。

具体来说，取值范围如下：
- `-tile-columns 0`：禁用 tile 列分割，即不使用 tile 并行编码。
- `-tile-columns 1`：使用 1 列 tile 进行并行编码。
- `-tile-columns 2`：使用 2 列 tile 进行并行编码。
- 依此类推，最大值为 `-tile-columns 6`，表示使用 6 列 tile 进行并行编码。

通过设置合适的 `-tile-columns` 参数值，您可以根据您的系统配置和需求来优化 libvpx-vp9 编码器的性能。请注意，选择合适的 tile 列数可以提高编码速度和效率，但也可能会增加编码器的复杂性。

如果您需要进一步的帮助或有其他问题，请提供更多关于您的需求和环境的信息，这样我可以为您提供更具体的指导。