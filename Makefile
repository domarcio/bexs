.PHONY: all
all: build
FORCE: ;

SHELL := env FILENAME=$(FILENAME) $(SHELL)

.PHONY: build

clean:
	@echo "|bexs| -- CLEAN BINARIES FILES -- |bexs|"
	@rm -rf bin/*

image:
	@echo "|bexs| -- GENERATE DOCKER IMAGE -- |bexs|"
	@docker build --network=host -t bexs-nogues -f ./docker/Dockerfile .

build-cmd:
	@echo "|bexs| -- GENERATE CMD BINARY FILE -- |bexs|"
	@docker run --name bexs-nogues-build-cmd --rm -v $(PWD):/app:Z -w /app -e ENV=dev bexs-nogues ./scripts/make.sh build-cmd

run-cmd:
	@echo "|bexs| -- RUNNING CMD -- |bexs|"
	@docker run -it --name bexs-nogues-build-cmd --rm -v $(PWD):/app:Z -w /app -e ENV=dev bexs-nogues ./scripts/make.sh run-cmd $(FILENAME)

build-api:
	@echo "|bexs| -- GENERATE API BINARY FILE -- |bexs|"
	@docker run --name bexs-nogues-build-api --rm -v $(PWD):/app:Z -w /app -e ENV=dev bexs-nogues ./scripts/make.sh build-api

run-api:
	@echo "|bexs| -- RUNNING CMD -- |bexs|"
	@docker run -it --name bexs-nogues-build-api --rm -p 7007:7007 -v $(PWD):/app:Z -w /app -e ENV=dev bexs-nogues ./scripts/make.sh run-api

test:
	@echo "|bexs| -- EXECUTE ALL TESTS -- |bexs|"
	@rm -f ./data/storage/testing.routes.csv
	@touch ./data/storage/testing.routes.csv
	@docker run --name bexs-nogues-test --rm -v $(PWD):/app:Z -w /app bexs-nogues ./scripts/make.sh test