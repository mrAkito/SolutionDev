# SolutionDev

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
# 例) migrate-file TableName="create_goal"

# マイグレートを実行
make migrate-up

# ダウンマイグレーションを実行
# 全てのダウンマイグレーションを実行する"- y"がついてるので注意
make migrate-down
```

## 挙動について
- コード上で新たに外部パッケージを必要とする際は、新たにシェルを立ち上げて上記に記載されてる方法で`make tidy`などをするとgo.modファイルが更新されます。
