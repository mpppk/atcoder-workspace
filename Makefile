SHELL = /bin/bash
AAAnumber := AAA=int,int8,int16,int32,int64,float32,float64
AAAint := AAA=int,int8,int16,int32,int64
BBBnumber := BBB=int,int8,int16,int32,int64,float32,float64
YYY := YYY=rune,string,int,int8,int16,int32,int64,float32,float64
ZZZ := ZZZ=rune,string,int,int8,int16,int32,int64,float32,float64
REPO_PATH := github.com/mpppk/atcoder
COPY_FILES := ./lib/type0.go ./lib/number0.go ./lib/string.go ./lib/input.go ./lib/utl.go ./lib/graph.go

.PHONY: build
build:
	go build -o ${pkg}/main ${pkg}/main.go ${pkg}/bundled_gen.go

# 指定したパッケージのテストを実施します
# ex) make test pkg=abc158/A
.PHONY: test
test:
	$(MAKE) build pkg=${pkg}
	atcoder-tools test --dir ./${pkg} --exec ./main -s

# 自動生成されたコードを削除します
.PHONY: clean
clean:
	find ${pkg} -name "bundled_*.go" | xargs rm
	find ${pkg} -name "must-*.go" | xargs rm
	find ${pkg} -name "gen-*.go" | xargs rm

# 必要なツールをインストールします。2020/03/29時点では、Go対応パッチがatcoder-toolsには取り込まれていないので、
# 実際にはhttps://github.com/nu50218/atcoder-toolsをpip install -eして利用しています
.PHONY: setup
setup:
	pip3 install atcoder-tools
	go get github.com/mpppk/atcoder
	go get github.com/mpppk/goofy
	go get github.com/cheekybits/genny

# atocderへの提出用に、指定された設問パッケージのソースコードを1ファイルにバンドルしたソースコードをクリップボードへコピーします
# pbcopyを利用しているので、macでしか動きません
# ex) make bundle pkg=abc158/A
.PHONY: bundle
bundle:
	gollup --dirs ./${pkg},./lib ${pkg}/main.go | goimports | pbcopy

# 指定したコンテストの実施環境を作成します。
# 各設問のパッケージ、バンドルされたライブラリ、テストなどが生成されます。
# ex) make new contest=abc158
.PHONY: new
new:
	atcoder-tools gen --workspace . --lang go --template ./templates/main.tmpl ${contest}
	find ./${contest}/**/*.go -type f | xargs goimports -w

.PHONY: generate
generate: mustify-gen

# エラーを返す関数をMustXXXとしてラップした関数としてmust-xxx.goに出力します
.PHONY: mustify-gen
mustify-gen: genny-lib
	mustify ./lib | xargs goimports -w

.PHONY: genny-lib
genny-lib:
	genny -in='./lib/number.go' -out='./lib/gen-number.go' gen "$(AAAnumber)"
	genny -in='./lib/number2.go' -out='./lib/gen-number2.go' gen "$(AAAnumber) $(BBBnumber)"
	genny -in='./lib/int.go' -out='./lib/gen-int.go' gen "$(AAAint)"
	genny -in='./lib/type.go' -out='./lib/gen-type.go' gen "$(ZZZ)"
	genny -in='./lib/type2.go' -out='./lib/gen-type2.go' gen "$(YYY) $(ZZZ)"
	genny -in='./lib/misc.go' -out='./lib/gen-misc.go' gen "$(AAAnumber)"
	genny -in='./lib/input-number.go' -out='./lib/gen-input-number.go' gen "$(AAAnumber)"
