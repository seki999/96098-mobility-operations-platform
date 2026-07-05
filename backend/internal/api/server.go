package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"96098-mobility-operations-platform/backend/internal/models"
	"96098-mobility-operations-platform/backend/internal/repository"
	"96098-mobility-operations-platform/backend/internal/service"
)

// Server は HTTP handler 群と CORS 設定を保持します。
type Server struct {
	service       *service.OperationsService
	allowedOrigin string
}

// NewServer は API サーバーを構築します。
func NewServer(service *service.OperationsService, allowedOrigin string) *Server {
	return &Server{service: service, allowedOrigin: allowedOrigin}
}

// Routes は標準ライブラリの ServeMux で REST API を定義します。
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /api/vehicles", s.listVehicles)
	mux.HandleFunc("POST /api/vehicles", s.createVehicle)
	mux.HandleFunc("/api/vehicles/", s.vehicleByID)
	mux.HandleFunc("GET /api/drivers", s.listDrivers)
	mux.HandleFunc("POST /api/drivers", s.createDriver)
	mux.HandleFunc("/api/drivers/", s.driverByID)
	mux.HandleFunc("GET /api/operation-tasks", s.listTasks)
	mux.HandleFunc("POST /api/operation-tasks", s.createTask)
	mux.HandleFunc("/api/operation-tasks/", s.taskByID)
	mux.HandleFunc("GET /api/incidents", s.listIncidents)
	mux.HandleFunc("POST /api/incidents", s.createIncident)
	mux.HandleFunc("/api/incidents/", s.incidentByID)
	mux.HandleFunc("GET /api/dashboard/summary", s.dashboardSummary)
	return s.withMiddleware(mux)
}

// withMiddleware はログ、CORS、JSON API 共通ヘッダーを付与します。
func (s *Server) withMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		w.Header().Set("Access-Control-Allow-Origin", s.allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// health は稼働確認用エンドポイントです。
func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "service": "mobility-operations-api"})
}

// listVehicles は車両一覧を返します。
func (s *Server) listVehicles(w http.ResponseWriter, r *http.Request) {
	items, err := s.service.ListVehicles(r.Context(), r.URL.Query().Get("status"), r.URL.Query().Get("q"))
	writeResult(w, items, err)
}

// createVehicle は車両を登録します。
func (s *Server) createVehicle(w http.ResponseWriter, r *http.Request) {
	var input models.Vehicle
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := s.service.CreateVehicle(r.Context(), input)
	writeResultWithStatus(w, item, err, http.StatusCreated)
}

