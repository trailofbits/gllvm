
develop:
	go install github.com/SRI-CSL/gllvm/cmd/...



dumpsections: 
	readelf -x .gllvm_flags data/hello
	readelf -x .gllvm_flags data/hello2

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
	rm -f data/hello data/hello2 [td]*/.*.o [td]*/.*.bc [td]*/*.o
