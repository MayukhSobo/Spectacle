OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
BIN_EXT := ""
ifeq ($(OS),windows)
	BIN_EXT := ".exe"
endif

.DEFAULT_GOAL := run

run: build
	./build/spectacle$(BIN_EXT)

build:
	mkdir -p build
	GOOS=$(OS) go build -o build/spectacle$(BIN_EXT) main.go

clean:
	rm -rf build

.PHONY: run build clean 