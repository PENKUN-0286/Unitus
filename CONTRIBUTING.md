# Unitus Contributing Guide

Unitus プロジェクトへの貢献を歓迎します！

## コード規約

### Go
- [Effective Go](https://golang.org/doc/effective_go)に従う
- `golangci-lint` でチェック
- テストカバレッジ: 最小 80%

### Dart/Flutter
- [Dart Style Guide](https://dart.dev/guides/language/effective-dart/style)に従う
- `dart analyze` で確認
- [Flutter Best Practices](https://flutter.dev/docs/testing)

## 貢献フロー

1. **Fork する**
   ```bash
   git clone https://github.com/your-username/unitus.git
   ```

2. **フィーチャーブランチを作成**
   ```bash
   git checkout -b feature/my-feature
   ```

3. **変更を加える**
   - コードを実装
   - テストを追加
   - ドキュメントを更新

4. **テストを実行**
   ```bash
   # Go
   go test ./...
   golangci-lint run
   
   # Flutter
   flutter analyze
   flutter test
   ```

5. **コミットする**
   ```bash
   git commit -m "[feat]: Add new feature"
   git push origin feature/my-feature
   ```

6. **Pull Request を作成**
   - 明確な説明を記入
   - 関連するIssueをリンク
   - スクリーンショットを添付（UI変更の場合）

## コミットメッセージ形式

```
[タイプ]: 短い説明（50文字以内）

より詳細な説明があればここに記入。
複数行に渡る場合も OK。

- リスト形式で詳細を記入
- 理由や背景も説明すること

Fixes #123
Related #456
```

### タイプ

- `feat`: 新機能
- `fix`: バグ修正
- `docs`: ドキュメント
- `style`: コード整形
- `refactor`: リファクタリング
- `test`: テスト追加・修正
- `chore`: ビルド・依存性更新

## Pull Request チェックリスト

- [ ] ブランチが最新の `develop` にマージ可能
- [ ] ユニットテストが実行されている
- [ ] コードが Style Guide に従っている
- [ ] ドキュメントが更新されている
- [ ] コミットメッセージが明確
- [ ] 1つのPRは1つの機能/修正のみ

## コード Review プロセス

1. 最低2人の Reviewer による承認が必要
2. CI/CD パイプラインが成功
3. Conflicts がない
4. コードカバレッジが維持

## 開発環境のセットアップ

```bash
# リポジトリをクローン
git clone https://github.com/your-username/unitus.git
cd unitus

# 環境変数を設定
cp .env.example .env

# Docker で開発環境を起動
docker-compose up -d

# 依存関係をインストール
cd server && go mod download
cd ../client && flutter pub get
```

## サポート

- 質問: GitHub Discussions を使用
- バグ報告: GitHub Issues を作成
- セキュリティ: security@unitus.dev に報告

---

Thank you for contributing! 🙏
