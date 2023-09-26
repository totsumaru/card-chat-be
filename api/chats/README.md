# 自分がhostのチャットを全て取得

自分がhostとなっているチャットを全て取得します。

ログイン必須。

```
[GET] /api/chats
```

### URL Params

なし

### Header

- `Authorization`(required)

```text
Authorization: Bearer [token]
```

### Success

- code: `200`

```json
{
  "chats": [
    {
      "id": "cb273580-8a04-4421-8141-e2bc48a89069"
    },
    {
      "id": "cb273580-8a04-4421-8141-e2bc48a89069"
    }
  ]
}
```

### Error

- code: `401` | `500`

```json
{
  "error": {
    "message": "ログインが必要です"
  }
}
```
