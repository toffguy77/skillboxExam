#!/usr/bin/env bash

# shellcheck disable=SC1009
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        echo "starting server..."
        ./bin/statusPage-linux-amd64
elif [[ "$OSTYPE" == "darwin"* ]]; then
        echo "starting server..."
        ./bin/statusPage-darwin-arm64
elif [[ "$OSTYPE" == "win32" ]]; then
        echo "starting server..."
        ./bin/statusPage-windows-amd64
else
        echo "exec: platform / architecture was not recognised. exiting..."
        exit 1
fi
