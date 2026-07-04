.PHONY: test test-email test-file test-path test-maps test-string test-task

test:
	go test -v ./...

test-email:
	go test -v ./email

test-file:
	go test -v ./file

test-path:
	go test -v ./pathutil

test-maps:
	go test -v ./maputil

test-string:
	go test -v ./strutil

test-task:
	go test -v ./csvutil
