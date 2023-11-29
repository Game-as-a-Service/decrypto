SERVER_PATH=./server


.PHONY: server_test
# server_test
server_test:
	cd $(SERVER_PATH) && go test -v ./... -cover