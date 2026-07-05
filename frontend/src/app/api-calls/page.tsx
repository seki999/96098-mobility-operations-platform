// ApiCallsPage は README と連動する API 呼び出し例を表示します。
export default function ApiCallsPage() {
  return (
    <>
      <div className="page-title">
        <div>
          <h1>API Calling Guide</h1>
          <p>Frontend から Go REST API を呼び出す設計を確認できます。</p>
        </div>
      </div>
      <pre className="api-box">{`curl http://localhost:8080/health
curl http://localhost:8080/api/vehicles
curl -X POST http://localhost:8080/api/incidents \\
  -H "Content-Type: application/json" \\
  -d '{"incidentCode":"INC-96098-NEW","operationTaskId":1,"severity":"low","status":"open","summary":"neutral test incident"}'`}</pre>
    </>
  );
}
