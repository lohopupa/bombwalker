#!/bin/bash
mkdir -p ./build

GOARCH=wasm GOOS=js go build -o ./build/main.wasm game/main.go

cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./build/
cp -r ./web/* ./build/
cp -r ./assets/web/* ./build/assets

