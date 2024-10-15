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

env GOOS=linux GOARCH=amd64 go build -o="$directory/jakvta-$version-linux" $PWD/src/jakvta/main.go
env GOOS=darwin GOARCH=amd64 go build -o="$directory/jakvta-$version-osx" $PWD/src/jakvta/main.go
env GOOS=windows GOARCH=amd64 go build -o="$directory/jakvta-$version-win.exe" $PWD/src/jakvta/main.go