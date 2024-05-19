.PHONY: buildserver
buildserver:
	go build -o ./bin/server.exe -v ./src/cmd

.PHONY: runserver
runserver:
	go build -o ./bin/server.exe -v ./src/cmd
	./bin/server.exe

.DEFAULT_GOAL := runserver