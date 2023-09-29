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
  "id": "cb273580-8a04-4421-8141-e2bc48a89069",
  "name": "鈴木 太郎",
  "avatar_url": "https://example.com/avatar.jpg",
  "headline": "お客様の笑顔のために働いています。",
  "introduction": "初めまして。私は鈴木太郎です。",
  "company": {
    "name": "株式会社ArGate",
    "position": "営業部 営業一課",
    "tel": "123-456-7890",
    "email": "john.doe@example.com",
    "website": "https://techcorp.example.com"
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
