# 開発ガイド

## 環境セットアップ

### 必要なツール

1. **Go** (v1.21+)
2. **Dart** (v3.0+)
3. **Flutter** (latest)
4. **PostgreSQL** (v14+)
5. **Redis** (v7+)
6. **Docker** & **Docker Compose**
7. **Git**

### インストール手順

#### Windows

```powershell
# Go のインストール
# https://golang.org/dl より Go 1.21+ をダウンロード・インストール

# Dart/Flutter のインストール
# https://flutter.dev/docs/get-started/install より Flutter をダウンロード

# PostgreSQL のインストール
# https://www.postgresql.org/download/windows/

# Docker Desktop のインストール
# https://www.docker.com/products/docker-desktop
```

#### macOS

```bash
# Homebrew を使用
brew install go dart postgresql redis

# Flutter のインストール
git clone https://github.com/flutter/flutter.git -b stable
export PATH="$PATH:`pwd`/flutter/bin"
```

#### Linux (Ubuntu/Debian)

```bash
# Go のインストール
sudo apt-get install golang-go

# Dart のインストール
sudo apt-get install dart

# Flutter のインストール
git clone https://github.com/flutter/flutter.git -b stable
export PATH="$PATH:`pwd`/flutter/bin"

# PostgreSQL, Redis
sudo apt-get install postgresql redis-server
```

## クライアント開発

### プロジェクト初期化

```bash
cd client
flutter create . --platforms windows,macos,linux,web
flutter pub get
```

### フォルダ構成を構築

```bash
# lib/ 配下のフォルダを作成
mkdir -p lib/{presentation,domain,data,config,core}
mkdir -p lib/presentation/{pages,widgets,controllers,themes}
mkdir -p lib/domain/{entities,repositories,usecases}
mkdir -p lib/data/{datasources,models,repositories}
mkdir -p lib/config/{routes,themes}
mkdir -p lib/core/{network,storage,utils}
```

### 主要パッケージをインストール

[client/pubspec.yaml](../client/pubspec.yaml) に以下を追加：

```yaml
dependencies:
  flutter:
    sdk: flutter
  
  # State Management
  riverpod: ^2.4.0
  flutter_riverpod: ^2.4.0
  
  # HTTP & Networking
  dio: ^5.3.0
  
  # Data Serialization
  freezed_annotation: ^2.4.1
  json_serializable: ^6.7.1
  
  # Local Storage
  hive: ^2.2.3
  hive_flutter: ^1.1.0
  
  # Routing
  go_router: ^10.0.0
  
  # UI Components
  flutter_animate: ^4.1.1
  gap: ^3.0.1
  
  # Utilities
  uuid: ^4.0.0
  intl: ^0.19.0

dev_dependencies:
  flutter_test:
    sdk: flutter
  flutter_lints: ^3.0.0
  build_runner: ^2.4.6
  freezed: ^2.4.1
```

### ビルド・実行

```bash
# Web
flutter run -d chrome

# Windows
flutter run -d windows

# macOS
flutter run -d macos

# Linux
flutter run -d linux
```

## サーバー開発

### プロジェクト初期化

```bash
cd server
go mod init github.com/your-org/unitus
```

### go.mod ファイル

```go
module github.com/your-org/unitus

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/gorm.io/gorm v1.25.5
    github.com/gorm.io/driver/postgres v1.5.7
    github.com/golang-jwt/jwt/v5 v5.0.0
    github.com/google/uuid v1.5.0
    github.com/gorilla/websocket v1.5.0
    github.com/joho/godotenv v1.5.1
    github.com/redis/go-redis/v9 v9.3.0
    golang.org/x/crypto v0.17.0
)
```

### ディレクトリ構成を作成

```bash
mkdir -p cmd/server
mkdir -p internal/{api,domain,usecase,repository,persistence,service,config}
mkdir -p internal/api/{handlers,middleware}
mkdir -p internal/persistence/{models,migrations}
mkdir -p pkg/{logger,errors,utils}
```

