SHELL = /bin/bash

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	find ${pkg} -type f -name "bundled_*.go" | xargs rm

.PHONY: bundle
bundle:
	make bundle-utl
	bundle -pkg main -prefix ' ' -dst github.com/mpppk/atcoder/${pkg} github.com/mpppk/atcoder/${pkg} | pbcopy

.PHONY: bundle-utl
bundle-utl:
	make generate
	bundle -o ${pkg}/bundled_utl.go -pkg main -dst github.com/mpppk/atcoder/${pkg} github.com/mpppk/atcoder/utl

.PHONY: generate
generate:
	genny -in='./slice/slice.go' -out='./utl/gen-slice.go' -pkg utl gen "Value=int,int16,int32,int64,float32,float64"

