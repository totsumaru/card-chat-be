# チャットの開始

ホストを設定します。

- ログインしているユーザーのみ

```
[POST] /api/chat/[chat-id]/start
```

### URL Params

なし

### Header

- `Authorization`(required)
- `Passcode`(required)

```text
Authorization: Bearer [token]
Passcode: [passcode]
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
    "message": "ログインしてください"
  }
}
```
