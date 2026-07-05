# Technical Design

## 技術選定理由

- Go: 小さな REST API を高速に実装でき、静的型付けと単体テストが扱いやすいです。
- Next.js / TypeScript: 一覧・詳細・Dashboard を型安全に構築できます。
- PostgreSQL: 関係性の強い運行管理データを外部キーで表現できます。
- Terraform: AWS 構成をレビュー可能なコードとして管理できます。
- GitHub Actions: lint、build、test、Terraform check を PR 単位で自動化できます。

## Go backend design

`internal/api`、`internal/service`、`internal/repository`、`internal/models` に分けています。Handler は HTTP と JSON、Service はユースケース、Repository は SQL を担当します。

## Next.js frontend design

App Router を使い、画面ごとに server component を配置します。検索 UI は query string を使う client component として切り出しています。

## PostgreSQL design

主要テーブルは `vehicles`、`drivers`、`operation_tasks`、`service_areas`、`incidents`、`operation_logs` です。`operation_tasks` は車両、担当者、エリアに外部キーを持ちます。

## Error and logging

API は内部エラー内容をレスポンスに出さず、ログへ記録します。CORS は `ALLOWED_ORIGIN` で環境別に制御します。

## Test and CI/CD

Backend は service 層の単体テストを含みます。CI は frontend lint/build、backend test/build、Docker build、Terraform fmt/validate を実行します。
