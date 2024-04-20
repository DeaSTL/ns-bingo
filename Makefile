build:
	mkdir -p ./bin
	cp ./tiles.txt ./bin/
	cd src;templ generate
	cd src;go build -o ../bin/main

run: build
	./bin/main
