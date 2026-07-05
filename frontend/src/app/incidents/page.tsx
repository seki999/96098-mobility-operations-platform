import { Badge, DataTable } from "@/components/DataTable";
import { FilterBar } from "@/components/FilterBar";
import { getIncidents } from "@/lib/api";
import { Suspense } from "react";

// IncidentsPage renders neutral incident tracking records.
export default async function IncidentsPage({ searchParams }: { searchParams: { status?: string; q?: string } }) {
  const incidents = await getIncidents(searchParams.status, searchParams.q);
  const open = incidents.filter((incident) => incident.status === "open").length;
  const investigating = incidents.filter((incident) => incident.status === "investigating").length;

  return (
    <>
      <div className="page-title">
        <div>
          <h1>Incident Management</h1>
          <p>Review delayed status updates, routing checks, and operational exceptions.</p>
        </div>
      </div>
      <section className="grid stats" style={{ gridTemplateColumns: "repeat(3, minmax(0, 1fr))", marginBottom: 16 }}>
        <div className="card stat-card"><div className="metric">Total events</div><div className="metric-value">{incidents.length}</div><div className="delta">neutral records</div></div>
        <div className="card stat-card"><div className="metric">Open</div><div className="metric-value">{open}</div><div className="delta">needs triage</div></div>
        <div className="card stat-card"><div className="metric">Investigating</div><div className="metric-value">{investigating}</div><div className="delta">operator review</div></div>
      </section>
      <Suspense fallback={<div className="toolbar">Loading filters...</div>}>
        <FilterBar statuses={["open", "investigating", "resolved"]} />
      </Suspense>
      <DataTable
        rows={incidents}
        columns={[
          { key: "incidentCode", label: "Incident", render: (row) => <span className="code-cell">{row.incidentCode}</span> },
          { key: "severity", label: "Severity", render: (row) => <Badge value={row.severity} /> },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "summary", label: "Summary" }
        ]}
      />
    </>
  );
}
