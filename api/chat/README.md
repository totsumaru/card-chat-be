# チャットの取得

誰でもチャットページは開けるようにするため、極力エラーは返さずに`status`で返します。

クリティカルな場合(悪意のあるリクエストなど)のみエラーを返します。

```
[GET] /api/chat/[chat-id]
```

### URL Params

なし

### Header

- `Authorization`(ログインしている場合のみ)

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
- ゲストの場合のみパスコードを返します
    - ※通知用Emailの設定のために返します

```json
{
  "id": "cb273580-8a04-4421-8141-e2bc48a89069",
  "status": "guest",
  "passcode": "123456"
}
```

### Error

- code: `404` | `500`

```json
{
  "error": {
    "message": "チャットが存在していません"
  }
}
```
