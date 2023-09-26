# 既読処理

チャットの既読処理をします。

- ホストのみ

```
[POST] /api/chat/[chat-id]/read
```

### URL Params

なし

### Header

- `Authorization`(required)

```text
Authorization: Bearer [token]
```

### Body

なし

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
