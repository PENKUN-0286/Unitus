# Unitus - アーキテクチャ設計書

## システム概要

Unitusは、マルチテナント対応のクラウドベースタスク管理プラットフォームです。各チームが独立してサーバーをホストでき、複数のクライアント（Windows、Mac、Linux、Web）から接続可能です。

## 全体アーキテクチャ図

```
┌─────────────────────────────────────────┐
│         クライアント層（Dart/Flutter）    │
├──────────────────┬──────────────────────┤
│ Windows │ Mac │ Linux │ Web Browser     │
└──────────────────┼──────────────────────┘
                   │
                   │ REST API / WebSocket
                   │
┌──────────────────┴──────────────────────┐
│      サーバー層（Go）                      │
├─────────────────────────────────────────┤
│ ┌──────────────┐ ┌──────────────────┐  │
│ │ API Server   │ │ WebSocket Server │  │
│ │ (Gin)        │ │ (Gorilla)        │  │
│ └──────────────┘ └──────────────────┘  │
│ ┌──────────────────────────────────┐   │
│ │ Business Logic Layer             │   │
│ │ - Tasks, Docs, Goals, etc        │   │
│ └──────────────────────────────────┘   │
│ ┌──────────────────────────────────┐   │
│ │ Data Access Layer (GORM)         │   │
│ └──────────────────────────────────┘   │
└──────────────────┬──────────────────────┘
                   │
        ┌──────────┴──────────┐
        │                     │
┌───────▼────────┐  ┌────────▼──────────┐
│  Database      │  │  External Services│
│  (PostgreSQL)  │  │  - AI APIs        │
│                │  │  - Storage        │
└────────────────┘  └───────────────────┘
```

## クライアント層（Dart/Flutter）

### テクノロジースタック

- **Flutter**: UIフレームワーク
- **GetX**: 状態管理
- **Dio**: HTTP通信
- **Freezed**: イミュータブルデータクラス
- **Riverpod**: 依存性注入
- **Hive**: ローカルストレージ

### ディレクトリ構成

```
client/
├── lib/
│   ├── main.dart
│   ├── presentation/           # UI層
│   │   ├── pages/
│   │   ├── widgets/
│   │   ├── controllers/
│   │   └── themes/
│   ├── domain/                 # ビジネスロジック層
│   │   ├── entities/
│   │   ├── repositories/
│   │   └── usecases/
│   ├── data/                   # データ層
│   │   ├── datasources/
│   │   ├── models/
│   │   └── repositories/
│   ├── config/                 # 設定
│   │   ├── routes.dart
│   │   ├── theme.dart
│   │   └── constants.dart
│   └── core/                   # ユーティリティ
│       ├── network/
│       ├── storage/
│       └── utils/
├── pubspec.yaml
└── README.md
```

### 主要機能モジュール

1. **認証・ユーザー管理**
   - ログイン・登録
   - サーバー接続設定
   - プロファイル管理

2. **Task管理**
   - タスク一覧表示
   - タスク作成・編集・削除
   - ステータス更新
   - リアルタイム同期

3. **Docs（ドキュメント）**
   - リッチテキストエディタ
   - リアルタイムコラボレーション

4. **Chat**
   - メッセージング
   - リアルタイム通知

5. **Dashboard**
   - カスタマイズ可能なウィジェット
   - リアルタイムデータ表示

## サーバー層（Go）

### テクノロジースタック

- **Gin**: Webフレームワーク
- **GORM**: ORM
- **PostgreSQL**: データベース
- **JWT**: 認証トークン
- **Gorilla WebSocket**: リアルタイム通信
- **Redis**: キャッシング・セッション管理

### ディレクトリ構成

```
server/
├── cmd/
│   └── server/
│       └── main.go             # エントリーポイント
├── internal/
│   ├── api/                    # APIハンドラー
│   │   ├── routes.go
│   │   ├── handlers/
│   │   │   ├── task_handler.go
│   │   │   ├── doc_handler.go
│   │   │   ├── chat_handler.go
│   │   │   ├── ai_handler.go
│   │   │   └── ...
│   │   └── middleware/
│   │       ├── auth.go
│   │       ├── cors.go
│   │       └── logging.go
│   ├── domain/                 # ドメインモデル
│   │   ├── task.go
│   │   ├── document.go
│   │   ├── user.go
│   │   ├── team.go
│   │   └── ...
│   ├── usecase/                # ビジネスロジック
│   │   ├── task_usecase.go
│   │   ├── doc_usecase.go
│   │   ├── ai_usecase.go
│   │   └── ...
│   ├── repository/             # データアクセス層
│   │   ├── task_repository.go
│   │   ├── doc_repository.go
│   │   ├── user_repository.go
│   │   └── ...
│   ├── persistence/            # データベース
│   │   ├── db.go
│   │   ├── migrations/
│   │   └── models/
│   ├── service/                # 外部サービス統合
│   │   ├── ai_service.go
│   │   ├── notification_service.go
│   │   └── storage_service.go
│   └── config/
│       └── config.go
├── pkg/
│   ├── logger/
│   ├── errors/
│   └── utils/
├── go.mod
├── go.sum
├── .env.example
└── Dockerfile
```

