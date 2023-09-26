# メールアドレスの登録/変更

通知用のメールアドレスを変更します。

Headerのパスコードで認証します。

```
[POST] /api/chat/[chat-id]/email
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
  "email": "test@gmail.com"
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
    "message": "パスコードが一致しません"
  }
}
```
