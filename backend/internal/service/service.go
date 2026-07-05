package service

import (
	"context"

	"96098-mobility-operations-platform/backend/internal/models"
	"96098-mobility-operations-platform/backend/internal/repository"
)

// OperationsService は業務ユースケースを handler から分離します。
type OperationsService struct {
	repo repository.Repository
}

// NewOperationsService は Repository を注入して service を構築します。
func NewOperationsService(repo repository.Repository) *OperationsService {
	return &OperationsService{repo: repo}
}

// ListVehicles は車両一覧を取得します。
func (s *OperationsService) ListVehicles(ctx context.Context, status, q string) ([]models.Vehicle, error) {
	return s.repo.ListVehicles(ctx, repository.NormalizeStatus(status), q)
}

// CreateVehicle は車両登録を実行します。
func (s *OperationsService) CreateVehicle(ctx context.Context, v models.Vehicle) (models.Vehicle, error) {
	return s.repo.CreateVehicle(ctx, v)
}

// GetVehicle は車両詳細を取得します。
func (s *OperationsService) GetVehicle(ctx context.Context, id int64) (models.Vehicle, error) {
	return s.repo.GetVehicle(ctx, id)
}

// UpdateVehicle は車両情報を更新します。
func (s *OperationsService) UpdateVehicle(ctx context.Context, id int64, v models.Vehicle) (models.Vehicle, error) {
	return s.repo.UpdateVehicle(ctx, id, v)
}

// DeleteVehicle は車両を削除します。
func (s *OperationsService) DeleteVehicle(ctx context.Context, id int64) error {
	return s.repo.DeleteVehicle(ctx, id)
}

// ListDrivers はドライバー一覧を取得します。
func (s *OperationsService) ListDrivers(ctx context.Context, status, q string) ([]models.Driver, error) {
	return s.repo.ListDrivers(ctx, repository.NormalizeStatus(status), q)
}

// CreateDriver はドライバー登録を実行します。
func (s *OperationsService) CreateDriver(ctx context.Context, d models.Driver) (models.Driver, error) {
	return s.repo.CreateDriver(ctx, d)
}

// GetDriver はドライバー詳細を取得します。
func (s *OperationsService) GetDriver(ctx context.Context, id int64) (models.Driver, error) {
	return s.repo.GetDriver(ctx, id)
}

// UpdateDriver はドライバー情報を更新します。
func (s *OperationsService) UpdateDriver(ctx context.Context, id int64, d models.Driver) (models.Driver, error) {
	return s.repo.UpdateDriver(ctx, id, d)
}

// DeleteDriver はドライバーを削除します。
func (s *OperationsService) DeleteDriver(ctx context.Context, id int64) error {
	return s.repo.DeleteDriver(ctx, id)
}

// ListTasks は運行タスク一覧を取得します。
func (s *OperationsService) ListTasks(ctx context.Context, status, q string) ([]models.OperationTask, error) {
	return s.repo.ListTasks(ctx, repository.NormalizeStatus(status), q)
}

// CreateTask は運行タスク登録を実行します。
func (s *OperationsService) CreateTask(ctx context.Context, t models.OperationTask) (models.OperationTask, error) {
	return s.repo.CreateTask(ctx, t)
}

// GetTask は運行タスク詳細を取得します。
func (s *OperationsService) GetTask(ctx context.Context, id int64) (models.OperationTask, error) {
	return s.repo.GetTask(ctx, id)
}

// UpdateTask は運行タスクを更新します。
func (s *OperationsService) UpdateTask(ctx context.Context, id int64, t models.OperationTask) (models.OperationTask, error) {
	return s.repo.UpdateTask(ctx, id, t)
}

// ListIncidents は異常イベント一覧を取得します。
func (s *OperationsService) ListIncidents(ctx context.Context, status, q string) ([]models.Incident, error) {
	return s.repo.ListIncidents(ctx, repository.NormalizeStatus(status), q)
}

// CreateIncident は異常イベント登録を実行します。
func (s *OperationsService) CreateIncident(ctx context.Context, i models.Incident) (models.Incident, error) {
	return s.repo.CreateIncident(ctx, i)
}

// GetIncident は異常イベント詳細を取得します。
func (s *OperationsService) GetIncident(ctx context.Context, id int64) (models.Incident, error) {
	return s.repo.GetIncident(ctx, id)
}

// UpdateIncident は異常イベントを更新します。
func (s *OperationsService) UpdateIncident(ctx context.Context, id int64, i models.Incident) (models.Incident, error) {
	return s.repo.UpdateIncident(ctx, id, i)
}

// DeleteIncident は異常イベントを削除します。
func (s *OperationsService) DeleteIncident(ctx context.Context, id int64) error {
	return s.repo.DeleteIncident(ctx, id)
}

// DashboardSummary はダッシュボード KPI を取得します。
func (s *OperationsService) DashboardSummary(ctx context.Context) (models.DashboardSummary, error) {
	return s.repo.DashboardSummary(ctx)
}
