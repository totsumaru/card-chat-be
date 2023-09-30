# チャットの作成

空のチャットを作成します。

- ADMIN(運営)しか実行できません

```
[POST] /api/chat/create
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
  "chat_id": "cb273580-8a04-4421-8141-e2bc48a89069",
  "passcode": "123456"
}
```

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
    "message": "認証できません"
  }
}
```
