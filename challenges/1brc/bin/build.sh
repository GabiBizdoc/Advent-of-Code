#!/bin/bash

platforms=("linux/amd64" "linux/386" "linux/arm64" "windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64")

program_name="generator"

> "$program_name/checksums.txt"
for platform in "${platforms[@]}"; do
    # Split the platform string
    IFS='/' read -ra split <<< "$platform"
    GOOS="${split[0]}"
    GOARCH="${split[1]}"

    # Name of the output binary
    output_name="$program_name/$program_name-$GOOS-$GOARCH"

    if [[ "windows" == $GOOS ]]; then
     output_name="$output_name.exe"
    fi

    # Build for the current platform
    env GOOS=$GOOS GOARCH=$GOARCH go build -o "$output_name" "../$program_name"

    # Calculate checksum
    checksum=$(sha256sum "$output_name" | awk '{ print $1 }')
    echo "$checksum  $output_name" >> "$program_name/checksums.txt"
done
