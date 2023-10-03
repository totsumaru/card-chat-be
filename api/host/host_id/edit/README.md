# ホスト情報の変更

ホストのプロフィール情報を変更します。

- ログインしているホストのみ

```
[POST] /api/host/[host-id]/edit
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
Content-Type: multipart/form-data
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

```json
{
  "host": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "John Doe",
    "avatar_url": "https://example.com/avatar.jpg",
    "headline": "Experienced Software Developer",
    "introduction": "With over 10 years of experience in the tech industry, I have honed my skills in full-stack development, cloud computing, and machine learning.",
    "company": {
      "name": "TechCorp",
      "position": "Senior Software Engineer",
      "tel": "123-456-7890",
      "email": "john.doe@techcorp.com",
      "website": "https://techcorp.com"
    }
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
