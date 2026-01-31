# Unitus - オープンソースタスク管理プラットフォーム

Unitusは、現代的なタスク管理アプリケーションです。複数のチームをサポートし、各チームが独立したサーバーを運用できます。

## 🎯 特徴

- **マルチプラットフォーム対応**: Windows、Mac、Linux、Web
- **チーム別サーバー対応**: 各チームが独立してサーバーをホスト可能
- **豊富な機能セット**:
  - 📋 Tasks - タスク管理
  - 📄 Docs - ドキュメント管理
  - 🎯 Goals - 目標管理
  - 🎨 Whiteboards - ホワイトボード
  - 📊 Dashboards - ダッシュボード
  - 💬 Chat View - チャット機能
  - 🤖 AI Assistant - カスタマイズ可能なAI統合（Gemini、ChatGPT対応）
- **モダンでスケーラブル**: Go + Dartで構築された高速で保守性の高い設計

## 📋 プロジェクト構成

```
unitus/
├── client/              # Flutterクライアント（マルチプラットフォーム）
├── server/              # Goバックエンド
├── docs/                # ドキュメント
├── docker-compose.yml   # 開発環境のセットアップ
└── README.md
```

## 🚀 クイックスタート

### 前提条件

- Dart SDK 3.0+
- Go 1.21+
- Docker & Docker Compose

### クライアント側のセットアップ

```bash
cd client
flutter pub get
flutter run
```

### サーバー側のセットアップ

```bash
cd server
go mod download
go run ./cmd/server/main.go
```

## 🏗️ アーキテクチャ

### クライアント（Dart + Flutter）
- **Flutter** - マルチプラットフォームUIフレームワーク
- **GetX** - 状態管理
- **Dio** - HTTP通信
- **モダンなUI/UX** - Material Design 3ベース

### サーバー（Go）
- **Gin** - Webフレームワーク
- **GORM** - ORM
- **JWT** - 認証
- **WebSocket** - リアルタイム通信

## 📱 機能詳細

### Tasks（タスク管理）
- タスクの作成・編集・削除
- 優先度、期限の設定
- ステータス管理
- タグ・カテゴリ分類
- タスク検索・フィルタリング

### Docs（ドキュメント）
- リッチテキストエディタ
- コラボレーション編集
- バージョン管理
- フォルダ体系

### Goals（目標管理）
- 長期目標の設定
- マイルストーン管理
- 進捗トラッキング

### Whiteboards（ホワイトボード）
- リアルタイム共有ホワイトボード
- 図形描画、テキスト
- チーム間コラボレーション

### Dashboards（ダッシュボード）
- カスタマイズ可能なダッシュボード
- リアルタイムデータ表示
- メトリクス・チャート表示

### Chat View（チャット）
- チームメッセージング
- 1対1チャット
- スレッド機能
- ファイル共有

### AI Assistant（AI統合）
- チームごとのカスタマイズ可能なAPI
- Gemini、ChatGPT対応
- タスク提案、ドキュメント作成支援など
- API キー管理インターフェース

## 🔐 セキュリティ

- JWT ベースの認証
- エンドツーエンド暗号化対応
- API認証キーの安全な管理

## 📝 開発ガイド

詳細な開発ガイドは [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md) を参照してください。

## 🤝 貢献

このプロジェクトはオープンソースです。貢献は大歓迎です！

## 📄 ライセンス

MIT License

## 📞 サポート

問題が発生した場合は、GitHubのIssueを作成してください。

---

**Made with ❤️ for teams that want to organize together**
