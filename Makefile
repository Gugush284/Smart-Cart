.PHONY: server
server:
	cd ./server && make

.PHONY: client
client:
	cd ./third_party/tg_client && make

.PHONY: device
device:
	cd ./Smart-Cart-Device && make

.PHONY: dtest
dtest:
	cd ./Smart-Cart-Device && make connect

.DEFAULT_GOAL := server