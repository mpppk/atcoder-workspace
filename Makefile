AAAnumber := AAA=int,int8,int16,int32,int64,float32,float64
AAAint := AAA=int,int8,int16,int32,int64
AAAnonnum := AAA=rune,bool,string
BBBnumber := BBB=int,int8,int16,int32,int64,float32,float64
YYY := YYY=rune,bool,string,int,int8,int16,int32,int64,float32,float64
ZZZ := ZZZ=rune,bool,string,int,int8,int16,int32,int64,float32,float64
CONTESTS_DIR := contents
SUBMIT_DIR := submit
SUBMIT_FILE := submit.go
SUBMIT_FILE_PATH := ${SUBMIT_DIR}/${SUBMIT_FILE}

# 提出用ファイルから実行ファイルを生成します
# ex) make build pkg=abc158/A
.PHONY: build
build: bundle
	go build -o ${CONTESTS_DIR}/${pkg}/main ${CONTESTS_DIR}/${pkg}/${SUBMIT_DIR}/${SUBMIT_FILE}

# コードをatcoderへ提出します FIXME: atcoder-toolsのreturn codeが255
# ex) make submit pkg=abc158/A
.PHONY: submit
submit:
	$(MAKE) build pkg=${pkg}
	atcoder-tools submit --timeout 2 --dir ./${CONTESTS_DIR}/${pkg} --exec ./main --code ${CONTESTS_DIR}/${pkg}/${SUBMIT_DIR}/${SUBMIT_FILE} -u ${flag}

# 指定したパッケージのテストを実施します
# ex) make test pkg=abc158/A
.PHONY: test
test: bundle
	$(MAKE) build pkg=${pkg}
	atcoder-tools test --dir ./${CONTESTS_DIR}/${pkg} --exec ./main -s

# 自動生成されたコードを削除します
# ex) make clean pkg=lib
.PHONY: clean
clean:
	find ${pkg} -name "must-*.go" | xargs rm
	find ${pkg} -name "gen-*.go" | xargs rm
	find ${pkg} -name "submit/submit.go" | xargs rm

# 必要なツールをインストールします。2020/03/29時点では、Go対応パッチがatcoder-toolsには取り込まれていないので、
# 実際にはhttps://github.com/nu50218/atcoder-toolsをpip install -eして利用しています
.PHONY: setup
setup:
	go get github.com/mpppk/mustify
	go get github.com/mpppk/gollup
	go get github.com/cheekybits/genny
	pip3 install atcoder-tools

# atocderへの提出用コードを生成します。
# ex) make bundle pkg=abc158/A
.PHONY: bundle
bundle:
	mkdir -p ${CONTESTS_DIR}/${pkg}/${SUBMIT_DIR}
	gollup ./${CONTESTS_DIR}/${pkg} ./lib > ${SUBMIT_FILE}
	mv ${SUBMIT_FILE} ${CONTESTS_DIR}/${pkg}/${SUBMIT_DIR}

# 指定したコンテストの実施環境を作成します。
# 各設問のパッケージやテストなどが生成されます。
# ex) make new contest=abc158
.PHONY: new
new:
	atcoder-tools gen --workspace ./${CONTESTS_DIR} --lang go --template ./templates/main.tmpl ${contest}
	find ./${CONTESTS_DIR}/${contest}/**/*.go -type f | xargs goimports -w

# ライブラリのテストを行います.
.PHONY: test-lib
test-lib:
	go test ./lib/...

# コードの自動生成を行います。
.PHONY: generate
generate: mustify

# エラーを返す関数をMustXXXとしてラップした関数としてmust-xxx.goに出力します
.PHONY: mustify
mustify: genny
	find ./lib/*.go -type f | xargs -I{} bash -c 'mustify {} > $$(dirname {})/must-$$(basename {})'
	find ./lib/*.go -type f | xargs -I{} bash -c 'test -s {} || rm {}'

# ジェネリクスを利用したコードから実際に利用するコードを生成します
.PHONY: genny
genny:
	$(MAKE) clean pkg=lib
	genny -in='./lib/number.go' -out='./lib/gen-number.go' gen "$(AAAnumber)"
	genny -in='./lib/math.go' -out='./lib/gen-math.go' gen "$(AAAnumber)"
	genny -in='./lib/number2.go' -out='./lib/gen-number2.go' gen "$(AAAnumber) $(BBBnumber)"
	genny -in='./lib/int.go' -out='./lib/gen-int.go' gen "$(AAAint)"
	genny -in='./lib/type.go' -out='./lib/gen-type.go' gen "$(ZZZ)"
	genny -in='./lib/type2.go' -out='./lib/gen-type2.go' gen "$(YYY) $(ZZZ)"
	genny -in='./lib/misc.go' -out='./lib/gen-misc.go' gen "$(AAAnumber)"
	genny -in='./lib/input-number.go' -out='./lib/gen-input-number.go' gen "$(AAAnumber)"
