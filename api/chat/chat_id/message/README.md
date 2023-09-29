# メッセージの送信

メッセージを送信します。

自分がホストであれば`host`、パスコードで認証した場合は`guest`として登録します。

```
[POST] /api/chat/[chat-id]/message
```

### URL Params

なし

### Header

`Authorization`,`Passcode`のどちらか一方は必ず必要です。

- `Authorization`

```text
Authorization: Bearer [token]
```

- `Passcode`

```text
Passcode: [passcode]
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

- code: `401` | `404` | `500`

```json
{
  "error": {
    "message": "認証できません"
  }
}
```
