#!/bin/sh

# Determine the user's operating system.
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

# Determine the user's architecture.
ARCH=$(uname -m)

# Set the release version.
VERSION="v0.1.3"

# Define the URLs for different operating systems and architectures.
if [ "$OS" = "darwin" ]; then
  URL="https://github.com/Rudy1021/goToPost/releases/download/${VERSION}/gtp"
elif [ "$OS" = "linux" ]; then
  if [ "$ARCH" = "x86_64" ]; then
    URL="https://github.com/Rudy1021/goToPost/releases/download/${VERSION}/gtp"
  elif [ "$ARCH" = "arm64" ]; then
    URL="https://github.com/Rudy1021/goToPost/releases/download/${VERSION}/gtp"
  else
    echo "Unsupported architecture: $ARCH"
    exit 1
  fi
elif [ "$OS" = "windows" ]; then
  URL="https://github.com/Rudy1021/goToPost/releases/download/${VERSION}/gtp.exe"
else
  echo "Unsupported operating system: $OS"
  exit 1
fi

# Define the installation directory.
INSTALL_DIR="$(go env GOPATH)/bin"

# Download and install the binary.
echo "Downloading and installing GoToPost..."
curl -sSfL "$URL" -o "$INSTALL_DIR/gtp"
chmod +x "$INSTALL_DIR/gtp"

echo "GoToPost installed successfully to $INSTALL_DIR"
