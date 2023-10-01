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

なし

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
