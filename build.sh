#!/bin/bash

# Get the latest version from Git tag
#version=$(git describe --abbrev=0 --tags)
version="0.0.1"

directory="releases/$version"
if [ ! -d "$directory" ]; then
    mkdir -p "$directory"
    echo "Directory created: $directory"
else
    echo "Directory already exists: $directory"
fi

env GOOS=linux GOARCH=amd64 go build -o="$directory/jakvta-$version-linux-x64" $PWD/src/jakvta/main.go
env GOOS=darwin GOARCH=amd64 go build -o="$directory/jakvta-$version-osx-x64" $PWD/src/jakvta/main.go
env GOOS=windows GOARCH=amd64 go build -o="$directory/jakvta-$version-win-x64.exe" $PWD/src/jakvta/main.go
env GOOS=linux GOARCH=arm64 go build -o="$directory/jakvta-$version-linux-arm" $PWD/src/jakvta/main.go
env GOOS=darwin GOARCH=arm64 go build -o="$directory/jakvta-$version-osx-arm" $PWD/src/jakvta/main.go
env GOOS=windows GOARCH=arm64 go build -o="$directory/jakvta-$version-win-arm.exe" $PWD/src/jakvta/main.go
