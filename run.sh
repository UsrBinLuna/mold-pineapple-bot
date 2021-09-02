#!/usr/bin/env bash
if [ ! -f .env ]; then echo "File .env not found, aborting" && exit 1; fi
source .env
echo "Sourced .env"
if [ ! -f mold-go ]; then echo "Binary 'mold-go' not found, please run 'go build'"&& exit 1; fi
./mold-go -t $TOKEN
