build:
	mkdir -p ./bin
	cp ./tiles.txt ./bin/
	templ generate
	go build -o ./bin/main

run: build
	./bin/main
