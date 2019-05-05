package main

import (
	"github.com/labstack/echo"
	"net/http"
	"sampleAPI/dbConteroller"
	"strconv"
)

func main() {
	// 初期化
	e := echo.New()

	// ルーティング
	e.GET("/get", getHandler)
	e.POST("/post", postHandler)
	e.PUT("/put/:id", putHandler)
	e.DELETE("/delete/:id", deleteHandler)

	// サーバの起動
	e.Logger.Fatal(e.Start(":12345"))
}

// GETリクエスト
func getHandler(c echo.Context) error {
	// データベースからデータを取得
	data, err := dbConteroller.GetData()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

// POSTリクエスト
func postHandler(c echo.Context) error {
	// リクエストデータを取得
	var user dbConteroller.Post
	if err := c.Bind(&user); err != nil {
		return err
	}

	// データベースにデータを挿入
	err := dbConteroller.PostData(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success")

}

// PUTリクエスト
func putHandler(c echo.Context) error {
	// リクエストデータを取得
	var user dbConteroller.Post
	if err := c.Bind(&user); err != nil {
		return err
	}
	// 更新するデータのIDを取得
	req := c.Param("id")
	id, err := strconv.Atoi(req)
	if err != nil {
		return err
	}
	user.Id = id

	// データベース内の指定されたIDのデータを更新
	err = dbConteroller.PutData(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success")
}

// DELETEリクエスト
func deleteHandler(c echo.Context) error {
	// リクエストデータを取得
	req := c.Param("id")
	id, err := strconv.Atoi(req)
	if err != nil {
		return err
	}

	// データベース内の指定されたIDのデータを削除
	err = dbConteroller.DeleteData(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success")
}
