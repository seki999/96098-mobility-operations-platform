import { Badge } from "@/components/DataTable";
import { fallbackTasks, getOperationTasks } from "@/lib/api";

// OperationTaskDetailPage は運行タスクの詳細ビューです。
export default async function OperationTaskDetailPage({ params }: { params: { id: string } }) {
  const tasks = await getOperationTasks();
  const task = tasks.find((item) => item.id === Number(params.id)) ?? fallbackTasks[0];
  return (
    <>
      <div className="page-title">
        <div>
          <h1>{task.taskCode}</h1>
          <p>運行タスクの担当車両、担当者、優先度、メモを確認します。</p>
        </div>
      </div>
      <div className="grid" style={{ gridTemplateColumns: "repeat(2, minmax(0, 1fr))" }}>
        <div className="card"><div className="metric">Status</div><div className="metric-value"><Badge value={task.status} /></div></div>
        <div className="card"><div className="metric">Priority</div><div className="metric-value"><Badge value={task.priority} /></div></div>
        <div className="card"><div className="metric">Vehicle ID</div><div className="metric-value">{task.vehicleId}</div></div>
        <div className="card"><div className="metric">Driver ID</div><div className="metric-value">{task.driverId}</div></div>
      </div>
      <div className="card" style={{ marginTop: 18 }}>
        <div className="metric">Notes</div>
        <p>{task.notes}</p>
      </div>
    </>
  );
}
