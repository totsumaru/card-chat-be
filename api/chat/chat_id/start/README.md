# チャットの開始

ホストを設定します。

- ログインしているユーザーのみが設定できます

```
[POST] /api/chat/[chat-id]/start
```

### URL Params

なし

### Header

- `Authorization`(required)

```text
Authorization: Bearer [token]
```

- `Content-Type`

```text
Content-Type: application/x-www-form-urlencoded
```

### Body(Form)

| Field Name     | Type   | 
|----------------|--------|
| `display_name` | string |

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
  }
}
```

### Error

- code: `401` | `500`

```json
{
  "error": {
    "message": "ログインしてください"
  }
}
```
