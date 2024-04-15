.PHONY: buildserver
buildserver:
	go build -o ./bin/server.exe -v ./cmd

.PHONY: runserver
runserver:
	go build -o ./bin/server.exe -v ./cmd
	./bin/server.exe

.DEFAULT_GOAL := runserver