#!/bin/bash
PORT=5555;

echo "Starting simple TCP server on $PORT..."

while :; do nc -l -p $PORT; sleep 1; done
