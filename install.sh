#!/bin/bash

echo "Installing CPTool"

GITHUB_REPO=jauhararifin/cptool
RELEASES_URL=https://api.github.com/repos/$GITHUB_REPO/releases
INSTALL_DIR=$HOME/.cptool

VERSION=$(curl -L $RELEASES_URL 2> /dev/null | grep -sEoa "\"tag_name\": \"v[0-9]+\.[0-9]+\.[0-9]+\"" | head -n 1 | grep -sEoa [0-9]+\.[0-9]+\.[0-9]+)
echo "Using version v$VERSION"

OS=linux
ARCH=`uname -m`
if [[ $ARCH == 'x86_64' ]]; then
  ARCH=amd64
else
  ARCH=386
fi
CPTOOL_NAME=cptool_${VERSION}_${OS}_${ARCH}.tar.gz
echo "Using OS=${OS} ARCH=${ARCH}"

echo "Downloading binary"
curl -Lo /tmp/cptool.tar.gz https://github.com/$GITHUB_REPO/releases/download/v$VERSION/$CPTOOL_NAME > /dev/null 2> /dev/null

echo "Verifying downloaded binary"
BINARY_CHECKSUM=$(curl -L https://github.com/$GITHUB_REPO/releases/download/v$VERSION/cptool_checksums.txt 2> /dev/null | grep $CPTOOL_NAME | head -c 64)
DOWNLOADED_CHECKSUM=$(sha256sum /tmp/cptool.tar.gz | head -c 64)
if [[ "$BINARY_CHECKSUM" != "$DOWNLOADED_CHECKSUM" ]]; then
    echo "Downloaded binary corrupted"
    exit -1
fi

echo "Installing binary"
mkdir -p $INSTALL_DIR
tar -xf /tmp/cptool.tar.gz -C $INSTALL_DIR

echo "Setting up environment variables"
if [ -z $CPTOOL_DIR ]; then
    export CPTOOL_DIR="$INSTALL_DIR"
    export PATH=$PATH:$CPTOOL_DIR
    echo "" >> $HOME/.bashrc
    echo "export CPTOOL_DIR=\"$INSTALL_DIR\"" >> $HOME/.bashrc
    echo "export PATH=\"\$PATH:\$CPTOOL_DIR\"" >> $HOME/.bashrc
fi

echo "Installed"
