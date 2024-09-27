.PHONY: all build run clean

all: build

build:
	go build -o kobo-manga

run: build
	./kobo-manga

clean:
	rm -f kobo-manga