# card-chat-be

チャットカードのバックエンドです

## API

### チャットを取得します

クエリパラメーターで取得方法を分別します。

```
[GET] /api/chat
```

### 自分がホストのチャットを全て取得します

※ログイン必須

```
[GET] /api/chat
```

Query

```
?chat-id=[chatID]
?type=["passcode" | "login" | "cookie"]
```

Header

```
Bearer [token]
```

body

- なし
