SHELL = /bin/bash
AAAnumber := AAA=int,int8,int16,int32,int64,float32,float64
AAAint := AAA=int,int8,int16,int32,int64
BBBnumber := BBB=int,int8,int16,int32,int64,float32,float64
YYY := YYY=rune,string,int,int8,int16,int32,int64,float32,float64
ZZZ := ZZZ=rune,string,int,int8,int16,int32,int64,float32,float64

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	find ${pkg} -name "bundled_*.go" | xargs rm
	find ${pkg} -name "must-*.go" | xargs rm
	find ${pkg} -name "gen-*.go" | xargs rm

.PHONY: bundle
bundle: bundle-gen
	bundle -pkg main -prefix ' ' -dst github.com/mpppk/atcoder/${pkg} github.com/mpppk/atcoder/${pkg} | pbcopy

.PHONY: bundle-gen
bundle-gen: mustify-gen
	bundle -o ${pkg}/bundled_gen.go -pkg main -dst github.com/mpppk/atcoder/${pkg} github.com/mpppk/atcoder/gen

.PHONY: mustify-gen
mustify-gen: genny-gen
	goofy mustify --file ./gen/input.go
	find ./gen/*.go -type f | xargs goimports -w

.PHONY: genny-gen
genny-gen: genny-lib
	cp ./lib/number0.go ./gen/number0.go
	cp ./lib/string.go ./gen/string.go
	cp ./lib/input.go ./gen/input.go
	cp ./lib/utl.go ./gen/utl.go
	cp ./lib/graph.go ./gen/graph.go
	genny -in='./lib/number.go' -out='./gen/gen-number.go' gen "$(AAAnumber)"
	genny -in='./lib/number2.go' -out='./gen/gen-number2.go' gen "$(AAAnumber) $(BBBnumber)"
	genny -in='./lib/int.go' -out='./gen/gen-int.go' gen "$(AAAint)"
	genny -in='./lib/type.go' -out='./gen/gen-type.go' gen "$(ZZZ)"
	genny -in='./lib/type2.go' -out='./gen/gen-type2.go' gen "$(YYY) $(ZZZ)"
	genny -in='./lib/misc.go' -out='./gen/gen-misc.go' gen "$(AAAnumber)"
	genny -in='./lib/input-number.go' -out='./gen/gen-input-number.go' gen "$(AAAnumber)"

.PHONY: genny-lib
genny-lib:
	genny -in='./lib/number.go' -out='./lib/gen-number.go' gen "$(AAAnumber)"
	genny -in='./lib/number2.go' -out='./lib/gen-number2.go' gen "$(AAAnumber) $(BBBnumber)"
	genny -in='./lib/int.go' -out='./lib/gen-int.go' gen "$(AAAint)"
	genny -in='./lib/type.go' -out='./lib/gen-type.go' gen "$(ZZZ)"
	genny -in='./lib/type2.go' -out='./lib/gen-type2.go' gen "$(YYY) $(ZZZ)"
	genny -in='./lib/misc.go' -out='./lib/gen-misc.go' gen "$(AAAnumber)"
	genny -in='./lib/input-number.go' -out='./lib/gen-input-number.go' gen "$(AAAnumber)"
