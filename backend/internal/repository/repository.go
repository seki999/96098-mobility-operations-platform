package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"96098-mobility-operations-platform/backend/internal/models"
)

// Repository は handler/service から見た永続化処理の契約です。
type Repository interface {
	ListVehicles(context.Context, string, string) ([]models.Vehicle, error)
	CreateVehicle(context.Context, models.Vehicle) (models.Vehicle, error)
	GetVehicle(context.Context, int64) (models.Vehicle, error)
	UpdateVehicle(context.Context, int64, models.Vehicle) (models.Vehicle, error)
	DeleteVehicle(context.Context, int64) error
	ListDrivers(context.Context, string, string) ([]models.Driver, error)
	CreateDriver(context.Context, models.Driver) (models.Driver, error)
	GetDriver(context.Context, int64) (models.Driver, error)
	UpdateDriver(context.Context, int64, models.Driver) (models.Driver, error)
	DeleteDriver(context.Context, int64) error
	ListTasks(context.Context, string, string) ([]models.OperationTask, error)
	CreateTask(context.Context, models.OperationTask) (models.OperationTask, error)
	GetTask(context.Context, int64) (models.OperationTask, error)
	UpdateTask(context.Context, int64, models.OperationTask) (models.OperationTask, error)
	ListIncidents(context.Context, string, string) ([]models.Incident, error)
	CreateIncident(context.Context, models.Incident) (models.Incident, error)
	GetIncident(context.Context, int64) (models.Incident, error)
	UpdateIncident(context.Context, int64, models.Incident) (models.Incident, error)
	DeleteIncident(context.Context, int64) error
	DashboardSummary(context.Context) (models.DashboardSummary, error)
}

// ErrNotFound は指定 ID のデータが存在しない場合に返します。
var ErrNotFound = errors.New("resource not found")

// PostgresRepository は PostgreSQL を利用した Repository 実装です。
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository は DB 接続を受け取り repository を構築します。
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// ListVehicles は検索文字列とステータスで車両一覧を絞り込みます。
func (r *PostgresRepository) ListVehicles(ctx context.Context, status, q string) ([]models.Vehicle, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, code, type, status, service_area_id, utilization_pct, last_checked_at
		FROM vehicles
		WHERE ($1 = '' OR status = $1)
		  AND ($2 = '' OR lower(code) LIKE lower('%' || $2 || '%') OR lower(type) LIKE lower('%' || $2 || '%'))
		ORDER BY id`, status, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Vehicle
	for rows.Next() {
		var v models.Vehicle
		if err := rows.Scan(&v.ID, &v.Code, &v.Type, &v.Status, &v.ServiceAreaID, &v.UtilizationPct, &v.LastCheckedAt); err != nil {
			return nil, err
		}
		items = append(items, v)
	}
	return items, rows.Err()
}

// CreateVehicle は新しい車両を登録します。
func (r *PostgresRepository) CreateVehicle(ctx context.Context, v models.Vehicle) (models.Vehicle, error) {
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO vehicles (code, type, status, service_area_id, utilization_pct, last_checked_at)
		VALUES ($1,$2,$3,$4,$5,COALESCE($6, now()))
		RETURNING id, last_checked_at`, v.Code, v.Type, v.Status, v.ServiceAreaID, v.UtilizationPct, nullableTime(v.LastCheckedAt)).
		Scan(&v.ID, &v.LastCheckedAt)
	return v, err
}

// GetVehicle は ID で車両を取得します。
func (r *PostgresRepository) GetVehicle(ctx context.Context, id int64) (models.Vehicle, error) {
	var v models.Vehicle
	err := r.db.QueryRowContext(ctx, `
		SELECT id, code, type, status, service_area_id, utilization_pct, last_checked_at
		FROM vehicles WHERE id=$1`, id).
		Scan(&v.ID, &v.Code, &v.Type, &v.Status, &v.ServiceAreaID, &v.UtilizationPct, &v.LastCheckedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return v, ErrNotFound
	}
	return v, err
}

// UpdateVehicle は既存車両を更新します。
func (r *PostgresRepository) UpdateVehicle(ctx context.Context, id int64, v models.Vehicle) (models.Vehicle, error) {
	err := r.db.QueryRowContext(ctx, `
		UPDATE vehicles
		SET code=$2, type=$3, status=$4, service_area_id=$5, utilization_pct=$6, last_checked_at=now()
		WHERE id=$1
		RETURNING id, last_checked_at`, id, v.Code, v.Type, v.Status, v.ServiceAreaID, v.UtilizationPct).
		Scan(&v.ID, &v.LastCheckedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return v, ErrNotFound
	}
	return v, err
}

