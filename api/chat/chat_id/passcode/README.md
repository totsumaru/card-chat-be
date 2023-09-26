# パスコードでチャットを取得

パスコードを使用してチャットを取得します。

パスコードが一致した場合は、cookieに保存します。

ここで取得した場合は、必ず`guest`として送信されます。

```
[GET] /api/chat/[chat-id]/passcode
```

### URL Params

なし

### Header

- `Passcode`(required)

```text
Passcode: [passcode]
```

### Success

- code: `200`

```json
{
  "id": "cb273580-8a04-4421-8141-e2bc48a89069",
  "status": "guest"
}
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