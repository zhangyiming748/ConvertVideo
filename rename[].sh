#!/bin/bash

# 指定文件夹路径和文件扩展名
folder_path="/mnt/e/video/joi"
file_extension=".mkv"

# 遍历文件夹下的所有文件
for file in "$folder_path"/*$file_extension; do
  # 使用sed命令替换文件名中的[]及其中的内容
  new_file=$(echo "$file" | sed 's/ \[.*\]//g')

  # 如果新文件名与原文件名不同，则重命名文件
  if [ "$file" != "$new_file" ]; then
    mv "$file" "$new_file"
    echo "Renamed $file to $new_file"
  fi
done
