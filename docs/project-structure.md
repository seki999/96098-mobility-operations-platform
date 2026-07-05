# Project Structure

```text
96098-mobility-operations-platform/
  frontend/                 Next.js / TypeScript UI
  backend/                  Go REST API
  backend/migrations/       PostgreSQL schema and seed
  infrastructure/terraform/ AWS reference IaC
  docs/                     Architecture and setup docs
  mock-data/                Neutral portfolio data
  screenshots/              README image outputs
  scripts/                  Screenshot and seed helpers
  .github/workflows/        GitHub Actions CI
```

## 主なファイル

- `README.md`: Portfolio 全体の説明、実行手順、技術対応表。
- `docker-compose.yml`: frontend/backend/postgres を一括起動します。
- `backend/cmd/server/main.go`: Go API の entrypoint。
- `backend/internal/api/server.go`: REST endpoint、CORS、JSON error handling。
- `frontend/src/app/page.tsx`: Dashboard 画面。
- `infrastructure/terraform/main.tf`: AWS 参照構成。

## 関係性

Frontend は Go API を呼び出します。Backend は PostgreSQL を読み書きします。Docker はローカル再現性を担保し、Terraform は AWS での実運用案を表します。
