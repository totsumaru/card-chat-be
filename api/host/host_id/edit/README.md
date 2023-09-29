# ホスト情報の変更

ホストのプロフィール情報を変更します。

- ログインしているホストのみ

```
[POST] /api/host/[host-id]/edit
```

### URL Params

なし

### Header

- `Content-Type`

```text
Content-Type: multipart/form-data[README.md](..%2F..%2Fcreate%2FREADME.md)
```

- `Authorization`(required)

```text
Authorization: Bearer [token]
```

### Body(Form)

| Field Name     | Type   | 
|----------------|--------|
| `avatar`       | File   |
| `name`         | String |
| `headline`     | String |
| `introduction` | String |
| `company_name` | String |
| `position`     | String |
| `tel`          | String |
| `email`        | String |
| `website`      | String |

### Success

- code: `200`

```text
レスポンスBodyなし
```

### Error

- code: `400` | `401` | `500`

```json
{
  "error": {
    "message": "認証できません"
  }
}
```
