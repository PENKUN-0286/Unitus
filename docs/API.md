# Unitus API リファレンス

## 基本情報

- **ベースURL**: `http://localhost:8080/api/v1`
- **認証**: JWT Bearer Token
- **Content-Type**: `application/json`

## 認証エンドポイント

### ログイン

```
POST /auth/login
```

**リクエスト:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**レスポンス:**
```json
{
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "name": "User Name"
  },
  "token": "jwt-token",
  "refresh_token": "refresh-token"
}
```

### ユーザー登録

```
POST /auth/signup
```

**リクエスト:**
```json
{
  "email": "user@example.com",
  "password": "password123",
  "name": "User Name"
}
```

**レスポンス:**
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "name": "User Name",
  "created_at": "2024-01-01T00:00:00Z"
}
```

## タスク管理 API

### タスク一覧取得

```
GET /teams/{teamId}/tasks?page=1&limit=20&status=todo
```

**クエリパラメータ:**
- `page` (int): ページ番号 (デフォルト: 1)
- `limit` (int): 件数 (デフォルト: 20)
- `status` (string): ステータス (todo, in_progress, done)
- `assigned_to` (uuid): 割り当て先ユーザーID

**レスポンス:**
```json
{
  "data": [
    {
      "id": "uuid",
      "title": "Task Title",
      "description": "Task Description",
      "status": "todo",
      "priority": "high",
      "assigned_to": "uuid",
      "due_date": "2024-12-31T23:59:59Z",
      "created_by": "uuid",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 100,
    "total_pages": 5
  }
}
```

### タスク作成

```
POST /teams/{teamId}/tasks
```

**リクエスト:**
```json
{
  "title": "New Task",
  "description": "Task Description",
  "priority": "high",
  "assigned_to": "uuid",
  "due_date": "2024-12-31T23:59:59Z"
}
```

**レスポンス:** (201 Created)
```json
{
  "id": "uuid",
  "title": "New Task",
  "status": "todo",
  "priority": "high",
  "created_at": "2024-01-01T00:00:00Z"
}
```

### タスク更新

```
PUT /teams/{teamId}/tasks/{taskId}
```

**リクエスト:**
```json
{
  "title": "Updated Task",
  "status": "in_progress",
  "priority": "medium"
}
```

### タスク削除

```
DELETE /teams/{teamId}/tasks/{taskId}
```

## ドキュメント管理 API

### ドキュメント一覧

```
GET /teams/{teamId}/docs
```

### ドキュメント作成

```
POST /teams/{teamId}/docs
```

**リクエスト:**
```json
{
  "title": "Document Title",
  "content": {
    "blocks": [
      {
        "type": "paragraph",
        "text": "Document content"
      }
    ]
  }
}
```

### ドキュメント更新

```
PUT /teams/{teamId}/docs/{docId}
```

### ドキュメント削除

```
DELETE /teams/{teamId}/docs/{docId}
```

## チャット API

### メッセージ取得

```
GET /teams/{teamId}/messages?channel=general&limit=50
```

### メッセージ送信

```
POST /teams/{teamId}/messages
```

**リクエスト:**
```json
{
  "channel": "general",
  "content": "Message content"
}
```

### WebSocket 接続

```
WS /ws/teams/{teamId}/chat?token=jwt-token
```

## AI API

### AI 設定確認

```
GET /teams/{teamId}/ai/config
```

**レスポンス:**
```json
{
  "id": "uuid",
  "provider": "gemini",
  "model_name": "gemini-pro",
  "configured": true
}
```

### AI 設定変更

```
POST /teams/{teamId}/ai/configure
```

**リクエスト:**
```json
{
  "provider": "gemini",
  "api_key": "your-api-key",
  "model_name": "gemini-pro"
}
```

### AI クエリ実行

```
POST /teams/{teamId}/ai/query
```

**リクエスト:**
```json
{
  "prompt": "Create a task from this requirement..."
}
```

**レスポンス:**
```json
{
  "response": "AI response text",
  "provider": "gemini",
  "usage": {
    "input_tokens": 100,
    "output_tokens": 50
  }
}
```

### AI 設定削除

```
DELETE /teams/{teamId}/ai/config
```

## ダッシュボード API

### ダッシュボード取得

```
GET /teams/{teamId}/dashboards/{dashboardId}
```

### ダッシュボード作成

```
POST /teams/{teamId}/dashboards
```

## エラーレスポンス

```json
{
  "error": "Error message",
  "status": 400,
  "timestamp": "2024-01-01T00:00:00Z"
}
```

### ステータスコード

- `200` OK
- `201` Created
- `400` Bad Request
- `401` Unauthorized
- `403` Forbidden
- `404` Not Found
- `500` Internal Server Error

---

詳細情報については、各ハンドラーのドキュメントを参照してください。
