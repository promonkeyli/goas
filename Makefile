# basic info of project
MODULE_NAME := goas

TEST_CMD := cmd/test/main.go

# run test command
test:
	go run $(TEST_CMD)