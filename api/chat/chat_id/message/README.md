# メッセージの送信

メッセージを送信します。

自分がホストであれば`host`、cookieのパスコードで認証した場合は`guest`として登録します。

kindが`text`の場合はtextのみ、`image`の場合は画像ファイルがが必要です。

```
[POST] /api/chat/[chat-id]/message
```

### URL Params

- `kind`(required)
    - `text` | `image`

```
?kind=text
```

### Header

`Authorization`(hostで送信する場合は必要)

- `Authorization`

```text
Authorization: Bearer [token]
```

- `Content-Type`

```text
Content-Type: multipart/form-data
```

### Body(Form)

| Field Name | Type   | Memo         | 
|------------|--------|--------------|
| `image`    | File   | kind=`image` |
| `text`     | String | kind=`text`  |

### Success

- code: `200`

```text
レスポンスBodyなし
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
