# メールアドレスの登録/変更

通知用のメールアドレスを変更します。

Headerのパスコードで認証します。

※パスコードはFEでstateで管理します。チャットを取得する時に、毎回パスコードも返します。

```
[POST] /api/chat/[chat-id]/email
```

### URL Params

なし

### Header

- `Content-Type`

```text
Content-Type: application/x-www-form-urlencoded
```

### Body(Form)

| Field Name | Type   | 
|------------|--------|
| `email`    | string |

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
    "message": "認証できません"
  }
}
```
