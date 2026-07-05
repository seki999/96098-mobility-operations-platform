import { Badge, DataTable } from "@/components/DataTable";
import { getDashboardSummary, getOperationTasks, getVehicles } from "@/lib/api";

// Home renders the operational overview used for portfolio screenshots.
export default async function Home() {
  const [summary, vehicles, tasks] = await Promise.all([getDashboardSummary(), getVehicles(), getOperationTasks()]);
  const stats = [
    ["Fleet size", summary.totalVehicles, "registered units"],
    ["Active fleet", summary.activeVehicles, "ready for dispatch"],
    ["Drivers", summary.availableDrivers, "available now"],
    ["Open tasks", summary.openTasks, "planned or running"],
    ["Incidents", summary.openIncidents, "review queue"]
  ] as const;

  return (
    <>
      <div className="page-title">
        <div>
          <h1>Operations Overview</h1>
          <p>Fleet availability, task pressure, and incident status for a neutral mobility operation.</p>
        </div>
      </div>
      <section className="grid stats">
        {stats.map(([label, value, note]) => (
          <div className="card stat-card" key={label}>
            <div className="metric">{label}</div>
            <div className="metric-value">{value}</div>
            <div className="delta">{note}</div>
          </div>
        ))}
      </section>
      <section className="dashboard-grid">
        <div className="card ops-map">
          <span className="route route-a" />
          <span className="route route-b" />
          <span className="route route-c" />
          <span className="node node-a">N1</span>
          <span className="node node-b warn">C2</span>
          <span className="node node-c">S3</span>
          <span className="node node-d danger">I1</span>
          <div className="map-caption">
            <div className="metric">Service area telemetry</div>
            <div className="metric-value">{summary.avgUtilizationPct.toFixed(1)}%</div>
            <p>Average vehicle utilization across three neutral service zones.</p>
          </div>
        </div>
        <div className="card">
          <div className="metric">Dispatch activity</div>
          <div className="activity-list" style={{ marginTop: 12 }}>
            {tasks.slice(0, 3).map((task) => (
              <div className="activity-item" key={task.id}>
                <strong>{task.taskCode}</strong>
                <span>
                  <Badge value={task.status} /> <Badge value={task.priority} />
                </span>
                <p>{task.notes}</p>
              </div>
            ))}
          </div>
        </div>
      </section>
      <h2>Recent Operation Tasks</h2>
      <DataTable
        rows={tasks}
        columns={[
          { key: "taskCode", label: "Task", render: (row) => <span className="code-cell">{row.taskCode}</span> },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "priority", label: "Priority", render: (row) => <Badge value={row.priority} /> },
          { key: "notes", label: "Notes" }
        ]}
      />
      <h2>Vehicle Utilization</h2>
      <DataTable
        rows={vehicles}
        columns={[
          { key: "code", label: "Vehicle", render: (row) => <span className="code-cell">{row.code}</span> },
          { key: "type", label: "Type" },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "utilizationPct", label: "Utilization", render: (row) => `${row.utilizationPct}%` }
        ]}
      />
    </>
  );
}
