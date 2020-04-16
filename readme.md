# go-atcoder-workspace

atcoderにGoで参加するための設定がセットアップ済みのリポジトリです。  
現時点ではLinuxとMacに対応しています。  
Windows向けには、WSLとDockerでのセットアップ手順を今後追加予定です。

## 準備
以下をあらかじめインストールしておいてください。
* python3.5以上 ([atcoder-tools](https://github.com/kyuridenamida/atcoder-tools)を動かすのに必要です)
* Go ([mustify](https://github.com/mpppk/mustify)と[gollup](https://github.com/mpppk/gollup)を動かすのに必要です。あと実装したコードのコンパイルにも当然必要です。)

## Setup
```shell
$ git clone https://github.com/mpppk/go-atcoder-workspace
$ cd go-atcoder-workspace
$ make setup # atcoder-tools, mustify, gollupをインストールします
$ make generate # ライブラリコードの自動生成を行います
```

### Note
2020/4/16時点ではatcoder-toolsにGoサポートのパッチがマージされていないので、代わりにforkされたリポジトリからインストールします。
```shell
$ git clone https://github.com/nu50218/atcoder-tools
$ cd atcoder-tools
$ pip3 install -e .
```

### コンテスト実施の流れ
abc999に参加する設定で、コンテストへの参加からサブミットまでを説明します。

1. コンテストを開始するには`make new contest=abc999`を実行します
    * `abc999`というディレクトリが生成されます
    * `abc999`以下にA~Fまでの各設問に対応するディレクトリが生成されます
    * 各設問ディレクトリの`main.go`には、設問内容に応じて入力を受け取るコードがあらかじめ実装されています。
    * `lib` package内のメソッドを利用できます。
1. ライブラリの変更が必要な場合は、`lib`以下のファイルを変更し、`make generate`を実行します。
1. `in_*.txt`と対応する`out_*.txt`を追加することで、テストを追加できます。
1. `make test pkg=abc999/A`を実行すると、設問の入力例に応じたテストが実行されます。
1. `make submit pkg=abc999/A`を実行すると、コードを提出します。
    * submit前にtestが自動で実行されます。失敗した場合は提出を中止するので安心です。

## ディレクトリ構成

* `lib` 汎用的に利用するライブラリのコードをおきます。コードにgennyを利用した場合、そのままでは使えません。`make generate`を実行して実際に利用するコードを生成してください。
* `templates` `make new`で生成する`main.go`のテンプレートを置きます。
