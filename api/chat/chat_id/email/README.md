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

- `Passcode`(required)

```text
Passcode: [passcode]
```

### Body(Form)

| Field Name | Type   | 
|------------|--------|
| `email`    | string |

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
