# 自分がhostのチャットを全て取得

自分がhostとなっているチャットを全て取得します。

ログイン必須。

```
[GET] /api/chats
```

### URL Params

なし

### Header

- `Authorization`(required)

```text
Authorization: Bearer [token]
```

### Success

- code: `200`

```json
{
  "chats": [
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
      "latest_message": {
        "id": "123e4567-e89b-12d3-a456-426614174001",
        "chat_id": "123e4567-e89b-12d3-a456-426614174002",
        "from_id": "123e4567-e89b-12d3-a456-426614174003",
        "content": "Hello, how are you?",
        "created": "2023-09-29T10:30:00Z"
      }
    }
  ]
}
```

### Error

- code: `401` | `500`

```json
{
  "error": {
    "message": "ログインが必要です"
  }
}
```
