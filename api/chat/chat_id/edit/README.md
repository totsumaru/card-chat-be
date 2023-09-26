# チャット情報の変更

チャットの情報を変更します。

- ホストのみ

```
[POST] /api/chat/[chat-id]/edit
```

### URL Params

なし

### Header

- `Authorization`(required)

```text
Authorization: Bearer [token]
```

### Body

```json
{
  "id": "cb273580-8a04-4421-8141-e2bc48a89069",
  "display_name": "鈴木 太郎"
}
```

### Success

- code: `200`

```text
レスポンスBodyなし
```

### Error

- code: `401` | `404` | `500`

```json
{
  "error": {
    "message": "ホストではありません"
  }
}
```
