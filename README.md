# golang REST API boilerplate
golang製 REST API ボイラーテンプレート

# 技術構成
- mysql: 8.0.28
- go: 1.19.2
- net/http
- gorm
- sql-migrate
- gorilla/sessions
- github actions

# 仕様
TODOリストAPI
- 認証機能は UserId をセッションに保存

## ルーティング

### トップ
|   | メソッド | URI | 認証有無 |
| - | ------- | --- | ------ |
| - | GET | / | なし |

### ユーザー認証API
|   | メソッド | URI | 認証有無 | クエリ | リクエストボディ |
| - | ------- | --- | ------ | - | - |
| ログイン | POST | /auth/sign_in/ | なし | | [AuthSignInRequestDto](./controller/dto/auth.dto.go#L3) |
| 会員登録 | POST | /auth/sign_up/ | なし |  |[AuthSignUpRequestDto](./controller/dto/auth.dto.go#L8) |

### ログインユーザーAPI
|   | メソッド | URI | 認証有無 | クエリ | リクエストボディ |
| - | ------- | --- | ------ | - | - |
| ログインユーザー情報取得 | GET | /auth/user/ | あり | | 
| ユーザー更新 | PUT | /auth/user/ | あり | | [AuthUserUpdateRequestDto](./controller/dto/auth.dto.go#L14) |
| ユーザー削除 | DELETE | /auth/user/ | あり | | |

### TodoAPI
|   | メソッド | URI | 認証有無 | クエリ | リクエストボディ |
| - | ------- | --- | ------ | - | - |
| ユーザーに紐づく全Todoデータ取得 | GET | /todos/ | あり | | |
| 単一Todoデータ取得 | POST | /todos/id/ | あり | | |
| Todo新規作成 | POST | /todos/ | あり | | [TodoCreateRequestDto](./controller/dto/todo.dto.go#L3) |
| Todo更新 | PUT | /todos/id/ | あり | | [TodoCreateRequestDto](./controller/dto/todo.dto.go#L8) |
| Todo削除 | DELETE | /todos/id/ | あり |

## ミドルウェア対応パス表
| URI | [Cors](./middleware/auth.middleware.go) | [WriteHeader](./middleware/write_header.middleware.go) | [Auth](./middleware/auth.middleware.go) | [SetHttpContext](./middleware/set_http_context.middleware.go) |
|-|-|-|-|-|
|/|x|x|x|x|
|/toods/*/|○|○|○|○|
|/auth/sign_in/|○|○|x|○|
|/auth/sign_up/|○|○|x|○|
|/auth/user/*/|○|○|○|○|

# 環境構築
## 1. ルートディレクトリに「.env」ファイルの用意
```
touch .env
```
下記環境変数をセット
|変数名|説明|
|----|----|
|GO_MODE|稼働環境|
|DB_USER|DBユーザー|
|DB_PASSWORD|DBユーザーパスワード|
|DB_HOST|DBホスト|
|DB_NAME|データベース名|
|SECRET_HASH_KEY|セッションキー|
|MAIL_AUTH_EMAIL|メール送信アカウントのメールアドレス|
|MAIL_AUTH_PASSWORD|メール送信アカウントのパスワード|
|MAIL_FROM_NAME|メール送信元名|
|MAIL_FROM_EMAIL|メール送信元のアドレス|

## 2. イメージビルドとDB作成・マイグレーション
```
make setup
```

## 3. APIスタート
```
make start
```


# Makefile
### setup
イメージビルドとDB作成・マイグレーション

### db.create
DB作成

### db.migrate
DBマイグレーション

### db.seed
シードデータ投入

### start
APIスタート（Dockerコンテナ起動）

### end
APIストップ（Dockerコンテナ停止）

### entry-server-container
Dockerサーバーコンテナに ash で入る

### entry-db-container
Docker DBコンテナに bash で入る