HOST_OS=$(uname -s)

if [ $HOST_OS == "Linux" ]; then
    xhost -local:root
elif [ $HOST_OS == "Darwin" ]; then
    export DISPLAY="host.docker.internal:0"
    xhost - 127.0.0.1
fi
