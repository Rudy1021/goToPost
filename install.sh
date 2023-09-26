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
  URL="https://github.com/Rudy1021/goToPost/releases/download/v${VERSION}/goToPost_${VERSION}_darwin_amd64.tar.gz"
elif [ "$OS" = "linux" ]; then
  if [ "$ARCH" = "x86_64" ]; then
    URL="https://github.com/Rudy1021/goToPost/releases/download/v${VERSION}/goToPost_${VERSION}_linux_amd64.tar.gz"
  elif [ "$ARCH" = "arm64" ]; then
    URL="https://github.com/Rudy1021/goToPost/releases/download/v${VERSION}/goToPost_${VERSION}_linux_arm64.tar.gz"
  else
    echo "Unsupported architecture: $ARCH"
    exit 1
  fi
elif [ "$OS" = "windows" ]; then
  URL="https://github.com/Rudy1021/goToPost/releases/download/${VERSION}/goToPost_${VERSION}_windows_amd64.exe"
else
  echo "Unsupported operating system: $OS"
  exit 1
fi

# Define the installation directory.
INSTALL_DIR="$(go env GOPATH)/bin"

# Download and install the binary.
echo "Downloading and installing GoToPost..."
# curl -sSfL "$URL" -o "$INSTALL_DIR/gtp"

curl -LJO $URL
tar -xzvf goToPost_0.1.3_darwin_amd64.tar.gz -C $INSTALL_DIR
mv $INSTALL_DIR/GoToPost $INSTALL_DIR/gtp
rm ./goToPost_0.1.3_darwin_amd64.tar.gz
echo "GoToPost installed successfully to $INSTALL_DIR"