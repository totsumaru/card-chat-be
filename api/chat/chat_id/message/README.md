# メッセージの送信

メッセージを送信します。

自分がホストであれば`host`、cookieのパスコードで認証した場合は`guest`として登録します。

```
[POST] /api/chat/[chat-id]/message
```

### URL Params

なし

### Header

`Authorization`(hostで送信する場合は必要)

- `Authorization`

```text
Authorization: Bearer [token]
```

- `Content-Type`

```text
Content-Type: application/x-www-form-urlencoded
```

### Body(Form)

| Field Name | Type   | 
|------------|--------|
| `content`  | String |

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
