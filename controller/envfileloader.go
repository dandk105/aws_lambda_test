package controller

//このファイルは、.envファイルに記載された環境変数を読み込むためのものです。
//環境変数の読み込みには、godotenvというライブラリを使用しています。

import (
	"log"

	"github.com/joho/godotenv"
)

// 環境変数を読み込む関数を外部のファイルから呼び出すために、main関数を定義
func LoadEnvfile() {
	godotenv.Load()
	log.Println("Load .env file")
}
