# define compiled bin
NAME=nvwa-server

.PHONY: clean

build:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go build -ldflags "-s -w" -o ${NAME} -mod=vendor main.go
	@echo "${NAME} builded ... ok"

init:
	GO111MODULE=on go mod init

clean:
	rm -f ${NAME}
	@echo "remove ${NAME} ... ok"
