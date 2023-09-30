# チャットの取得

誰でもチャットページは開けるようにするため、極力エラーは返さずに`status`で返します。

クリティカルな場合(悪意のあるリクエストなど)のみエラーを返します。

認証できたら、Websocketでリアルタイム通信を実装します。

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

- `Passcode`(ゲストでチャット取得済みの場合のみ)

```text
Passcode: [passcode]
```

### Success

- code: `200`
- レスポンスのステータス
    - チャットが開始していない:
        - ログイン済み: `first-is-login`
        - ログインしていない: `first-not-login`
    - チャットが開始している:
        - ホスト: `host`
        - ホストでは無い:
            - cookieまたはheaderのパスコードが一致: `guest`
            - cookieまたはheaderのパスコードが一致しない: `visitor`
- ゲストの場合のみパスコードを返します
    - ※通知用Emailの設定のために返します

```json
{
  "status": "host",
  "chat": {
    "id": "a123b456-7890-c123-d456-7890e123f456",
    "passcode": "123456",
    "host_id": "b123c456-7890-d123-e456-7890f123g456",
    "guest": {
      "display_name": "John Doe",
      "memo": "Some notes about the guest.",
      "email": "john.doe@example.com"
    },
    "is_read": false,
    "is_closed": false,
    "last_message": "2023-09-29T10:30:00Z"
  },
  "messages": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174001",
      "chat_id": "123e4567-e89b-12d3-a456-426614174002",
      "from_id": "123e4567-e89b-12d3-a456-426614174003",
      "content": "Hello, how are you?",
      "created": "2023-09-29T10:30:00Z"
    }
  ]
}
```

### Error

- code: `500`

```json
{
  "error": {
    "message": "チャットが存在していません"
  }
}
```
