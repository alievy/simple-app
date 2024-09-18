.PHONY: build run clean

build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./main.go

run: build
	go run ./main.go

clean:
	rm -f web/app.wasm
