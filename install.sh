#!/bin/bash

# 檢測操作系統和架構
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# 定義 GitHub 存儲庫擁有者、存儲庫名稱和發布版本
OWNER="Rudy1021"
REPO="goToPost"
VERSION="v0.1.1"

# 創建關聯映射，將操作系統和架構映射到文件名
declare -A FILE_MAP
FILE_MAP=(
  ["darwin_arm64"]="goToPost_Darwin_arm64.tar.gz"
  ["darwin_x86_64"]="goToPost_Darwin_x86_64.tar.gz"
  ["linux_arm64"]="goToPost_Linux_arm64.tar.gz"
  ["linux_i386"]="goToPost_Linux_i386.tar.gz"
  ["linux_x86_64"]="goToPost_Linux_x86_64.tar.gz"
  ["windows_arm64"]="goToPost_Windows_arm64.zip"
  ["windows_i386"]="goToPost_Windows_i386.zip"
  ["windows_x86_64"]="goToPost_Windows_x86_64.zip"
)

# 確保操作系統和架構的組合存在於映射中
KEY="$OS_$ARCH"
if [ -n "${FILE_MAP[$KEY]}" ]; then
  FILE_NAME="${FILE_MAP[$KEY]}"
else
  echo "Unsupported operating system and architecture combination: $OS $ARCH"
  exit 1
fi

# 構建下載 URL
DOWNLOAD_URL="https://github.com/$OWNER/$REPO/releases/download/$VERSION/$FILE_NAME"

# 使用 curl 下載文件
curl -L -o "$FILE_NAME" "$DOWNLOAD_URL"

# 提示下載完成
echo "Downloaded file to current directory: $FILE_NAME"
