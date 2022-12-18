BINARY=advents

build_go:
	go build -o ./${BINARY} ./${BINARY}.go
	
run_go:
	./${BINARY}

god: build_go run_go

clean_go:
	rm -f ./${BINARY}