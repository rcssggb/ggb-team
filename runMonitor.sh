#!/usr/bin/env bash

HOST_OS=$(uname -s)
if [ $HOST_OS == "Linux" ]; then
    echo "linux"
elif [ $HOST_OS == "Darwin" ]; then
    echo "mac"
fi
