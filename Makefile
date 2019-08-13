test:
	@@go test -v -cover -race -covermode=atomic ./test/...

.PHONY: test
