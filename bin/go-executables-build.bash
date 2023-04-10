#!/usr/bin/env bash

package_name=$1
src_directory=$2
dest_directory=$3

if [[ -z "$package_name" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

platforms=("windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64" "linux/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $dest_directory/$output_name $src_directory
    if [ $? -ne 0 ]; then
           echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
    chmod +x "$dest_directory"/$output_name
done
