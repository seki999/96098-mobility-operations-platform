import Link from "next/link";
import { Badge, DataTable } from "@/components/DataTable";
import { FilterBar } from "@/components/FilterBar";
import { getOperationTasks } from "@/lib/api";
import { Suspense } from "react";

// OperationTasksPage は運行タスク一覧と詳細画面への導線を提供します。
export default async function OperationTasksPage({ searchParams }: { searchParams: { status?: string; q?: string } }) {
  const tasks = await getOperationTasks(searchParams.status, searchParams.q);
  return (
    <>
      <div className="page-title">
        <div>
          <h1>Operation Tasks</h1>
          <p>配車、点検、エリア調整などの運転関連タスクを管理します。</p>
        </div>
      </div>
      <Suspense fallback={<div className="toolbar">Loading filters...</div>}>
        <FilterBar statuses={["planned", "in_progress", "completed"]} />
      </Suspense>
      <DataTable
        rows={tasks}
        columns={[
          { key: "taskCode", label: "Task", render: (row) => <Link href={`/operation-tasks/${row.id}`}>{row.taskCode}</Link> },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "priority", label: "Priority", render: (row) => <Badge value={row.priority} /> },
          { key: "vehicleId", label: "Vehicle" },
          { key: "notes", label: "Notes" }
        ]}
      />
    </>
  );
}
