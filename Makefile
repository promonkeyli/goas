# basic info of project
MODULE_NAME := goas

GOAS_CMD := cmd/goas/main.go

# run test command
test: build
	./bin/goas -dir "./example/cmd,./example/internal" -output "example/api"

# goas
goas:
	go run $(GOAS_CMD)

# build goas binary
build:
	go build -o bin/goas $(GOAS_CMD)