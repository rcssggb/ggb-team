#!/usr/bin/env bash

bash ./x11expose.sh

HOST_OS=$(uname -s)
if [ $HOST_OS == "Linux" ]; then
    export DISPLAY=${DISPLAY}
elif [ $HOST_OS == "Darwin" ]; then
    export DISPLAY="host.docker.internal:0"
else
    echo "unknown host os"
    exit
fi


docker-compose up

bash ./x11disable.sh
