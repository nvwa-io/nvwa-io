# define compiled agent name
NAME=nvwa-agent

.PHONY: clean build tidy vendor

# build nvwa-agent
build:
	GO111MODULE=on go build -ldflags "-s -w" -o ${NAME} -mod=vendor main.go
#	GO111MODULE=on go build -ldflags "-s -w" -o ${NAME}  main.go
	@echo "${NAME} builded ... ok"

tidy:
	GO111MODULE=on go mod tidy

vendor:
	GO111MODULE=on go mod vendor

test-notify:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
	GO111MODULE=on go build -ldflags "-s -w" -o ${NAME} -mod=vendor main.go

	./${NAME} notify --app=abc notify  --message="Hello world" --users=abc,bcd,def -c=etc/config.toml

test-build-info:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
	GO111MODULE=on go build -ldflags "-s -w" -o ${NAME} -mod=vendor main.go

	export GIT_BRANCH=master && export GIT_COMMIT=abcdefghijklmn && export BUILD_NUMBER=5 && \
	./${NAME} build-info --build-id=1 --status=building -c=etc/config.toml

test-tar:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
	GO111MODULE=on go build -ldflags "-s -w" -o ${NAME} -mod=vendor main.go

	export WORKSPACE=`pwd` && export GIT_BRANCH=master && export GIT_COMMIT=abcdefghijklmn && \
	./${NAME} tar --app=nvwa-demo --build-id=1 --files=main.go,option/config.go --excludes=logic -c=etc/config.toml

test-pull:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
	GO111MODULE=on go build -ldflags "-s -w" -o ${NAME} -mod=vendor main.go

	./${NAME} pull --app=nvwa-demo --version-pkg=nvwa-demo.11.master.abcdef.20190101010101.tar.gz  -c=etc/config.toml

test-push:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
	GO111MODULE=on go build -ldflags "-s -w" -o ${NAME} -mod=vendor main.go

	./${NAME} push --app=nvwa-demo --build-id=1 --version-pkg=nvwa-demo.1.master.abcdefgh.20190112142046.tar.gz -c=etc/config.toml


clean:
	rm -f ${NAME}
	@echo "remove ${NAME} ... ok"