### APIエンドポイント設計

#### Tasks
```
GET    /api/v1/teams/{teamId}/tasks
POST   /api/v1/teams/{teamId}/tasks
GET    /api/v1/teams/{teamId}/tasks/{taskId}
PUT    /api/v1/teams/{teamId}/tasks/{taskId}
DELETE /api/v1/teams/{teamId}/tasks/{taskId}
```

#### Documents
```
GET    /api/v1/teams/{teamId}/docs
POST   /api/v1/teams/{teamId}/docs
GET    /api/v1/teams/{teamId}/docs/{docId}
PUT    /api/v1/teams/{teamId}/docs/{docId}
DELETE /api/v1/teams/{teamId}/docs/{docId}
```

#### Chat
```
GET    /api/v1/teams/{teamId}/messages
POST   /api/v1/teams/{teamId}/messages
WS     /ws/teams/{teamId}/chat  # WebSocket
```

#### AI
```
POST   /api/v1/teams/{teamId}/ai/configure    # APIキー設定
GET    /api/v1/teams/{teamId}/ai/config
POST   /api/v1/teams/{teamId}/ai/query
DELETE /api/v1/teams/{teamId}/ai/config
```

## データベーススキーマ

### 主要テーブル

#### teams
```
- id (UUID, PK)
- name (VARCHAR)
- description (TEXT)
- created_at
- updated_at
```

#### users
```
- id (UUID, PK)
- email (VARCHAR, UNIQUE)
- password_hash (VARCHAR)
- name (VARCHAR)
- created_at
- updated_at
```

#### team_members
```
- id (UUID, PK)
- team_id (UUID, FK)
- user_id (UUID, FK)
- role (ENUM: admin, member, viewer)
- joined_at
```

#### tasks
```
- id (UUID, PK)
- team_id (UUID, FK)
- title (VARCHAR)
- description (TEXT)
- status (ENUM: todo, in_progress, done)
- priority (ENUM: low, medium, high, urgent)
- assigned_to (UUID, FK to users)
- due_date (TIMESTAMP)
- created_by (UUID, FK)
- created_at
- updated_at
```

#### documents
```
- id (UUID, PK)
- team_id (UUID, FK)
- title (VARCHAR)
- content (JSONB)
- version (INT)
- created_by (UUID, FK)
- created_at
- updated_at
```

#### messages
```
- id (UUID, PK)
- team_id (UUID, FK)
- channel (VARCHAR)
- content (TEXT)
- sender_id (UUID, FK)
- created_at
```

#### ai_configs
```
- id (UUID, PK)
- team_id (UUID, FK)
- provider (ENUM: gemini, chatgpt, custom)
- api_key_encrypted (VARCHAR)
- model_name (VARCHAR)
- created_at
- updated_at
```

## 認証フロー

1. **ログイン**
   - ユーザーがサーバーアドレスとクレデンシャルを入力
   - サーバーが JWT トークンを発行
   - クライアントがトークンをローカルに保存

2. **API リクエスト**
   - クライアントが全てのAPI呼び出しに JWT をヘッダーに含める
   - サーバーが JWT を検証
   - 有効な場合、リクエスト処理

3. **リアルタイム通信**
   - WebSocket 接続時に JWT で認証
   - 継続的な接続を維持

## AI統合アーキテクチャ

```
┌──────────────┐
│ クライアント   │
└──────┬───────┘
       │
       ▼
┌──────────────────────┐
│ Unitus Server        │
│ - AI Proxy Layer     │
│ - Rate Limiting      │
│ - Caching            │
└──────┬───────────────┘
       │
   ┌───┴───┐
   ▼       ▼
┌──────┐ ┌──────┐
│Google│ │OpenAI│
│Gemini│ │API   │
└──────┘ └──────┘
```

**特徴:**
- チームごとに異なるAI プロバイダーを設定可能
- API キーの暗号化保存
- レート制限の実装
- レスポンスキャッシング

## セキュリティ考慮事項

1. **認証**
   - JWT 使用
   - Refresh Token ローテーション

2. **データ保護**
   - API キーの暗号化
   - HTTPS/TLS 必須
   - SQL インジェクション対策

3. **アクセス制御**
   - ロールベースアクセス制御（RBAC）
   - チーム内権限管理

## デプロイメント

### Docker Compose による開発環境

```yaml
version: '3.8'
services:
  server:
    build: ./server
    ports:
      - "8080:8080"
    environment:
      DATABASE_URL: postgresql://user:pass@db:5432/unitus
    depends_on:
      - db
      - redis

  db:
    image: postgres:16
    environment:
      POSTGRES_DB: unitus
      POSTGRES_PASSWORD: password

  redis:
    image: redis:7-alpine
```

### 本番環境

- Kubernetes or Docker Swarm
- CI/CD パイプライン（GitHub Actions）
- 監視・ログ収集（Prometheus, ELK）

## パフォーマンス最適化

1. **キャッシング**
   - Redis キャッシング
   - クライアント側キャッシング

2. **データベース**
   - インデックス最適化
   - クエリ最適化

3. **リアルタイム同期**
   - WebSocket による効率的な通信
   - デルタ同期

---

詳細な実装については、各コンポーネントの README を参照してください。
