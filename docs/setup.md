# Setup

## 前提

- Go 1.22+
- Node.js 20+
- Docker Desktop
- Terraform 1.7+

## ローカル起動

```bash
cd 96098-mobility-operations-platform
docker compose up -d postgres
cd backend
go mod tidy
go run ./cmd/server
```

別ターミナル:

```bash
cd frontend
npm install
npm run dev
```

## Docker 起動

```bash
docker compose up --build
```

- Frontend: http://localhost:3000
- Backend health: http://localhost:8080/health
- PostgreSQL: localhost:5432

## PostgreSQL 初期化

Docker Compose 初回起動時に `backend/migrations/001_init.sql` と `002_seed.sql` が実行されます。再投入する場合は `scripts/seed-db.sh` を使います。

## Terraform check

```bash
cd infrastructure/terraform
terraform fmt -check -recursive
terraform init -backend=false
terraform validate
```

## GitHub Actions

`.github/workflows/ci.yml` が PR / push で frontend、backend、Docker、Terraform を検査します。

## FAQ

- API が落ちる場合: `DATABASE_URL` と PostgreSQL 起動状態を確認します。
- CORS エラー: `ALLOWED_ORIGIN` が frontend URL と一致しているか確認します。
- screenshot が生成できない場合: frontend 起動後に Playwright Chromium を install します。
