#!/bin/bash

INSTALL_DIR=$HOME/.cptool

curl -Lo /tmp/cptool.tar.gz https://github.com/jauhararifin/cptool/archive/v0.0.1.tar.gz > /dev/null 2> /dev/null
mkdir -p $INSTALL_DIR
tar -xf /tmp/cptool.tar.gz -C $INSTALL_DIR --strip 1

if [ -z $CPTOOL_DIR ]; then
    export CPTOOL_DIR="$INSTALL_DIR"
    export PATH=$PATH:$CPTOOL_DIR
    echo "" >> $HOME/.bashrc
    echo "export CPTOOL_DIR=\"$INSTALL_DIR\"" >> $HOME/.bashrc
    echo "export PATH=\"\$PATH:\$CPTOOL_DIR\"" >> $HOME/.bashrc
fi

echo "Installed"
