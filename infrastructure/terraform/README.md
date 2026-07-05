# Terraform 設計

このディレクトリは AWS 上の参照構成を表す IaC テンプレートです。実行には AWS 認証情報と `db_password` の安全な指定が必要です。

## 含まれる設計

- VPC と public subnet
- App 用 Security Group
- RDS PostgreSQL 用 Security Group
- CloudWatch Logs
- ECR
- RDS PostgreSQL
- App Runner による Go API コンテナ公開案

## ローカルチェック

```bash
terraform fmt -check -recursive
terraform init -backend=false
terraform validate
```

本番では `terraform.tfvars` や CI secret を使い、秘密情報を GitHub にコミットしません。
