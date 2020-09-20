SERVICE_NAME=gfb-mping

GOOSes = solaris linux windows darwin freebsd
GOARCHes = amd64 386

.PHONY: all dep build test lint

all: dep build test lint

dep:
	GOPROXY=direct go mod download

build:
	go build -o ./${SERVICE_NAME} ./
test:
	go test -race ./...

lint:
	golangci-lint run

install:
	go install ./

release:
	rm -f ./build/*
	$(foreach goos,$(GOOSes), \
		$(foreach goarch,$(GOARCHes), \
			GOOS=${goos} GOARCH=${goarch} go build -o ./build/${SERVICE_NAME}_${goos}_${goarch} ./ || true ; \
		) \
	)
	zip -mr ./build/release-`date +"%y-%m-%d-%H-%M-%S"`.zip ./build/*
