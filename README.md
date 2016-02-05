cagrep
====

cagrepは、cacheされた情報を検索するツールです。
検索対象をcacheしておくため、次回以降の検索が非常に高速となります。

といいたいところですが、現時点でスピードは出ていません。

## Description

以下の動作を行います。

 1. 指定フォルダ以下を検索しcacheする(cagrep --server)
 2. cagrepで検索する

## VS.
grepと比較すると以下の特徴があります。
静的検査など、繰り返し何度も検索する用途では特に有効です。

 * 良い所
  * 先にcacheを作るため2回目以降の検索が高速
 * 悪い所
  * on memory cacheのため、メモリ消費量が大きい

## Requirement

特になし

## Usage

    Usage: cagrep [OPTION]... PATTERN
    Search for PATTERN
    Example: cagrep 'hello world'

    -s, --server              serverモードで起動する
    -p, --port                serverのport番号を指定する

以下のようにコマンドプロンプトから実行できます。

    $ cagrep --server .
    README.md
    cagrep-client.go
    cagrep-server.go
    cagrep.go

その後、別のコマンドプロンプトから以下のように入力することで検索結果が表示されます。

    $ cagrep import
    cagrep-client.go:3:import (
    cagrep-server.go:3:import (
    cagrep.go:3:import (

詳細はヘルプコマンドを実行してください。

    $ cagrep  --help


## Install

TBD.

## Build

本ツールは、golangで作成されています。

    $ go version
    go version go1.5.3 windows/amd64

以下のコマンドで実行体を作成することができます。

    $ go build

なお、以下のライブラリをインストールする必要があります。

    $ go get github.com/codegangsta/cli
    $ go get github.com/hoisie/web
    $ go get golang.org/x/text/transform
    $ go get golang.org/x/text/encoding/japanese


## Contribution

TBD.

## Licence

[MIT](http://opensource.org/licenses/mit-license.php)

## Author

[sago35](https://github.com/sago35)
