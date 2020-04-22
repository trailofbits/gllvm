develop:
	go install github.com/SRI-CSL/gllvm/cmd/...



dumpsections: check
	readelf -x .gllvm_flags data/hello

deps:
	sudo apt install clang llvm
	touch $@

check: develop deps
	 go test -v ./tests


tests: develop deps
	WLLVM_OUTPUT_LEVEL=DEBUG ; go test -v ./tests

format:
	gofmt -s -w shared/*.go tests/*.go cmd/*/*.go


clean:
	rm -f data/hello data/helloworld data/hello.bc [td]*/.helloworld.c.o [td]*/.helloworld.c.o.bc
