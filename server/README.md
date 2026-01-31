# Unitus Backend Server

Go で実装された高速で拡張性の高いバックエンドサーバー。

## セットアップ

```bash
go mod download
go run ./cmd/server/main.go
```

## ディレクトリ構成

```
cmd/
└── server/
    └── main.go                 # エントリーポイント

internal/
├── api/                        # API層
│   ├── routes.go
│   ├── handlers/
│   │   ├── task_handler.go
│   │   ├── doc_handler.go
│   │   ├── chat_handler.go
│   │   ├── ai_handler.go
│   │   ├── auth_handler.go
│   │   └── ...
│   └── middleware/
│       ├── auth.go
│       ├── cors.go
│       └── logging.go
├── domain/                     # ドメインモデル
│   ├── task.go
│   ├── document.go
│   ├── user.go
│   ├── team.go
│   ├── message.go
│   ├── ai_config.go
│   └── ...
├── usecase/                    # ビジネスロジック
│   ├── task_usecase.go
│   ├── doc_usecase.go
│   ├── chat_usecase.go
│   ├── ai_usecase.go
│   └── ...
├── repository/                 # データアクセス層
│   ├── task_repository.go
│   ├── doc_repository.go
│   ├── user_repository.go
│   └── ...
├── persistence/                # データベース
│   ├── db.go
│   ├── models/
│   └── migrations/
├── service/                    # 外部サービス統合
│   ├── ai_service.go
│   ├── notification_service.go
│   └── storage_service.go
└── config/
    └── config.go

pkg/
├── logger/
├── errors/
└── utils/
```

## API ドキュメント

詳細は [docs/API.md](../docs/API.md) を参照してください。

## ビルド

```bash
go build -o unitus-server ./cmd/server/main.go
```

## テスト

```bash
go test ./...
```

## Docker

```bash
docker build -t unitus-server .
docker run -p 8080:8080 unitus-server
```
