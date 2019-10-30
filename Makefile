.PHONY : all server clean

all: server

server:
	go build -o dist/server server.go

clean:
	rm -rf dist/*