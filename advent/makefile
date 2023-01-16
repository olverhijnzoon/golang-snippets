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

compile_c:
	gcc -o adventC advent.c

run_c:
	./${BINARY}C

cdo: compile_c run_c

clean_binaries:
	rm -f ./*X ./*C
