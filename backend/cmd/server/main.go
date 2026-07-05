package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"96098-mobility-operations-platform/backend/internal/api"
	"96098-mobility-operations-platform/backend/internal/config"
	"96098-mobility-operations-platform/backend/internal/repository"
	"96098-mobility-operations-platform/backend/internal/service"
)

// main は API サーバーの起動点です。環境変数から設定を読み込み、DB 接続とルーティングを初期化します。
func main() {
	cfg := config.Load()
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database open failed: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Printf("database ping failed, API still starts for health check: %v", err)
	}

	repo := repository.NewPostgresRepository(db)
	svc := service.NewOperationsService(repo)
	server := api.NewServer(svc, cfg.AllowedOrigin)

	addr := ":" + cfg.Port
	log.Printf("mobility operations API listening on %s", addr)
	if err := http.ListenAndServe(addr, server.Routes()); err != nil {
		log.Printf("server stopped: %v", err)
		os.Exit(1)
	}
}
