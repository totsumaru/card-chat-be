# チャットの開始

ホストを設定します。

- ログインしているユーザーのみが設定できます

```
[POST] /api/chat/[chat-id]/start
```

### URL Params

なし

### Header

- `Authorization`(required)

```text
Authorization: Bearer [token]
```

### Body(Form)

| Field Name     | Type   | 
|----------------|--------|
| `display_name` | string |

### Success

- code: `200`

```text
レスポンスBodyなし
```

### Error

- code: `401` | `500`

```json
{
  "error": {
    "message": "ログインしてください"
  }
}
```
