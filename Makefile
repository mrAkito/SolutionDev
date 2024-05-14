.PHONY: run down stop tidy migrate-file migrate-up migrate-down

# Docker Composeの起動（ログをターミナルに流すために[-d]は使用しない）
run:
	docker-compose up

stop:
	docker-compose stop

# Docker Composeの停止とボリュームの削除
down:
	docker-compose down -v

tidy:
	docker exec SolutionDev-app go mod tidy

# マイグレーション用ファイル作成
migrate-file:
	docker exec SolutionDev-app migrate create -ext sql -dir ./cmd/migrate/data -seq $(TableName)

migrate-up:
	docker exec SolutionDev-app migrate -database "mysql://mysql:mysql@tcp(SolutionDev-db:3306)/SolutionDB" -path ./cmd/migrate/data up

# 全てのダウンマイグレーションを実行するecho 'y'がついてるので注意
migrate-down:
	echo 'y' | docker-compose exec -T backend migrate -database "mysql://mysql:mysql@tcp(SolutionDev-db:3306)/SolutionDB" -path ./cmd/migrate/data down