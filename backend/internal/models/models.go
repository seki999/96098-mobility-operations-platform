package models

import "time"

// Vehicle は運用対象車両を表します。実データではなく中立的なサンプル属性のみを扱います。
type Vehicle struct {
	ID             int64     `json:"id"`
	Code           string    `json:"code"`
	Type           string    `json:"type"`
	Status         string    `json:"status"`
	ServiceAreaID  int64     `json:"serviceAreaId"`
	UtilizationPct float64   `json:"utilizationPct"`
	LastCheckedAt  time.Time `json:"lastCheckedAt"`
}

// Driver は運転担当者を表します。個人情報にならないよう仮名コードを利用します。
type Driver struct {
	ID            int64     `json:"id"`
	Code          string    `json:"code"`
	DisplayName   string    `json:"displayName"`
	Status        string    `json:"status"`
	SkillLevel    string    `json:"skillLevel"`
	ServiceAreaID int64     `json:"serviceAreaId"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// OperationTask は配車や点検などの運転関連タスクを表します。
type OperationTask struct {
	ID            int64     `json:"id"`
	TaskCode      string    `json:"taskCode"`
	VehicleID     int64     `json:"vehicleId"`
	DriverID      int64     `json:"driverId"`
	ServiceAreaID int64     `json:"serviceAreaId"`
	Status        string    `json:"status"`
	Priority      string    `json:"priority"`
	ScheduledAt   time.Time `json:"scheduledAt"`
	Notes         string    `json:"notes"`
}

// Incident は運行中に発生した異常イベントを表します。
type Incident struct {
	ID              int64     `json:"id"`
	IncidentCode    string    `json:"incidentCode"`
	OperationTaskID int64     `json:"operationTaskId"`
	Severity        string    `json:"severity"`
	Status          string    `json:"status"`
	Summary         string    `json:"summary"`
	DetectedAt      time.Time `json:"detectedAt"`
}

// DashboardSummary は運用ダッシュボードに表示する集計値を表します。
type DashboardSummary struct {
	TotalVehicles     int64   `json:"totalVehicles"`
	ActiveVehicles    int64   `json:"activeVehicles"`
	AvailableDrivers  int64   `json:"availableDrivers"`
	OpenTasks         int64   `json:"openTasks"`
	OpenIncidents     int64   `json:"openIncidents"`
	AvgUtilizationPct float64 `json:"avgUtilizationPct"`
}
