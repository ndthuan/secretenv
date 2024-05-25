#!/bin/bash
goVers="1.18 1.19 1.20 1.21 1.22"
for ver in $goVers; do
    echo "Running tests for Go $ver"
    
    GO_IMAGE="golang:$ver-alpine" docker compose run --rm test
done
