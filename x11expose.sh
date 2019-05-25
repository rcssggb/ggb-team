HOST_OS=$(uname -s)

if [ $HOST_OS == "Linux" ]; then
    export DISPLAY=${DISPLAY}
    xhost +local:root
elif [ $HOST_OS == "Darwin" ]; then
    export DISPLAY="host.docker.internal:0"
    xhost + 127.0.0.1
else
    echo "unknown host os"
    exit
fi
