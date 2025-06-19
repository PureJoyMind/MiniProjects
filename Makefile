.PHONY: build
build:
	@go build -o bin\main.exe main.go

.PHONY: run
run: build
	@.\bin\main.exe $(args)

.PHONY: rungif
rungif: build
	@./bin/main.exe >out.gif
