FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN GOARCH=wasm GOOS=js go build -o build/main.wasm game/main.go

FROM nginx:latest

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/build/main.wasm .
COPY web/* .

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
