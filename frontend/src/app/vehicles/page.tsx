import { Badge, DataTable } from "@/components/DataTable";
import { FilterBar } from "@/components/FilterBar";
import { getVehicles } from "@/lib/api";
import { Suspense } from "react";

// VehiclesPage は車両 CRUD の一覧側 UI を表します。
export default async function VehiclesPage({ searchParams }: { searchParams: { status?: string; q?: string } }) {
  const vehicles = await getVehicles(searchParams.status, searchParams.q);
  return (
    <>
      <div className="page-title">
        <div>
          <h1>Vehicle Management</h1>
          <p>車両タイプ、稼働状態、サービスエリア、稼働率を管理します。</p>
        </div>
      </div>
      <Suspense fallback={<div className="toolbar">Loading filters...</div>}>
        <FilterBar statuses={["active", "maintenance", "standby"]} />
      </Suspense>
      <DataTable
        rows={vehicles}
        columns={[
          { key: "code", label: "Code" },
          { key: "type", label: "Type" },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "serviceAreaId", label: "Area" },
          { key: "utilizationPct", label: "Utilization", render: (row) => `${row.utilizationPct}%` }
        ]}
      />
    </>
  );
}
