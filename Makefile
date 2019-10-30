.PHONY : all server clean

all: server

server:
	go build -o dist/server main.go

clean:
	rm -rf dist/*