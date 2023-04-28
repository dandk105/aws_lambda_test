package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func insertUsername(endpoint string, username string, password string, database string) {

	// データベースに接続
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, endpoint, database))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// データを挿入するためのクエリを作成
	query := "INSERT INTO users (id, name, email) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// パラメータを設定してクエリを実行
	_, err = stmt.Exec(1, "John Smith", "john@example.com")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Data inserted successfully")
}
