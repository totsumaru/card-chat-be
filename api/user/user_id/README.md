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
  "name": "鈴木 太郎"
}
```

### Error

- code: `404` | `500`

```json
{
  "error": {
    "message": "ホストが存在していません"
  }
}
```
