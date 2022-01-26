ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

deps:
	go mod tidy
	cd web && npm i

build:
	cd web && npm run build  
	go build .

update:
	cd web && npm update