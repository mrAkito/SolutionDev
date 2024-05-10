# try-standard-layout

## コマンド群
```sh
# コンテナ立ち上げ
make run

# ボリューム削除 (イメージは残ります)
make down

# 例) go mod tidy をコンテナ内で行いたい場合
make tidy

# マイグレーション用fileを作成
make migrate-file TableName="ここにテーブル名を記入"

# マイグレートを実行
make migrate-up

# ダウンマイグレーションを実行
# 全てのダウンマイグレーションを実行する"- y"がついてるので注意
make migrate-down

# 例) postgresに接続する場合
docker exec -it try-standard-layout-db psql -U mysql -d develop
```

## layout
```sh
/cmd
	main.go         # ここでインスタンス作成　例) router.New()
	/migrate
	  /data
	    migrate-up.sql    # golang-migrateにより自動生成されたマイグレーションファイル群
	    migrate-down.sql  # こっちはテーブル削除用
/internal
	/router
		router.go	# ここでechoインスタンスを作成するメソッドを実装 func New(){}
		group.go
	/server
		server.go	# serverインスタンスを作成するメソッドを実装 func New(){}
		product.go
	/database
		/request
		/response
			product.go
		/mysql
			mysql.go		#dbのconnectやcloseを行う
	/domain         # テーブルの構造を定義する
		product.go
```

## 挙動について
- コード上で新たに外部パッケージを必要とする際は、新たにシェルを立ち上げて上記に記載されてる方法でコンテナ内で`go mod tidy`などをするとgo.modファイルが更新されます。