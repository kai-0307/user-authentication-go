### PostgreSQL がインストールされているか確認

```
postgres --version
```

### もしインストールされていない場合（Mac の場合）

```
brew install postgresql@14
```

### Homebrew でインストールした場合（Mac）

```
brew services start postgresql
```

または

```
pg_ctl -D /usr/local/var/postgres start
```

### PostgreSQL に接続

```
psql postgres
```

### データベース作成

```
CREATE DATABASE testdb;
```

### ユーザーの確認（必要な場合）

```
\du
```

### データベースの確認

```
\l
```

### サービスの状態確認（Mac）

```
brew services list
```

または

```
pg_ctl status -D /usr/local/var/postgres
```

### データベースの URL を環境に合わせて設定

```
export POSTGRESQL_URL='postgres://postgres:your_password@localhost:5432/testdb?sslmode=disable'
```

### マイグレーションを実行

```
migrate -database ${POSTGRESQL_URL} -path backend/migrations up
```

### データベースに接続

```
psql testdb
```

### テーブル一覧を表示

```
\dt
```

### users テーブルの構造を確認

```
\d users
```

### ヘルスチェック

```
$ curl http://localhost:8080/api/health
{"status":"ok"}%
```

### ユーザー登録

```
curl -X POST http://localhost:8080/api/auth/register \
 -H "Content-Type: application/json" \
 -d '{
"username": "testuser",
"email": "test@example.com",
"password": "password123"
}'
```

成功してユーザーが作成され、以下の情報が返却される想定

UUID: "f89648fd-e45e-429b-b202-f2beb9f47851"
ユーザー名: "testuser"
メール: "test@example.com"
タイムスタンプ情報も正しく記録

### 登録したユーザーでログイン

```
curl -X POST http://localhost:8080/api/auth/login \
 -H "Content-Type: application/json" \
 -d '{
"email": "test@example.com",
"password": "password123"
}'
```

成功して JWT トークンが発行

### 取得したトークンを使ってプロフィールを取得

```
curl http://localhost:8080/api/user/profile \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZjg5NjQ4ZmQtZTQ1ZS00MjliLWIyMDItZjJiZWI5ZjQ3ODUxIiwidXNlcm5hbWUiOiJ0ZXN0dXNlciIsImV4cCI6MTczNzY0Nzk1MiwibmJmIjoxNzM3NTYxNTUyLCJpYXQiOjE3Mzc1NjE1NTJ9.8E9dm-DAq6JKt_D9Wd-8C23oycA36tvEPZ8uJhwoxJY"
```

返ってくる値

```
{"message":"get profile","user_id":"f89648fd-e45e-429b-b202-f2beb9f47851"}%
```
