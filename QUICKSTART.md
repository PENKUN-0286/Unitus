# クイックスタートガイド - Unitus

Unitusをローカルマシンで実行するための手順です。

## 📋 前提条件

以下がインストール済みであることを確認してください：

- **Go** v1.21 以上
- **Dart/Flutter** 最新版
- **Docker** & **Docker Compose**
- **Git**

## 🚀 セットアップ手順

### 1. リポジトリのクローン

```bash
cd c:\Users\Gaming
git clone https://github.com/your-org/unitus.git
cd Unitus
```

### 2. 環境変数の設定

```bash
# Windows (PowerShell)
Copy-Item .env.example .env

# Linux/macOS
cp .env.example .env
```

`.env` ファイルを必要に応じて編集してください。

### 3. Docker環境の起動

```bash
docker-compose up -d
```

このコマンドで以下が起動します：
- PostgreSQL (ポート 5432)
- Redis (ポート 6379)
- Go サーバー (ポート 8080)
- PgAdmin (ポート 5050)
- Redis Commander (ポート 8081)

### 4. データベースの初期化（初回のみ）

```bash
# サーバーコンテナにアクセス
docker-compose exec server sh

# マイグレーション実行
go run ./cmd/server/main.go migrate

# コンテナを出る
exit
```

### 5. クライアントのセットアップ

#### 別のターミナルウィンドウで実行：

```bash
cd client
flutter pub get
```

## 🎮 アプリケーション起動

### サーバーの確認

```bash
# サーバーが起動しているか確認
curl http://localhost:8080/health

# レスポンス例：
# {"status":"ok","server":"Unitus v1.0.0"}
```

### クライアントの起動

#### Web版

```bash
cd client
flutter run -d chrome
```

#### Windows版

```bash
cd client
flutter run -d windows
```

#### macOS版

```bash
cd client
flutter run -d macos
```

#### Linux版

```bash
cd client
flutter run -d linux
```

## 🔑 ログイン情報

開発環境用のテストアカウント：

```
メールアドレス: dev@unitus.local
パスワード: dev123456
```

## 🛠️ 開発ツール

### データベース管理

**PgAdmin にアクセス:**
```
URL: http://localhost:5050
ユーザー: admin@unitus.local
パスワード: password
```

### Redisの確認

**Redis Commander:**
```
URL: http://localhost:8081
```

### API テスト

```bash
# ログイン
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "dev@unitus.local",
    "password": "dev123456"
  }'

# レスポンスから token を取得
```

## 📱 マルチプラットフォーム対応状況

| プラットフォーム | ステータス | 起動コマンド |
|-----------------|-----------|-----------|
| Web | ✅ 実装中 | `flutter run -d chrome` |
| Windows | ✅ 実装中 | `flutter run -d windows` |
| macOS | ✅ 実装中 | `flutter run -d macos` |
| Linux | ✅ 実装中 | `flutter run -d linux` |

## 🐛 トラブルシューティング

### ポートが既に使用されている

```bash
# Windows (PowerShell)
netstat -ano | findstr :8080

# Linux/macOS
lsof -i :8080
```

### Docker が起動しない

```bash
# ログを確認
docker-compose logs server

# コンテナをリセット
docker-compose down -v
docker-compose up -d
```

### Flutter ビルドエラー

```bash
cd client
flutter clean
flutter pub get
flutter run
```

### データベース接続エラー

```bash
# PostgreSQL が起動しているか確認
docker-compose logs postgres

# 再起動
docker-compose restart postgres
```

## 📚 その他のドキュメント

- [アーキテクチャ設計書](./docs/ARCHITECTURE.md)
- [開発ガイド](./docs/DEVELOPMENT.md)
- [API リファレンス](./docs/API.md)
- [UI/UX ガイドライン](./docs/UI_UX_GUIDE.md)

## 🤝 貢献

改善案やバグ報告は GitHub Issues で！

## 📞 サポート

問題が発生した場合：

1. ドキュメントを確認
2. ログを確認
3. GitHub Issues を検索
4. 新しい Issue を作成

---

**Happy coding! 🎉**
