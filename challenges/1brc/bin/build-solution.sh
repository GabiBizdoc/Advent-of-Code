#!/bin/bash

platforms=("linux/amd64" "linux/386" "linux/arm64" "windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64")

program_name="solution"

> "$program_name/checksums.txt"
for platform in "${platforms[@]}"; do

    IFS='/' read -ra split <<< "$platform"
    GOOS="${split[0]}"
    GOARCH="${split[1]}"

    output_name="$program_name/$program_name-$GOOS-$GOARCH"

    if [[ "windows" == $GOOS ]]; then
     output_name="$output_name.exe"
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o "$output_name" "../$program_name"

    checksum=$(sha256sum "$output_name" | awk '{ print $1 }')
    echo "$checksum  $output_name" >> "$program_name/checksums.txt"
done
