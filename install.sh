#!/bin/bash

INSTALL_DIR=$HOME/.cptool
VERSION=1.0.1
CPTOOL_NAME=cptool_${VERSION}_linux_amd64.tar.gz

curl -Lo /tmp/cptool.tar.gz https://github.com/jauhararifin/cptool/releases/download/v$VERSION/$CPTOOL_NAME > /dev/null 2> /dev/null
mkdir -p $INSTALL_DIR
tar -xf /tmp/cptool.tar.gz -C $INSTALL_DIR

if [ -z $CPTOOL_DIR ]; then
    export CPTOOL_DIR="$INSTALL_DIR"
    export PATH=$PATH:$CPTOOL_DIR
    echo "" >> $HOME/.bashrc
    echo "export CPTOOL_DIR=\"$INSTALL_DIR\"" >> $HOME/.bashrc
    echo "export PATH=\"\$PATH:\$CPTOOL_DIR\"" >> $HOME/.bashrc
fi

echo "Installed"
