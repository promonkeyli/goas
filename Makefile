# basic info of project
MODULE_NAME := goas

TEST_CMD := cmd/test/main.go
GOAS_CMD := cmd/goas/main.go

# run test command
test:
	go run $(TEST_CMD)

# goas
goas:
	go run $(GOAS_CMD)

# build goas binary
build:
	go build -o bin/goas $(GOAS_CMD)