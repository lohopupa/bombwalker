FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN GOARCH=wasm GOOS=js go build -o build/main.wasm game/main.go

FROM nginx:alpine

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/build/main.wasm .
COPY --from=builder /usr/local/go/misc/wasm/wasm_exec.js .
COPY web/* .
COPY web/assets ./assets

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
