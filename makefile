NAME=import-data

build:
	rm -rf build
	go build -o build/$(NAME) $(NAME).go

run:
	go run $(NAME).go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o build/$(NAME)-linux-arm $(NAME).go
	GOOS=linux GOARCH=arm64 go build -o build/$(NAME)-numero-linux-arm64 $(NAME).go
	GOOS=freebsd GOARCH=386 go build -o build/$(NAME)-numero-freebsd-386 $(NAME).go

all: $(NAME) build

docker-build: build
	@docker image build -t $(NAME) .


docker-run:
	@docker run --rm -v /Users/samuelsantos/git/import-data/resource:/files --name import-data --net import-data_net import-data:latest -file=/files/test.txt 