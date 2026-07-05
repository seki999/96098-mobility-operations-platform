import { Badge, DataTable } from "@/components/DataTable";
import { FilterBar } from "@/components/FilterBar";
import { getIncidents } from "@/lib/api";
import { Suspense } from "react";

// IncidentsPage は異常イベントの状況を一覧表示します。
export default async function IncidentsPage({ searchParams }: { searchParams: { status?: string; q?: string } }) {
  const incidents = await getIncidents(searchParams.status, searchParams.q);
  return (
    <>
      <div className="page-title">
        <div>
          <h1>Incident Management</h1>
          <p>運行中の遅延、端末応答、確認待ちなどの異常イベントを追跡します。</p>
        </div>
      </div>
      <Suspense fallback={<div className="toolbar">Loading filters...</div>}>
        <FilterBar statuses={["open", "investigating", "resolved"]} />
      </Suspense>
      <DataTable
        rows={incidents}
        columns={[
          { key: "incidentCode", label: "Incident" },
          { key: "severity", label: "Severity", render: (row) => <Badge value={row.severity} /> },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "summary", label: "Summary" }
        ]}
      />
    </>
  );
}
