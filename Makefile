SHELL = /bin/bash

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	find ${pkg} -name "bundled_*.go" | xargs rm
	find ${pkg} -name "must-*.go" | xargs rm
	find ${pkg} -name "gen-*.go" | xargs rm

.PHONY: bundle
bundle: bundle-utl
	bundle -pkg main -prefix ' ' -dst github.com/mpppk/atcoder/${pkg} github.com/mpppk/atcoder/${pkg} | pbcopy

.PHONY: bundle-utl
bundle-utl: generate
	bundle -o ${pkg}/bundled_utl.go -pkg main -dst github.com/mpppk/atcoder/${pkg} github.com/mpppk/atcoder/utl

.PHONY: generate
generate:
	genny -in='./generic/number.go' -out='./generic/gen-number.go' gen "Value=int,int8,int16,int32,int64,float32,float64"
	genny -in='./generic/int.go' -out='./generic/gen-int.go' gen "Value=int,int8,int16,int32,int64"
	genny -in='./generic/type.go' -out='./generic/gen-type.go' gen "Type=rune,string,int,int8,int16,int32,int64,float32,float64"
	genny -in='./generic/misc.go' -out='./generic/gen-misc.go' gen "Value=int,int8,int16,int32,int64,float32,float64"
	genny -in='./generic/number.go' -out='./utl/gen-number.go' -pkg utl gen "Value=int,int8,int16,int32,int64,float32,float64"
	genny -in='./generic/int.go' -out='./utl/gen-int.go' -pkg utl gen "Value=int,int8,int16,int32,int64"
	genny -in='./generic/type.go' -out='./utl/gen-type.go' -pkg utl gen "Type=rune,string,int,int8,int16,int32,int64,float32,float64"
	genny -in='./generic/misc.go' -out='./utl/gen-misc.go' -pkg utl gen "Value=int,int8,int16,int32,int64,float32,float64"
	goofy mustify --file ./utl/input.go --out ./utl/must-input.go
	goimports -w ./utl/must-input.go
	goofy mustify --file ./utl/gen-number.go --out ./utl/must-gen-number.go
	goimports -w ./utl/must-gen-number.go
	goofy mustify --file ./utl/gen-int.go --out ./utl/must-gen-int.go
	goimports -w ./utl/must-gen-int.go
	goofy mustify --file ./utl/gen-type.go --out ./utl/must-gen-type.go
	goimports -w ./utl/must-gen-type.go
	goofy mustify --file ./utl/gen-misc.go --out ./utl/must-gen-misc.go
	goimports -w ./utl/must-gen-misc.go
