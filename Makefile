.PHONY: all
all: build
FORCE: ;

SHELL := env ENV=$(ENV) FILENAME=$(FILENAME) $(SHELL)
ENV ?= dev

.PHONY: build

clean:
	@echo "|bexs| -- CLEAN BINARIES FILES -- |bexs|"
	@rm -rf bin/*

image:
	@echo "|bexs| -- GENERATE DOCKER IMAGE -- |bexs|"
	@docker build --network=host -t bexs-nogues -f ./docker/Dockerfile .

build-cmd:
	@echo "|bexs| -- GENERATE CMD BINARY FILE -- |bexs|"
	@docker run --name bexs-nogues-build-cmd --rm -v $(PWD):/app:Z -w /app -e ENV=$(ENV) bexs-nogues ./scripts/make.sh build-cmd

run-cmd:
	@echo "|bexs| -- RUNNING CMD -- |bexs|"
	@docker run -it --name bexs-nogues-build-cmd --rm -v $(PWD):/app:Z -w /app -e ENV=$(ENV) bexs-nogues ./scripts/make.sh run-cmd $(FILENAME)

test:
	@echo "|bexs| -- EXECUTE ALL TESTS -- |bexs|"
	@docker run --name bexs-nogues-test --rm -v $(PWD):/app:Z -w /app -e ENV=$(ENV) bexs-nogues ./scripts/make.sh test