package config

import "os"

// Config はアプリケーション起動時に必要な設定値をまとめます。
type Config struct {
	Port          string
	DatabaseURL   string
	AllowedOrigin string
}

// Load は環境変数から設定を読み込みます。秘密情報をコードに埋め込まないために env を利用します。
func Load() Config {
	return Config{
		Port:          valueOrDefault("PORT", "8080"),
		DatabaseURL:   valueOrDefault("DATABASE_URL", "postgres://mobility:mobility@localhost:5432/mobility_ops?sslmode=disable"),
		AllowedOrigin: valueOrDefault("ALLOWED_ORIGIN", "http://localhost:3000"),
	}
}

// valueOrDefault は未設定の環境変数に安全なローカル既定値を与えます。
func valueOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
