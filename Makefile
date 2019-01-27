SHELL = /bin/bash

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	find ${pkg} -type f -name "bundled_*.go" | xargs rm

.PHONY: bundle
bundle: bundle-utl
	bundle -pkg main -prefix ' ' -dst github.com/mpppk/atcoder/${pkg} github.com/mpppk/atcoder/${pkg} | pbcopy

.PHONY: bundle-utl
bundle-utl: generate
	bundle -o ${pkg}/bundled_utl.go -pkg main -dst github.com/mpppk/atcoder/${pkg} github.com/mpppk/atcoder/utl

.PHONY: generate
generate:
	genny -in='./slice/slice.go' -out='./slice/gen-slice.go' gen "Value=int,int8,int16,int32,int64,float32,float64"
	genny -in='./slice/string.go' -out='./slice/gen-string.go' gen "Type=rune,string,int,int8,int16,int32,int64,float32,float64"
	genny -in='./slice/misc.go' -out='./slice/gen-misc.go' gen "Value=int,int8,int16,int32,int64,float32,float64"
	genny -in='./slice/slice.go' -out='./utl/gen-slice.go' -pkg utl gen "Value=int,int8,int16,int32,int64,float32,float64"
	genny -in='./slice/string.go' -out='./utl/gen-string.go' -pkg utl gen "Type=rune,string,int,int8,int16,int32,int64,float32,float64"
	genny -in='./slice/misc.go' -out='./utl/gen-misc.go' -pkg utl gen "Value=int,int8,int16,int32,int64,float32,float64"


