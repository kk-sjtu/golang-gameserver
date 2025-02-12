#!/bin/bash

# 打印当前目录
echo "current dir: $(dirname "$0")"
cDir=$(dirname "$0")
SRC_DIR=$cDir
DST_DIR="$cDir/gen/"

# 打印 SRC_DIR 和 DST_DIR
echo "SRC_DIR: $SRC_DIR"
echo "DST_DIR: $DST_DIR"

# 创建 gen 和 doc 目录
mkdir -p "$DST_DIR"
mkdir -p "$cDir/doc"




# 使用 protoc 生成代码和文档
protoc \
  -I . \
    --go_out="$DST_DIR" \
  --validate_out="lang=go:$DST_DIR" \
  --doc_out="$cDir/doc" \
  --doc_opt=html,index.html \
  "$SRC_DIR"/proto/*.proto

echo "Code and documentation generated successfully!"