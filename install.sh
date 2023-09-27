#!/bin/sh

uname_os() {
  os=$(uname -s | tr '[:upper:]' '[:lower:]')
  case "$os" in
    cygwin_nt*) os="windows" ;;
    mingw*) os="windows" ;;
    msys_nt*) os="windows" ;;
  esac
  echo "$os"
}

# Determine the user's operating system.
OS=$(uname_os)

# Determine the user's architecture.
ARCH=$(uname -m)

# Set the release version.
TAG=$(curl -s "https://api.github.com/repos/Rudy1021/GoToPost/releases/latest" | grep -o '"tag_name": ".*"' | cut -d '"' -f 4)
VERSION="${TAG#v}"

# Define the URLs for different operating systems and architectures.
if [ "$OS" = "darwin" ]; then
  INSTALL_DIR="$(go env GOPATH)/bin"
  EXTENTION="tar.gz"
  URL="https://github.com/Rudy1021/goToPost/releases/download/v${VERSION}/goToPost_${VERSION}_darwin_amd64.${EXTENTION}"
  FILENAME="goToPost_${VERSION}_windows_amd64.${EXTENTION}"
elif [ "$OS" = "linux" ]; then
  if [ "$ARCH" = "x86_64" ]; then
    INSTALL_DIR="$(go env GOPATH)/bin"
    EXTENTION="tar.gz"
    URL="https://github.com/Rudy1021/goToPost/releases/download/v${VERSION}/goToPost_${VERSION}_linux_amd64.${EXTENTION}"
    FILENAME="goToPost_${VERSION}_windows_amd64.${EXTENTION}"
  elif [ "$ARCH" = "arm64" ]; then
    INSTALL_DIR="$(go env GOPATH)/bin"
    EXTENTION="tar.gz"
    URL="https://github.com/Rudy1021/goToPost/releases/download/v${VERSION}/goToPost_${VERSION}_linux_arm64.${EXTENTION}"
    FILENAME="goToPost_${VERSION}_windows_amd64.${EXTENTION}"
  else
    echo "Unsupported architecture: $ARCH"
    exit 1
  fi
elif [ "$OS" = "windows" ]; then
  INSTALL_DIR="$(go env GOPATH)\bin"
  EXTENTION="exe"
  URL="https://github.com/Rudy1021/goToPost/releases/download/${VERSION}/goToPost_${VERSION}_windows_amd64.${EXTENTION}"
  FILENAME="goToPost_${VERSION}_windows_amd64.${EXTENTION}"
else
  echo "Unsupported operating system: $OS"
  exit 1
fi

# Define the installation directory.


# Download and install the binary.
echo "Downloading and installing GoToPost..."
# curl -sSfL "$URL" -o "$INSTALL_DIR/gtp"

curl -LJO $URL

if [ "$OS" = "windows" ]; then
  mv $FILENAME $INSTALL_DIR
else
  tar -xzvf $FILENAME -C $INSTALL_DIR


  mv $INSTALL_DIR/GoToPost $INSTALL_DIR/gtp
  rm ./$FILENAME
fi

echo "GoToPost installed successfully to $INSTALL_DIR"