// DeleteVehicle は車両を削除します。
func (r *PostgresRepository) DeleteVehicle(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM vehicles WHERE id=$1`, id)
	if err != nil {
		return err
	}
	count, _ := result.RowsAffected()
	if count == 0 {
		return ErrNotFound
	}
	return nil
}

// ListDrivers は検索文字列とステータスでドライバー一覧を返します。
func (r *PostgresRepository) ListDrivers(ctx context.Context, status, q string) ([]models.Driver, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, code, display_name, status, skill_level, service_area_id, updated_at
		FROM drivers
		WHERE ($1 = '' OR status = $1)
		  AND ($2 = '' OR lower(code) LIKE lower('%' || $2 || '%') OR lower(display_name) LIKE lower('%' || $2 || '%'))
		ORDER BY id`, status, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Driver
	for rows.Next() {
		var d models.Driver
		if err := rows.Scan(&d.ID, &d.Code, &d.DisplayName, &d.Status, &d.SkillLevel, &d.ServiceAreaID, &d.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, d)
	}
	return items, rows.Err()
}

// CreateDriver はドライバーを登録します。
func (r *PostgresRepository) CreateDriver(ctx context.Context, d models.Driver) (models.Driver, error) {
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO drivers (code, display_name, status, skill_level, service_area_id)
		VALUES ($1,$2,$3,$4,$5)
		RETURNING id, updated_at`, d.Code, d.DisplayName, d.Status, d.SkillLevel, d.ServiceAreaID).
		Scan(&d.ID, &d.UpdatedAt)
	return d, err
}

// GetDriver は ID でドライバーを取得します。
func (r *PostgresRepository) GetDriver(ctx context.Context, id int64) (models.Driver, error) {
	var d models.Driver
	err := r.db.QueryRowContext(ctx, `
		SELECT id, code, display_name, status, skill_level, service_area_id, updated_at
		FROM drivers WHERE id=$1`, id).
		Scan(&d.ID, &d.Code, &d.DisplayName, &d.Status, &d.SkillLevel, &d.ServiceAreaID, &d.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return d, ErrNotFound
	}
	return d, err
}

// UpdateDriver はドライバー情報を更新します。
func (r *PostgresRepository) UpdateDriver(ctx context.Context, id int64, d models.Driver) (models.Driver, error) {
	err := r.db.QueryRowContext(ctx, `
		UPDATE drivers
		SET code=$2, display_name=$3, status=$4, skill_level=$5, service_area_id=$6, updated_at=now()
		WHERE id=$1
		RETURNING id, updated_at`, id, d.Code, d.DisplayName, d.Status, d.SkillLevel, d.ServiceAreaID).
		Scan(&d.ID, &d.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return d, ErrNotFound
	}
	return d, err
}

// DeleteDriver はドライバーを削除します。
func (r *PostgresRepository) DeleteDriver(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM drivers WHERE id=$1`, id)
	if err != nil {
		return err
	}
	count, _ := result.RowsAffected()
	if count == 0 {
		return ErrNotFound
	}
	return nil
}

// ListTasks は運行タスクを検索・絞り込みします。
func (r *PostgresRepository) ListTasks(ctx context.Context, status, q string) ([]models.OperationTask, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, task_code, vehicle_id, driver_id, service_area_id, status, priority, scheduled_at, notes
		FROM operation_tasks
		WHERE ($1 = '' OR status = $1)
		  AND ($2 = '' OR lower(task_code) LIKE lower('%' || $2 || '%') OR lower(notes) LIKE lower('%' || $2 || '%'))
		ORDER BY scheduled_at DESC`, status, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.OperationTask
	for rows.Next() {
		var t models.OperationTask
		if err := rows.Scan(&t.ID, &t.TaskCode, &t.VehicleID, &t.DriverID, &t.ServiceAreaID, &t.Status, &t.Priority, &t.ScheduledAt, &t.Notes); err != nil {
			return nil, err
		}
		items = append(items, t)
	}
	return items, rows.Err()
}

// CreateTask は運行タスクを登録します。
func (r *PostgresRepository) CreateTask(ctx context.Context, t models.OperationTask) (models.OperationTask, error) {
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO operation_tasks (task_code, vehicle_id, driver_id, service_area_id, status, priority, scheduled_at, notes)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		RETURNING id`, t.TaskCode, t.VehicleID, t.DriverID, t.ServiceAreaID, t.Status, t.Priority, t.ScheduledAt, t.Notes).Scan(&t.ID)
	return t, err
}

// GetTask は ID で運行タスクを取得します。
func (r *PostgresRepository) GetTask(ctx context.Context, id int64) (models.OperationTask, error) {
	var t models.OperationTask
	err := r.db.QueryRowContext(ctx, `
		SELECT id, task_code, vehicle_id, driver_id, service_area_id, status, priority, scheduled_at, notes
		FROM operation_tasks WHERE id=$1`, id).
		Scan(&t.ID, &t.TaskCode, &t.VehicleID, &t.DriverID, &t.ServiceAreaID, &t.Status, &t.Priority, &t.ScheduledAt, &t.Notes)
	if errors.Is(err, sql.ErrNoRows) {
		return t, ErrNotFound
	}
	return t, err
}

