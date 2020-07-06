EXTRACTOR=import-data
SERVER=app

clean:
	go clean

build: clean
	go build -o build/$(EXTRACTOR) $(EXTRACTOR).go
	go build -o build/$(SERVER) cmd/$(SERVER).go

run:
	go run $(EXTRACTOR).go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o build/$(EXTRACTOR)-linux-arm $(EXTRACTOR).go
	GOOS=linux GOARCH=arm64 go build -o build/$(EXTRACTOR)-numero-linux-arm64 $(EXTRACTOR).go
	GOOS=freebsd GOARCH=386 go build -o build/$(EXTRACTOR)-numero-freebsd-386 $(EXTRACTOR).go

	GOOS=linux GOARCH=arm go build -o build/$(SERVER)-linux-arm $(SERVER).go
	GOOS=linux GOARCH=arm64 go build -o build/$(SERVER)-numero-linux-arm64 $(SERVER).go
	GOOS=freebsd GOARCH=386 go build -o build/$(SERVER)-numero-freebsd-386 $(SERVER).go

all: $(EXTRACTOR) build

docker-build: build
	@docker image build -t $(EXTRACTOR) .


docker-run:
	@docker run --rm -v /Users/samuelsantos/git/import-data/resource:/files --EXTRACTOR import-data --net import-data_net import-data:latest -file=/files/test.txt 