# ホストを取得

ホストの情報を取得します。

プロフィールで表示する情報のため、ログインは不要です。

```
[GET] /api/host/[host-id]
```

### URL Params

なし

### Header

なし

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

- code: `500`

```json
{
  "error": {
    "message": "エラーが発生しました"
  }
}
```
