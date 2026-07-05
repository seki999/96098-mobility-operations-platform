-- 96098 mobility operations platform 初期スキーマ
-- 実在企業・実在車両・実在個人情報を扱わない Portfolio 用 DB 設計です。

CREATE TABLE IF NOT EXISTS service_areas (
  id BIGSERIAL PRIMARY KEY,
  code VARCHAR(40) NOT NULL UNIQUE,
  name VARCHAR(120) NOT NULL,
  city_label VARCHAR(120) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS vehicles (
  id BIGSERIAL PRIMARY KEY,
  code VARCHAR(40) NOT NULL UNIQUE,
  type VARCHAR(40) NOT NULL,
  status VARCHAR(40) NOT NULL,
  service_area_id BIGINT NOT NULL REFERENCES service_areas(id),
  utilization_pct NUMERIC(5,2) NOT NULL DEFAULT 0,
  last_checked_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT vehicles_utilization_range CHECK (utilization_pct >= 0 AND utilization_pct <= 100)
);

CREATE TABLE IF NOT EXISTS drivers (
  id BIGSERIAL PRIMARY KEY,
  code VARCHAR(40) NOT NULL UNIQUE,
  display_name VARCHAR(120) NOT NULL,
  status VARCHAR(40) NOT NULL,
  skill_level VARCHAR(40) NOT NULL,
  service_area_id BIGINT NOT NULL REFERENCES service_areas(id),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS operation_tasks (
  id BIGSERIAL PRIMARY KEY,
  task_code VARCHAR(40) NOT NULL UNIQUE,
  vehicle_id BIGINT NOT NULL REFERENCES vehicles(id),
  driver_id BIGINT NOT NULL REFERENCES drivers(id),
  service_area_id BIGINT NOT NULL REFERENCES service_areas(id),
  status VARCHAR(40) NOT NULL,
  priority VARCHAR(40) NOT NULL,
  scheduled_at TIMESTAMPTZ NOT NULL,
  notes TEXT NOT NULL DEFAULT '',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS incidents (
  id BIGSERIAL PRIMARY KEY,
  incident_code VARCHAR(40) NOT NULL UNIQUE,
  operation_task_id BIGINT NOT NULL REFERENCES operation_tasks(id),
  severity VARCHAR(40) NOT NULL,
  status VARCHAR(40) NOT NULL,
  summary TEXT NOT NULL,
  detected_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS operation_logs (
  id BIGSERIAL PRIMARY KEY,
  actor_code VARCHAR(80) NOT NULL,
  action VARCHAR(120) NOT NULL,
  target_type VARCHAR(80) NOT NULL,
  target_id BIGINT,
  message TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_vehicles_status ON vehicles(status);
CREATE INDEX IF NOT EXISTS idx_drivers_status ON drivers(status);
CREATE INDEX IF NOT EXISTS idx_operation_tasks_status ON operation_tasks(status);
CREATE INDEX IF NOT EXISTS idx_incidents_status ON incidents(status);
