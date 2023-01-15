BINARY=advent
build_go:
	go build -o ./${BINARY}X ./${BINARY}.go

run_go:
	./${BINARY}X

test_go: 
	go test -v

god: test_go build_go run_go

clean_go:
	rm -f ./${BINARY}

clean_binaries:
	rm -f ./*X
