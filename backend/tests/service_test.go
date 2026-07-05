package tests

import (
	"context"
	"testing"

	"96098-mobility-operations-platform/backend/internal/models"
	"96098-mobility-operations-platform/backend/internal/repository"
	"96098-mobility-operations-platform/backend/internal/service"
)

// fakeRepository は service の単体テスト用に最小限の Repository 契約を満たします。
type fakeRepository struct {
	vehicles []models.Vehicle
}

func (f *fakeRepository) ListVehicles(context.Context, string, string) ([]models.Vehicle, error) {
	return f.vehicles, nil
}
func (f *fakeRepository) CreateVehicle(_ context.Context, v models.Vehicle) (models.Vehicle, error) {
	v.ID = 99
	return v, nil
}
func (f *fakeRepository) GetVehicle(context.Context, int64) (models.Vehicle, error) {
	return models.Vehicle{}, repository.ErrNotFound
}
func (f *fakeRepository) UpdateVehicle(context.Context, int64, models.Vehicle) (models.Vehicle, error) {
	return models.Vehicle{}, nil
}
func (f *fakeRepository) DeleteVehicle(context.Context, int64) error { return nil }
func (f *fakeRepository) ListDrivers(context.Context, string, string) ([]models.Driver, error) {
	return nil, nil
}
func (f *fakeRepository) CreateDriver(context.Context, models.Driver) (models.Driver, error) {
	return models.Driver{}, nil
}
func (f *fakeRepository) GetDriver(context.Context, int64) (models.Driver, error) {
	return models.Driver{}, nil
}
func (f *fakeRepository) UpdateDriver(context.Context, int64, models.Driver) (models.Driver, error) {
	return models.Driver{}, nil
}
func (f *fakeRepository) DeleteDriver(context.Context, int64) error { return nil }
func (f *fakeRepository) ListTasks(context.Context, string, string) ([]models.OperationTask, error) {
	return nil, nil
}
func (f *fakeRepository) CreateTask(context.Context, models.OperationTask) (models.OperationTask, error) {
	return models.OperationTask{}, nil
}
func (f *fakeRepository) GetTask(context.Context, int64) (models.OperationTask, error) {
	return models.OperationTask{}, nil
}
func (f *fakeRepository) UpdateTask(context.Context, int64, models.OperationTask) (models.OperationTask, error) {
	return models.OperationTask{}, nil
}
func (f *fakeRepository) ListIncidents(context.Context, string, string) ([]models.Incident, error) {
	return nil, nil
}
func (f *fakeRepository) CreateIncident(context.Context, models.Incident) (models.Incident, error) {
	return models.Incident{}, nil
}
func (f *fakeRepository) GetIncident(context.Context, int64) (models.Incident, error) {
	return models.Incident{}, nil
}
func (f *fakeRepository) UpdateIncident(context.Context, int64, models.Incident) (models.Incident, error) {
	return models.Incident{}, nil
}
func (f *fakeRepository) DeleteIncident(context.Context, int64) error { return nil }
func (f *fakeRepository) DashboardSummary(context.Context) (models.DashboardSummary, error) {
	return models.DashboardSummary{TotalVehicles: 2}, nil
}

// TestListVehicles は service が repository の結果をそのまま返すことを確認します。
func TestListVehicles(t *testing.T) {
	svc := service.NewOperationsService(&fakeRepository{vehicles: []models.Vehicle{{ID: 1, Code: "VH-001"}}})
	items, err := svc.ListVehicles(context.Background(), "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 1 || items[0].Code != "VH-001" {
		t.Fatalf("unexpected vehicles: %#v", items)
	}
}

// TestDashboardSummary は dashboard 集計の受け渡しを確認します。
func TestDashboardSummary(t *testing.T) {
	svc := service.NewOperationsService(&fakeRepository{})
	summary, err := svc.DashboardSummary(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if summary.TotalVehicles != 2 {
		t.Fatalf("unexpected summary: %#v", summary)
	}
}