### main.go ファイル作成

```go
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/your-org/unitus/internal/config"
	"github.com/your-org/unitus/internal/api"
	"github.com/your-org/unitus/internal/persistence"
)

func init() {
	godotenv.Load()
}

func main() {
	cfg := config.LoadConfig()

	db, err := persistence.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := api.SetupRouter(db)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
```

### ビルド・実行

```bash
# 開発モード
go run ./cmd/server/main.go

# ビルド
go build -o unitus-server ./cmd/server/main.go

# 実行
./unitus-server
```

## データベースセットアップ

### PostgreSQL 初期化

```bash
# Windows (PostgreSQL がインストール済みの場合)
psql -U postgres -c "CREATE DATABASE unitus;"
psql -U postgres -d unitus -c "CREATE USER unitus WITH PASSWORD 'password';"
psql -U postgres -d unitus -c "GRANT ALL PRIVILEGES ON DATABASE unitus TO unitus;"
```

### マイグレーション

サーバー内で GORM オートマイグレーション機能を使用：

```go
// internal/persistence/db.go
type Migration struct {
	db *gorm.DB
}

func (m *Migration) Migrate() error {
	return m.db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.TeamMember{},
		&models.Task{},
		&models.Document{},
		&models.Message{},
		&models.AIConfig{},
		// ...
	)
}
```

## 開発フロー

### ブランチ戦略

```
main (本番)
├── develop (開発)
│   ├── feature/tasks-management
│   ├── feature/real-time-sync
│   └── feature/ai-integration
```

### コミットメッセージ

```
[機能]: 短い説明

より詳細な説明があればここに記入

Closes #123
```

### コード品質

#### Go
```bash
# Lint
golangci-lint run ./...

# Format
gofmt -s -w .

# Test
go test ./...
```

#### Dart/Flutter
```bash
# Analyze
dart analyze

# Format
dart format .

# Test
flutter test
```

## ローカル開発環境（Docker Compose）

### docker-compose.yml

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_DB: unitus
      POSTGRES_USER: unitus
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

  server:
    build:
      context: ./server
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      DATABASE_URL: postgresql://unitus:password@postgres:5432/unitus
      REDIS_URL: redis://redis:6379
      JWT_SECRET: your-secret-key
      ENVIRONMENT: development
    depends_on:
      - postgres
      - redis
    volumes:
      - ./server:/app

volumes:
  postgres_data:
```

### 起動

```bash
# イメージをビルド
docker-compose build

# コンテナを起動
docker-compose up -d

# ログを確認
docker-compose logs -f server

# 停止
docker-compose down
```

## テスティング

### Go テスト

```go
// internal/usecase/task_usecase_test.go
package usecase

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	// Arrange
	usecase := NewTaskUsecase(mockRepo)
	
	// Act
	task, err := usecase.CreateTask(context.Background(), taskInput)
	
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, task)
}
```

### Flutter テスト

```dart
// test/presentation/pages/task_page_test.dart
import 'package:flutter_test/flutter_test.dart';

void main() {
	group('TaskPage', () {
		testWidgets('TaskPage displays tasks', (WidgetTester tester) async {
			// Build and render
			await tester.pumpWidget(const MyApp());
			
			// Verify UI elements
			expect(find.text('Tasks'), findsOneWidget);
		});
	});
}
```

## デバッギング

### Go

```bash
# Delve デバッガを使用
go install github.com/go-delve/delve/cmd/dlv@latest
dlv debug ./cmd/server/main.go
```

### Flutter

```bash
# DevTools を使用
flutter pub global activate devtools
flutter pub global run devtools
```

## パフォーマンスプロファイリング

### Go

```bash
# CPU プロファイル
go tool pprof http://localhost:8080/debug/pprof/profile

# メモリプロファイル
go tool pprof http://localhost:8080/debug/pprof/heap
```

### Flutter

DevTools の Performance タブを使用。

---

詳細については、各コンポーネントのドキュメントを参照してください。
