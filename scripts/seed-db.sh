#!/usr/bin/env sh
set -eu

# Docker Compose の postgres コンテナに migration と seed を再投入します。
psql "${DATABASE_URL:-postgres://mobility:mobility@localhost:5432/mobility_ops?sslmode=disable}" -f backend/migrations/001_init.sql
psql "${DATABASE_URL:-postgres://mobility:mobility@localhost:5432/mobility_ops?sslmode=disable}" -f backend/migrations/002_seed.sql
