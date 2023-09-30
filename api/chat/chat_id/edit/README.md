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

### Body(Form)

| Field Name     | Type   | 
|----------------|--------|
| `display_name` | string |
| `memo`         | String |

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
