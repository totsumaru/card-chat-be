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
