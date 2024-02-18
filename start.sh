#!/usr/bash

export GOARCH=wasm
export GOOS=js

go build -o build/main.wasm game/game.go
python3 -m http.server 6969 --directory web
