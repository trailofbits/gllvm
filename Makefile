
develop:
	go install github.com/SRI-CSL/gllvm/cmd/...



dumpsections: 
	readelf -x .gllvm_flags data/hello

deps:
	sudo apt install clang llvm
	touch $@

check: develop deps
	GLLVM_EMBED_FRONTEND_ARGS=true go test -v ./tests


tests: develop deps
	WLLVM_OUTPUT_LEVEL=DEBUG GLLVM_EMBED_FRONTEND_ARGS=true go test -v ./tests

format:
	gofmt -s -w shared/*.go tests/*.go cmd/*/*.go


clean:
	rm -f data/hello data/hello data/helloworld data/hello.bc [td]*/.helloworld.c.o [td]*/.helloworld.c.o.bc
