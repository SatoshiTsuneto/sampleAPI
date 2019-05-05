package dbConteroller

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

// データベースに接続する関数
func connectingDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:secret@/user_info")
	if err != nil {
		return
	}
	return
}

// テーブルの行を取得する関数
func getRow(db *sql.DB) (count int, err error) {
	rows, err := db.Query("SELECT COUNT(id) FROM posts")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		_ = rows.Scan(&count)
	}
	return
}

// テーブルからデータの取得
func GetData() (data []Post, e error) {
	// データベースに接続
	db, err := connectingDatabase()
	if err != nil {
		e = err
		return
	}
	defer db.Close()

	// データの取得
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		e = err
		return
	}
	defer rows.Close()

	// 取得したデータの格納
	var line Post
	for rows.Next() {
		_ = rows.Scan(&line.Id, &line.Name, &line.Age)
		data = append(data, line)
	}

	return
}

// テーブルにデータを挿入する関数
func PostData(user Post) error {
	// データベースに接続
	db, err := connectingDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	// 行の取得（IDに使う）
	row, err := getRow(db)
	if err != nil {
		return err
	}
	user.Id = row + 1

	// 準備
	statement := "INSERT INTO posts VALUES (?, ?, ?)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// データの挿入
	_, err = stmt.Exec(user.Id, user.Name, user.Age)
	if err != nil {
		return err
	}

	return nil
}

// テーブルのデータの更新
func PutData(post Post) error {
	// データベースに接続
	db, err := connectingDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	// 準備
	statement := "UPDATE posts SET name = ?, age = ? WHERE id = ?"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// データの更新
	_, err = stmt.Exec(post.Name, post.Age, post.Id)
	if err != nil {
		return err
	}

	return nil
}

// テーブルからデータを削除
func DeleteData(id int) error {
	// データベースに接続
	db, err := connectingDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	// 準備
	statement := "DELETE FROM posts WHERE id = ?"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// データの更新
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
