import { Badge, DataTable } from "@/components/DataTable";
import { FilterBar } from "@/components/FilterBar";
import { getDrivers } from "@/lib/api";
import { Suspense } from "react";

// DriversPage はドライバーの稼働状態を確認する画面です。
export default async function DriversPage({ searchParams }: { searchParams: { status?: string; q?: string } }) {
  const drivers = await getDrivers(searchParams.status, searchParams.q);
  return (
    <>
      <div className="page-title">
        <div>
          <h1>Driver Management</h1>
          <p>仮名コード化した運転担当者の状態と担当エリアを管理します。</p>
        </div>
      </div>
      <Suspense fallback={<div className="toolbar">Loading filters...</div>}>
        <FilterBar statuses={["available", "assigned", "off_shift"]} />
      </Suspense>
      <DataTable
        rows={drivers}
        columns={[
          { key: "code", label: "Code" },
          { key: "displayName", label: "Display Name" },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "skillLevel", label: "Skill" },
          { key: "serviceAreaId", label: "Area" }
        ]}
      />
    </>
  );
}
