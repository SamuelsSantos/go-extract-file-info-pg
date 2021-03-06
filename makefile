EXTRACTOR=extractor-cmd
SERVER=server-extractor

VOLUME=$VOLUME_PATH
FILENAME=$FILENAME

clean:
	rm -rf build
	go clean -modcache

build: clean
	go build -o build/$(EXTRACTOR) import-data.go
	go build -o build/$(SERVER) cmd/app.go

run:
	go run $(EXTRACTOR).go

docker-build: build
	@docker image build -t $(EXTRACTOR) . -f extractor.Dockerfile
	@docker image build -t $(SERVER) . -f server.Dockerfile


docker-compose: docker-build
	@docker-compose up -d --build

docker-clean: docker-down
	@docker rmi -f $(EXTRACTOR)
	@docker rmi -f $(SERVER)
	@docker volume rm -f import-data_db_data
	@docker image prune -f

docker-down:
	@docker-compose down

docker-run:
	@docker run --rm -v $(VOLUME):/files --env-file=.env --name extractor-cmd --net import-data_net extractor-cmd:latest -file=/files/$(FILENAME)