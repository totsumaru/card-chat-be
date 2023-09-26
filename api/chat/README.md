# チャットの取得

誰でもチャットページは開けるようにするため、極力エラーは返さずに`status`で返します。

クリティカルな場合(悪意のあるリクエストなど)のみエラーを返します。

```
[GET] /api/chat
```

### URL Params

- `chat-id`(required): uuid
    - e.g. cb273580-8a04-4421-8141-e2bc48a89069

### Header

```text
Authorization: Bearer [token]
```

### Success

- code: `200`
- レスポンスのステータス
    - チャットが開始している:
        - ホスト: `host`
        - ホストでは無い:
            - cookieのパスコードが一致: `guest`
            - cookieのパスコードが一致しない: `visitor`
    - チャットが開始していない:
        - ログイン済み: `first-is-login`
        - ログインしていない: `first-not-login`

```json
{
  "id": "cb273580-8a04-4421-8141-e2bc48a89069",
  "status": "host"
}
```

### ERROR

- code: `404` | `500`