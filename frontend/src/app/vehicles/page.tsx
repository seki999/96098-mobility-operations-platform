import { Badge, DataTable } from "@/components/DataTable";
import { FilterBar } from "@/components/FilterBar";
import { getVehicles } from "@/lib/api";
import { Suspense } from "react";

// VehiclesPage renders fleet inventory and utilization controls.
export default async function VehiclesPage({ searchParams }: { searchParams: { status?: string; q?: string } }) {
  const vehicles = await getVehicles(searchParams.status, searchParams.q);
  const active = vehicles.filter((vehicle) => vehicle.status === "active").length;
  const average = vehicles.reduce((sum, vehicle) => sum + vehicle.utilizationPct, 0) / Math.max(vehicles.length, 1);

  return (
    <>
      <div className="page-title">
        <div>
          <h1>Fleet Management</h1>
          <p>Monitor vehicle readiness, operating zone, and utilization pressure.</p>
        </div>
      </div>
      <section className="grid stats" style={{ gridTemplateColumns: "repeat(3, minmax(0, 1fr))", marginBottom: 16 }}>
        <div className="card stat-card"><div className="metric">Inventory</div><div className="metric-value">{vehicles.length}</div><div className="delta">neutral fleet records</div></div>
        <div className="card stat-card"><div className="metric">Active</div><div className="metric-value">{active}</div><div className="delta">dispatch capable</div></div>
        <div className="card stat-card"><div className="metric">Average use</div><div className="metric-value">{average.toFixed(1)}%</div><div className="delta">rolling operations KPI</div></div>
      </section>
      <Suspense fallback={<div className="toolbar">Loading filters...</div>}>
        <FilterBar statuses={["active", "maintenance", "standby"]} />
      </Suspense>
      <DataTable
        rows={vehicles}
        columns={[
          { key: "code", label: "Code", render: (row) => <span className="code-cell">{row.code}</span> },
          { key: "type", label: "Type" },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "serviceAreaId", label: "Area" },
          { key: "utilizationPct", label: "Utilization", render: (row) => `${row.utilizationPct}%` }
        ]}
      />
    </>
  );
}
