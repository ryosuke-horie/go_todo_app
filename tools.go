//go:build tools

package main

// go generateのバージョン管理のために利用するファイル
// ビルドタグを指定しない実アプリケーション開発時には無視される

// moqはモックにふるまいを指定する際に型を意識して実装できるツール
import _ "github.com/matryer/moq"
