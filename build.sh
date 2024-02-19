#!/bin/bash

GOARCH=wasm GOOS=js go build -o ./build/main.wasm game/main.go

mkdir -p ./build

cp ./build/main.wasm ./build/
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./build/
cp -r ./web/* ./build/

