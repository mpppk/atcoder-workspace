SHELL = /bin/bash
AAAnumber := AAA=int,int8,int16,int32,int64,float32,float64
AAAint := AAA=int,int8,int16,int32,int64
BBBnumber := BBB=int,int8,int16,int32,int64,float32,float64
YYY := YYY=rune,string,int,int8,int16,int32,int64,float32,float64
ZZZ := ZZZ=rune,string,int,int8,int16,int32,int64,float32,float64
REPO_PATH := github.com/mpppk/atcoder
COPY_FILES := ./lib/type0.go ./lib/number0.go ./lib/string.go ./lib/input.go ./lib/utl.go ./lib/graph.go

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	find ${pkg} -name "bundled_*.go" | xargs rm
	find ${pkg} -name "must-*.go" | xargs rm
	find ${pkg} -name "gen-*.go" | xargs rm

.PHONY: bundle
bundle:
	bundle -pkg main -prefix ' ' -dst $(REPO_PATH)/${pkg} $(REPO_PATH)/${pkg} | pbcopy

.PHONY: new
new:
	atcoder-tools gen --workspace . --lang go --template ./templates/main.tmpl ${contest}
	find ./${contest}/**/*.go -type f | xargs goimports -w
	$(MAKE) bundle-gen-all contest=${contest}


.PHONY: bundle-gen-all
bundle-gen-all:
	$(MAKE) bundle-gen pkg=${contest}/A
	cp ${contest}/A/bundled_gen.go ${contest}/B
	cp ${contest}/A/bundled_gen.go ${contest}/C
	cp ${contest}/A/bundled_gen.go ${contest}/D
	cp ${contest}/A/bundled_gen.go ${contest}/E
	cp ${contest}/A/bundled_gen.go ${contest}/F

.PHONY: bundle-gen
bundle-gen: mustify-gen
	bundle -o ${pkg}/bundled_gen.go -pkg main -dst $(REPO_PATH)/${pkg} $(REPO_PATH)/gen

.PHONY: mustify-gen
mustify-gen: genny-gen
	goofy mustify --file ./gen/input.go
	find ./gen/*.go -type f | xargs goimports -w

.PHONY: genny-gen
genny-gen: genny-lib
	cp $(COPY_FILES) ./gen
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
