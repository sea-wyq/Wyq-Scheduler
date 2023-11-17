all: local

local: fmt vet
	GOOS=linux GOARCH=amd64 go build  -o=bin/wyq-scheduler ./cmd/scheduler

build: local
	docker build --no-cache -t registry.cnbita.com:5000/wuyiqiang/wyq-scheduler:v1 .

push: build
	docker push registry.cnbita.com:5000/wuyiqiang/wyq-scheduler:v1

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...
