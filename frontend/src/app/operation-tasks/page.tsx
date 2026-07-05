import Link from "next/link";
import { Badge, DataTable } from "@/components/DataTable";
import { FilterBar } from "@/components/FilterBar";
import { getOperationTasks } from "@/lib/api";
import { Suspense } from "react";

// OperationTasksPage renders dispatch tasks and links into task details.
export default async function OperationTasksPage({ searchParams }: { searchParams: { status?: string; q?: string } }) {
  const tasks = await getOperationTasks(searchParams.status, searchParams.q);
  const planned = tasks.filter((task) => task.status === "planned").length;
  const running = tasks.filter((task) => task.status === "in_progress").length;

  return (
    <>
      <div className="page-title">
        <div>
          <h1>Operation Tasks</h1>
          <p>Coordinate dispatch, inspection, and service-area balancing tasks.</p>
        </div>
      </div>
      <section className="grid stats" style={{ gridTemplateColumns: "repeat(3, minmax(0, 1fr))", marginBottom: 16 }}>
        <div className="card stat-card"><div className="metric">Task queue</div><div className="metric-value">{tasks.length}</div><div className="delta">visible records</div></div>
        <div className="card stat-card"><div className="metric">Planned</div><div className="metric-value">{planned}</div><div className="delta">awaiting dispatch</div></div>
        <div className="card stat-card"><div className="metric">Running</div><div className="metric-value">{running}</div><div className="delta">in operation</div></div>
      </section>
      <Suspense fallback={<div className="toolbar">Loading filters...</div>}>
        <FilterBar statuses={["planned", "in_progress", "completed"]} />
      </Suspense>
      <DataTable
        rows={tasks}
        columns={[
          { key: "taskCode", label: "Task", render: (row) => <Link className="code-cell" href={`/operation-tasks/${row.id}`}>{row.taskCode}</Link> },
          { key: "status", label: "Status", render: (row) => <Badge value={row.status} /> },
          { key: "priority", label: "Priority", render: (row) => <Badge value={row.priority} /> },
          { key: "vehicleId", label: "Vehicle" },
          { key: "notes", label: "Notes" }
        ]}
      />
    </>
  );
}
