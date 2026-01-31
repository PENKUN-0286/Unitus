# Unitus Client

Flutter ベースのマルチプラットフォーム クライアント。

## プラットフォーム対応

- ✅ Windows
- ✅ macOS
- ✅ Linux
- ✅ Web

## セットアップ

```bash
flutter pub get
flutter run
```

## ディレクトリ構成

```
lib/
├── main.dart                   # エントリーポイント
├── presentation/              # UI層
│   ├── pages/
│   ├── widgets/
│   ├── controllers/
│   └── themes/
├── domain/                     # ビジネスロジック層
│   ├── entities/
│   ├── repositories/
│   └── usecases/
├── data/                       # データ層
│   ├── datasources/
│   ├── models/
│   └── repositories/
├── config/                     # 設定
│   ├── routes.dart
│   └── theme.dart
└── core/                       # ユーティリティ
    ├── network/
    ├── storage/
    └── utils/
```

## 主要機能

- 🔐 サーバー接続・認証
- 📋 タスク管理
- 📄 ドキュメント管理
- 🎯 目標管理
- 🎨 ホワイトボード
- 📊 ダッシュボード
- 💬 チャット
- 🤖 AI アシスタント

## ビルド

### Web
```bash
flutter run -d chrome
```

### Windows
```bash
flutter run -d windows
```

### macOS
```bash
flutter run -d macos
```

### Linux
```bash
flutter run -d linux
```

## テスト

```bash
flutter test
```
