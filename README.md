## Go言語のフレームワーク「Echo」を用いたAPIのサンプル

### Go言語 version: 1.12.4

### フレームワーク: Echo(v4.1.5)

### データベース: MySQL(Ver 8.0.13 for Linux on x86_64)

### 実行環境: Ubuntu19.04

### プログラム内容
- GET   ：データベースからデータを取得
- POST  ：受け取ったJSONデータをデータベースに保存
- PUT   ：指定されたIDのデータを受け取ったJSONデータに更新
- DELETE：指定されたIDのデータを削除

### セットアップ
- > $ go get github.com/go-sql-driver/mysql
- > $ go get github.com/labstack/echo

### テスト
- データの取得
  > $ curl -X GET http://localhost:12345/get

- データの挿入
  > $ curl -H "Accept: application/json" -H "Content-type: application/json" -X POST http://localhost:12345/post -d '{"Name": "sample", "Age": 9999}'   

- データの更新
  > $ curl -H "Accept: application/json" -H "Content-type: application/json" -X PUT http://localhost:12345/put/1 -d '{"Name": "sample01", "Age": 20}'

- データの削除
  > $ curl -X DELETE http://localhost:12345/delete/1