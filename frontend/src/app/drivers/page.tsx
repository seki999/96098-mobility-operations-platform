import { Badge, DataTable } from "@/components/DataTable";
import { FilterBar } from "@/components/FilterBar";
import { getDrivers } from "@/lib/api";
import { Suspense } from "react";

// DriversPage renders anonymized driver availability and skill coverage.
export default async function DriversPage({ searchParams }: { searchParams: { status?: string; q?: string } }) {
  const drivers = await getDrivers(searchParams.status, searchParams.q);
  const available = drivers.filter((driver) => driver.status === "available").length;
  const assigned = drivers.filter((driver) => driver.status === "assigned").length;

  return (
    <>
      <div className="page-title">
        <div>
          <h1>Driver Management</h1>
          <p>Track anonymized operator availability, skill level, and service-area coverage.</p>
        </div>
      </div>
      <section className="grid stats" style={{ gridTemplateColumns: "repeat(3, minmax(0, 1fr))", marginBottom: 16 }}>
        <div className="card stat-card"><div className="metric">Registered</div><div className="metric-value">{drivers.length}</div><div className="delta">anonymized operators</div></div>
        <div className="card stat-card"><div className="metric">Available</div><div className="metric-value">{available}</div><div className="delta">ready for assignment</div></div>
        <div className="card stat-card"><div className="metric">Assigned</div><div className="metric-value">{assigned}</div><div className="delta">currently allocated</div></div>
      </section>
      <Suspense fallback={<div className="toolbar">Loading filters...</div>}>
        <FilterBar statuses={["available", "assigned", "off_shift"]} />
      </Suspense>
      <DataTable
        rows={drivers}
        columns={[
          { key: "code", label: "Code", render: (row) => <span className="code-cell">{row.code}</span> },
          { key: "displayName", label: "Display Name" },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "skillLevel", label: "Skill" },
          { key: "serviceAreaId", label: "Area" }
        ]}
      />
    </>
  );
}
