-- Portfolio 用の中立的な初期データです。実在の企業・車両・個人とは関係ありません。

INSERT INTO service_areas (id, code, name, city_label) VALUES
  (1, 'AREA-NORTH', 'North Mobility Zone', 'Neutral City North'),
  (2, 'AREA-CENTRAL', 'Central Mobility Zone', 'Neutral City Central'),
  (3, 'AREA-SOUTH', 'South Mobility Zone', 'Neutral City South')
ON CONFLICT (id) DO NOTHING;

INSERT INTO vehicles (id, code, type, status, service_area_id, utilization_pct, last_checked_at) VALUES
  (1, 'VH-96098-001', 'compact-ev', 'active', 1, 78.5, now() - interval '20 minutes'),
  (2, 'VH-96098-002', 'accessible-van', 'maintenance', 2, 35.0, now() - interval '2 hours'),
  (3, 'VH-96098-003', 'standard-ev', 'active', 2, 91.2, now() - interval '15 minutes'),
  (4, 'VH-96098-004', 'standard-ev', 'standby', 3, 42.4, now() - interval '55 minutes')
ON CONFLICT (id) DO NOTHING;

INSERT INTO drivers (id, code, display_name, status, skill_level, service_area_id) VALUES
  (1, 'DR-96098-A01', 'Operator Alpha', 'available', 'senior', 1),
  (2, 'DR-96098-B02', 'Operator Bravo', 'assigned', 'standard', 2),
  (3, 'DR-96098-C03', 'Operator Charlie', 'available', 'standard', 3),
  (4, 'DR-96098-D04', 'Operator Delta', 'off_shift', 'lead', 2)
ON CONFLICT (id) DO NOTHING;

INSERT INTO operation_tasks (id, task_code, vehicle_id, driver_id, service_area_id, status, priority, scheduled_at, notes) VALUES
  (1, 'TASK-96098-1001', 1, 1, 1, 'planned', 'high', now() + interval '30 minutes', 'Airport-style hub transfer simulation'),
  (2, 'TASK-96098-1002', 3, 2, 2, 'in_progress', 'medium', now() - interval '10 minutes', 'Central zone recurring operation'),
  (3, 'TASK-96098-1003', 4, 3, 3, 'completed', 'low', now() - interval '3 hours', 'South zone availability balancing')
ON CONFLICT (id) DO NOTHING;

INSERT INTO incidents (id, incident_code, operation_task_id, severity, status, summary, detected_at) VALUES
  (1, 'INC-96098-501', 2, 'medium', 'investigating', 'Delayed status update from vehicle terminal', now() - interval '12 minutes'),
  (2, 'INC-96098-502', 1, 'low', 'open', 'Route confirmation requires operator review', now() - interval '25 minutes')
ON CONFLICT (id) DO NOTHING;

INSERT INTO operation_logs (actor_code, action, target_type, target_id, message) VALUES
  ('system-seed', 'create', 'vehicle', 1, 'Initial neutral vehicle record created'),
  ('system-seed', 'create', 'task', 1, 'Initial neutral operation task created');

SELECT setval('service_areas_id_seq', 3, true);
SELECT setval('vehicles_id_seq', 4, true);
SELECT setval('drivers_id_seq', 4, true);
SELECT setval('operation_tasks_id_seq', 3, true);
SELECT setval('incidents_id_seq', 2, true);
