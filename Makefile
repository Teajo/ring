all: build install

build:
	go build
run:
	./ring
install:
	cp ring ~/.local/bin/ring