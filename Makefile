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
bundle: bundle-generic
	bundle -pkg main -prefix ' ' -dst github.com/mpppk/atcoder/${pkg} github.com/mpppk/atcoder/${pkg} | pbcopy

.PHONY: bundle-generic
bundle-generic: generate
	bundle -o ${pkg}/bundled_generic.go -pkg main -dst github.com/mpppk/atcoder/${pkg} github.com/mpppk/atcoder/generic

.PHONY: generate
generate:
	genny -in='./generic/number.go' -out='./generic/gen-number.go' gen "AAA=int,int8,int16,int32,int64,float32,float64"
	genny -in='./generic/number2.go' -out='./generic/gen-number2.go' gen "AAA=int,int8,int16,int32,int64,float32,float64 BBB=int,int8,int16,int32,int64,float32,float64"
	genny -in='./generic/int.go' -out='./generic/gen-int.go' gen "AAA=int,int8,int16,int32,int64"
	genny -in='./generic/type.go' -out='./generic/gen-type.go' gen "ZZZ=rune,string,int,int8,int16,int32,int64,float32,float64"
	genny -in='./generic/misc.go' -out='./generic/gen-misc.go' gen "AAA=int,int8,int16,int32,int64,float32,float64"
	genny -in='./generic/input-number.go' -out='./generic/gen-input-number.go' gen "AAA=int,int8,int16,int32,int64,float32,float64"
#	genny -in='./generic/number.go' -out='./utl/gen-number.go' -pkg utl gen "Value=int,int8,int16,int32,int64,float32,float64"
#	genny -in='./generic/number2.go' -out='./utl/gen-number2.go' -pkg utl gen "ZZZ=int,int8,int16,int32,int64,float32,float64 Type=int,int8,int16,int32,int64,float32,float64"
#	genny -in='./generic/int.go' -out='./utl/gen-int.go' -pkg utl gen "Value=int,int8,int16,int32,int64"
#	genny -in='./generic/type.go' -out='./utl/gen-type.go' -pkg utl gen "Type=rune,string,int,int8,int16,int32,int64,float32,float64"
#	genny -in='./generic/misc.go' -out='./utl/gen-misc.go' -pkg utl gen "Value=int,int8,int16,int32,int64,float32,float64"
#	genny -in='./utl/input-number.go' -out='./utl/gen-input-number.go' gen "Value=int,int8,int16,int32,int64,float32,float64"
	goofy mustify --file ./generic/input.go --out ./generic/must-input.go
	goimports -w ./generic/must-input.go
	goofy mustify --file ./generic/gen-number.go --out ./generic/must-gen-number.go
	goimports -w ./generic/must-gen-number.go
	goofy mustify --file ./generic/gen-number2.go --out ./generic/must-gen-number2.go
	goimports -w ./generic/must-gen-number2.go
	goofy mustify --file ./generic/gen-int.go --out ./generic/must-gen-int.go
	goimports -w ./generic/must-gen-int.go
	goofy mustify --file ./generic/gen-type.go --out ./generic/must-gen-type.go
	goimports -w ./generic/must-gen-type.go
	goofy mustify --file ./generic/gen-misc.go --out ./generic/must-gen-misc.go
	goimports -w ./generic/must-gen-misc.go
