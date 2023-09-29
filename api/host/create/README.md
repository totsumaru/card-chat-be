# ホストの作成

ホストのアカウントを新規作成します。

- ログインしているユーザーのみ、自身のアカウントを作成できます

```
[POST] /api/host/create
```

### URL Params

なし

### Header

- `Authorization`(required)

```
Authorization: Bearer [token]
```

### Body

```
リクエストBodyなし
```

### Success

- code: `200`

```
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
