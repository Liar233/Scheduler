#!/bin/bash

if [ -z "$1" ]
  then
    PORT=5555
  else
    PORT=$1
fi

echo $PORT

echo "Starting simple TCP server on $PORT..."

while :; do nc -l -p $PORT; sleep 1; done