// vehicleByID は車両詳細・更新・削除をメソッドで振り分けます。
func (s *Server) vehicleByID(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r.URL.Path, "/api/vehicles/")
	if !ok {
		return
	}
	switch r.Method {
	case http.MethodGet:
		item, err := s.service.GetVehicle(r.Context(), id)
		writeResult(w, item, err)
	case http.MethodPut:
		var input models.Vehicle
		if !decodeJSON(w, r, &input) {
			return
		}
		item, err := s.service.UpdateVehicle(r.Context(), id, input)
		writeResult(w, item, err)
	case http.MethodDelete:
		err := s.service.DeleteVehicle(r.Context(), id)
		writeResult(w, map[string]string{"status": "deleted"}, err)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

// listDrivers はドライバー一覧を返します。
func (s *Server) listDrivers(w http.ResponseWriter, r *http.Request) {
	items, err := s.service.ListDrivers(r.Context(), r.URL.Query().Get("status"), r.URL.Query().Get("q"))
	writeResult(w, items, err)
}

// createDriver はドライバーを登録します。
func (s *Server) createDriver(w http.ResponseWriter, r *http.Request) {
	var input models.Driver
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := s.service.CreateDriver(r.Context(), input)
	writeResultWithStatus(w, item, err, http.StatusCreated)
}

// driverByID はドライバー詳細・更新・削除をメソッドで振り分けます。
func (s *Server) driverByID(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r.URL.Path, "/api/drivers/")
	if !ok {
		return
	}
	switch r.Method {
	case http.MethodGet:
		item, err := s.service.GetDriver(r.Context(), id)
		writeResult(w, item, err)
	case http.MethodPut:
		var input models.Driver
		if !decodeJSON(w, r, &input) {
			return
		}
		item, err := s.service.UpdateDriver(r.Context(), id, input)
		writeResult(w, item, err)
	case http.MethodDelete:
		err := s.service.DeleteDriver(r.Context(), id)
		writeResult(w, map[string]string{"status": "deleted"}, err)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

// listTasks は運行タスク一覧を返します。
func (s *Server) listTasks(w http.ResponseWriter, r *http.Request) {
	items, err := s.service.ListTasks(r.Context(), r.URL.Query().Get("status"), r.URL.Query().Get("q"))
	writeResult(w, items, err)
}

// createTask は運行タスクを登録します。
func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	var input models.OperationTask
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := s.service.CreateTask(r.Context(), input)
	writeResultWithStatus(w, item, err, http.StatusCreated)
}

// taskByID は運行タスク詳細・更新をメソッドで振り分けます。
func (s *Server) taskByID(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r.URL.Path, "/api/operation-tasks/")
	if !ok {
		return
	}
	switch r.Method {
	case http.MethodGet:
		item, err := s.service.GetTask(r.Context(), id)
		writeResult(w, item, err)
	case http.MethodPut:
		var input models.OperationTask
		if !decodeJSON(w, r, &input) {
			return
		}
		item, err := s.service.UpdateTask(r.Context(), id, input)
		writeResult(w, item, err)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

// listIncidents は異常イベント一覧を返します。
func (s *Server) listIncidents(w http.ResponseWriter, r *http.Request) {
	items, err := s.service.ListIncidents(r.Context(), r.URL.Query().Get("status"), r.URL.Query().Get("q"))
	writeResult(w, items, err)
}

// createIncident は異常イベントを登録します。
func (s *Server) createIncident(w http.ResponseWriter, r *http.Request) {
	var input models.Incident
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := s.service.CreateIncident(r.Context(), input)
	writeResultWithStatus(w, item, err, http.StatusCreated)
}

// incidentByID は異常イベント詳細・更新・削除をメソッドで振り分けます。
func (s *Server) incidentByID(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r.URL.Path, "/api/incidents/")
	if !ok {
		return
	}
	switch r.Method {
	case http.MethodGet:
		item, err := s.service.GetIncident(r.Context(), id)
		writeResult(w, item, err)
	case http.MethodPut:
		var input models.Incident
		if !decodeJSON(w, r, &input) {
			return
		}
		item, err := s.service.UpdateIncident(r.Context(), id, input)
		writeResult(w, item, err)
	case http.MethodDelete:
		err := s.service.DeleteIncident(r.Context(), id)
		writeResult(w, map[string]string{"status": "deleted"}, err)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

// dashboardSummary は KPI 集計を返します。
func (s *Server) dashboardSummary(w http.ResponseWriter, r *http.Request) {
	item, err := s.service.DashboardSummary(r.Context())
	writeResult(w, item, err)
}

// decodeJSON は JSON decode 失敗時に安全なエラーを返します。
func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json body"})
		return false
	}
	return true
}

// parseID は URL パスから数値 ID を取り出します。
func parseID(w http.ResponseWriter, path, prefix string) (int64, bool) {
	id, err := strconv.ParseInt(strings.TrimPrefix(path, prefix), 10, 64)
	if err != nil || id <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return 0, false
	}
	return id, true
}

// writeResult は service/repository の結果を HTTP レスポンスへ変換します。
func writeResult(w http.ResponseWriter, payload any, err error) {
	writeResultWithStatus(w, payload, err, http.StatusOK)
}

// writeResultWithStatus は成功時ステータスを指定できるレスポンス補助関数です。
func writeResultWithStatus(w http.ResponseWriter, payload any, err error, status int) {
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": repository.PublicError(err)})
			return
		}
		log.Printf("request failed: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": repository.PublicError(err)})
		return
	}
	writeJSON(w, status, payload)
}

// writeJSON は JSON エンコードを一箇所にまとめます。
func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("json encode failed: %v", err)
	}
}
