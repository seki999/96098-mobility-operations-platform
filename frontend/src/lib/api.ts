export type Vehicle = {
  id: number;
  code: string;
  type: string;
  status: string;
  serviceAreaId: number;
  utilizationPct: number;
  lastCheckedAt: string;
};

export type Driver = {
  id: number;
  code: string;
  displayName: string;
  status: string;
  skillLevel: string;
  serviceAreaId: number;
  updatedAt: string;
};

export type OperationTask = {
  id: number;
  taskCode: string;
  vehicleId: number;
  driverId: number;
  serviceAreaId: number;
  status: string;
  priority: string;
  scheduledAt: string;
  notes: string;
};

export type Incident = {
  id: number;
  incidentCode: string;
  operationTaskId: number;
  severity: string;
  status: string;
  summary: string;
  detectedAt: string;
};

export type DashboardSummary = {
  totalVehicles: number;
  activeVehicles: number;
  availableDrivers: number;
  openTasks: number;
  openIncidents: number;
  avgUtilizationPct: number;
};

const API_BASE = process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080";

// fetchJson は API 呼び出しを一箇所に集約します。失敗時は画面で mock data に戻せるよう例外を投げます。
async function fetchJson<T>(path: string): Promise<T> {
  const response = await fetch(`${API_BASE}${path}`, { cache: "no-store" });
  if (!response.ok) {
    throw new Error(`API request failed: ${response.status}`);
  }
  return response.json() as Promise<T>;
}

export const fallbackSummary: DashboardSummary = {
  totalVehicles: 4,
  activeVehicles: 2,
  availableDrivers: 2,
  openTasks: 2,
  openIncidents: 2,
  avgUtilizationPct: 61.8
};

export const fallbackVehicles: Vehicle[] = [
  { id: 1, code: "VH-96098-001", type: "compact-ev", status: "active", serviceAreaId: 1, utilizationPct: 78.5, lastCheckedAt: new Date().toISOString() },
  { id: 2, code: "VH-96098-002", type: "accessible-van", status: "maintenance", serviceAreaId: 2, utilizationPct: 35, lastCheckedAt: new Date().toISOString() },
  { id: 3, code: "VH-96098-003", type: "standard-ev", status: "active", serviceAreaId: 2, utilizationPct: 91.2, lastCheckedAt: new Date().toISOString() }
];

export const fallbackDrivers: Driver[] = [
  { id: 1, code: "DR-96098-A01", displayName: "Operator Alpha", status: "available", skillLevel: "senior", serviceAreaId: 1, updatedAt: new Date().toISOString() },
  { id: 2, code: "DR-96098-B02", displayName: "Operator Bravo", status: "assigned", skillLevel: "standard", serviceAreaId: 2, updatedAt: new Date().toISOString() },
  { id: 3, code: "DR-96098-C03", displayName: "Operator Charlie", status: "available", skillLevel: "standard", serviceAreaId: 3, updatedAt: new Date().toISOString() }
];

export const fallbackTasks: OperationTask[] = [
  { id: 1, taskCode: "TASK-96098-1001", vehicleId: 1, driverId: 1, serviceAreaId: 1, status: "planned", priority: "high", scheduledAt: new Date().toISOString(), notes: "Hub transfer simulation" },
  { id: 2, taskCode: "TASK-96098-1002", vehicleId: 3, driverId: 2, serviceAreaId: 2, status: "in_progress", priority: "medium", scheduledAt: new Date().toISOString(), notes: "Central zone recurring operation" }
];

export const fallbackIncidents: Incident[] = [
  { id: 1, incidentCode: "INC-96098-501", operationTaskId: 2, severity: "medium", status: "investigating", summary: "Delayed status update from terminal", detectedAt: new Date().toISOString() },
  { id: 2, incidentCode: "INC-96098-502", operationTaskId: 1, severity: "low", status: "open", summary: "Route confirmation requires review", detectedAt: new Date().toISOString() }
];

// getDashboardSummary は dashboard KPI を取得します。
export async function getDashboardSummary() {
  try {
    return await fetchJson<DashboardSummary>("/api/dashboard/summary");
  } catch {
    return fallbackSummary;
  }
}

// getVehicles は車両一覧を取得します。検索条件は query string に変換します。
export async function getVehicles(status = "", q = "") {
  try {
    return await fetchJson<Vehicle[]>(`/api/vehicles?status=${encodeURIComponent(status)}&q=${encodeURIComponent(q)}`);
  } catch {
    return fallbackVehicles;
  }
}

// getDrivers はドライバー一覧を取得します。
export async function getDrivers(status = "", q = "") {
  try {
    return await fetchJson<Driver[]>(`/api/drivers?status=${encodeURIComponent(status)}&q=${encodeURIComponent(q)}`);
  } catch {
    return fallbackDrivers;
  }
}

// getOperationTasks は運行タスク一覧を取得します。
export async function getOperationTasks(status = "", q = "") {
  try {
    return await fetchJson<OperationTask[]>(`/api/operation-tasks?status=${encodeURIComponent(status)}&q=${encodeURIComponent(q)}`);
  } catch {
    return fallbackTasks;
  }
}

// getIncidents は異常イベント一覧を取得します。
export async function getIncidents(status = "", q = "") {
  try {
    return await fetchJson<Incident[]>(`/api/incidents?status=${encodeURIComponent(status)}&q=${encodeURIComponent(q)}`);
  } catch {
    return fallbackIncidents;
  }
}
