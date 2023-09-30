# パスコードでチャットを取得

パスコードを使用してチャットを取得します。

パスコードが一致した場合は、cookieに保存します。

ここで取得した場合は、必ず`guest`として送信されます。

※表示名,メモは送信されません。

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

- code: `401` | `500`

```json
{
  "error": {
    "message": "パスコードが一致しません"
  }
}
```