// UpdateTask は運行タスクを更新します。
func (r *PostgresRepository) UpdateTask(ctx context.Context, id int64, t models.OperationTask) (models.OperationTask, error) {
	err := r.db.QueryRowContext(ctx, `
		UPDATE operation_tasks
		SET task_code=$2, vehicle_id=$3, driver_id=$4, service_area_id=$5, status=$6, priority=$7, scheduled_at=$8, notes=$9
		WHERE id=$1
		RETURNING id`, id, t.TaskCode, t.VehicleID, t.DriverID, t.ServiceAreaID, t.Status, t.Priority, t.ScheduledAt, t.Notes).Scan(&t.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return t, ErrNotFound
	}
	return t, err
}

// ListIncidents は異常イベントを検索・絞り込みします。
func (r *PostgresRepository) ListIncidents(ctx context.Context, status, q string) ([]models.Incident, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, incident_code, operation_task_id, severity, status, summary, detected_at
		FROM incidents
		WHERE ($1 = '' OR status = $1)
		  AND ($2 = '' OR lower(incident_code) LIKE lower('%' || $2 || '%') OR lower(summary) LIKE lower('%' || $2 || '%'))
		ORDER BY detected_at DESC`, status, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Incident
	for rows.Next() {
		var i models.Incident
		if err := rows.Scan(&i.ID, &i.IncidentCode, &i.OperationTaskID, &i.Severity, &i.Status, &i.Summary, &i.DetectedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, rows.Err()
}

// CreateIncident は異常イベントを登録します。
func (r *PostgresRepository) CreateIncident(ctx context.Context, i models.Incident) (models.Incident, error) {
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO incidents (incident_code, operation_task_id, severity, status, summary, detected_at)
		VALUES ($1,$2,$3,$4,$5,COALESCE($6, now()))
		RETURNING id, detected_at`, i.IncidentCode, i.OperationTaskID, i.Severity, i.Status, i.Summary, nullableTime(i.DetectedAt)).
		Scan(&i.ID, &i.DetectedAt)
	return i, err
}

// GetIncident は ID で異常イベントを取得します。
func (r *PostgresRepository) GetIncident(ctx context.Context, id int64) (models.Incident, error) {
	var i models.Incident
	err := r.db.QueryRowContext(ctx, `
		SELECT id, incident_code, operation_task_id, severity, status, summary, detected_at
		FROM incidents WHERE id=$1`, id).
		Scan(&i.ID, &i.IncidentCode, &i.OperationTaskID, &i.Severity, &i.Status, &i.Summary, &i.DetectedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return i, ErrNotFound
	}
	return i, err
}

// UpdateIncident は異常イベントの状態や要約を更新します。
func (r *PostgresRepository) UpdateIncident(ctx context.Context, id int64, i models.Incident) (models.Incident, error) {
	err := r.db.QueryRowContext(ctx, `
		UPDATE incidents
		SET incident_code=$2, operation_task_id=$3, severity=$4, status=$5, summary=$6, detected_at=COALESCE($7, detected_at)
		WHERE id=$1
		RETURNING id, detected_at`, id, i.IncidentCode, i.OperationTaskID, i.Severity, i.Status, i.Summary, nullableTime(i.DetectedAt)).
		Scan(&i.ID, &i.DetectedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return i, ErrNotFound
	}
	return i, err
}

// DeleteIncident は異常イベントを削除します。
func (r *PostgresRepository) DeleteIncident(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM incidents WHERE id=$1`, id)
	if err != nil {
		return err
	}
	count, _ := result.RowsAffected()
	if count == 0 {
		return ErrNotFound
	}
	return nil
}

// DashboardSummary は主要 KPI を SQL で集計します。
func (r *PostgresRepository) DashboardSummary(ctx context.Context) (models.DashboardSummary, error) {
	var s models.DashboardSummary
	err := r.db.QueryRowContext(ctx, `
		SELECT
		  (SELECT count(*) FROM vehicles),
		  (SELECT count(*) FROM vehicles WHERE status='active'),
		  (SELECT count(*) FROM drivers WHERE status='available'),
		  (SELECT count(*) FROM operation_tasks WHERE status IN ('planned','in_progress')),
		  (SELECT count(*) FROM incidents WHERE status <> 'resolved'),
		  COALESCE((SELECT avg(utilization_pct) FROM vehicles), 0)`).Scan(
		&s.TotalVehicles, &s.ActiveVehicles, &s.AvailableDrivers, &s.OpenTasks, &s.OpenIncidents, &s.AvgUtilizationPct,
	)
	return s, err
}

// nullableTime はゼロ値の時刻を SQL NULL として扱うための補助関数です。
func nullableTime(t time.Time) any {
	if t.IsZero() {
		return nil
	}
	return t
}

// NormalizeStatus は API 入力のステータスを簡易的に正規化します。
func NormalizeStatus(v string) string {
	return strings.TrimSpace(strings.ToLower(v))
}

// PublicError は内部エラー詳細を外部へ出さないためのメッセージを返します。
func PublicError(err error) string {
	if errors.Is(err, ErrNotFound) {
		return "resource not found"
	}
	if err != nil {
		return fmt.Sprintf("request failed: %s", "internal error")
	}
	return ""
}
