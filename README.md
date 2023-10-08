# card-chat-be

チャットカードのバックエンドです

## API

一覧は[こちら](./api)

## インフラ構成

### Frontend

- Vercel
    - Prd
    - Stg

### Backend

- Render
    - Prd
    - Stg

### DB

- Render
    - Prd
    - Stg

### Auth

- Supabase
    - Prd/Stg/Dev共有

Supabaseのプロジェクトはdevからprdまで共有で作成しています。

ある程度の規模までユーザーが拡大できた時にstg用の環境を作成します。

現在は管理コストを低くするため、リスクを許容して共有しています。

#### なぜ共有なのか？

- Resendと統合できるアカウントが1つだけのため
- stg用のプロジェクトを作成すると、管理コストが高くなるため

### Email

- Resend
    - Prd/Stg/Dev共有