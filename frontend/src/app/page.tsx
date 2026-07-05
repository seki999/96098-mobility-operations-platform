import { Badge, DataTable } from "@/components/DataTable";
import { getDashboardSummary, getOperationTasks, getVehicles } from "@/lib/api";

// Home は運用 Dashboard として主要 KPI と直近タスクを表示します。
export default async function Home() {
  const [summary, vehicles, tasks] = await Promise.all([getDashboardSummary(), getVehicles(), getOperationTasks()]);
  return (
    <>
      <div className="page-title">
        <div>
          <h1>Operations Dashboard</h1>
          <p>移動出行運営管理平台の稼働率、タスク、異常状態を俯瞰します。</p>
        </div>
      </div>
      <section className="grid stats">
        {[
          ["Vehicles", summary.totalVehicles],
          ["Active", summary.activeVehicles],
          ["Drivers", summary.availableDrivers],
          ["Open Tasks", summary.openTasks],
          ["Incidents", summary.openIncidents]
        ].map(([label, value]) => (
          <div className="card" key={label}>
            <div className="metric">{label}</div>
            <div className="metric-value">{value}</div>
          </div>
        ))}
      </section>
      <div className="grid" style={{ marginTop: 18, gridTemplateColumns: "1.2fr 1fr" }}>
        <div className="card">
          <div className="metric">Average utilization</div>
          <div className="metric-value">{summary.avgUtilizationPct.toFixed(1)}%</div>
          <p>車両稼働率を運用判断の入口 KPI として表示します。</p>
        </div>
        <div className="card">
          <div className="metric">API endpoint</div>
          <pre className="api-box">GET /api/dashboard/summary</pre>
        </div>
      </div>
      <h2>Recent Operation Tasks</h2>
      <DataTable
        rows={tasks}
        columns={[
          { key: "taskCode", label: "Task" },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "priority", label: "Priority", render: (row) => <Badge value={row.priority} /> },
          { key: "notes", label: "Notes" }
        ]}
      />
      <h2>Vehicle Utilization</h2>
      <DataTable
        rows={vehicles}
        columns={[
          { key: "code", label: "Vehicle" },
          { key: "type", label: "Type" },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "utilizationPct", label: "Utilization", render: (row) => `${row.utilizationPct}%` }
        ]}
      />
    </>
  );
}
