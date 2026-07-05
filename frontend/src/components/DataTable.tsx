import React from "react";

type Column<T> = {
  key: keyof T | string;
  label: string;
  render?: (row: T) => React.ReactNode;
};

// DataTable is a shared dense table for operational list screens.
export function DataTable<T extends { id: number }>({ columns, rows }: { columns: Column<T>[]; rows: T[] }) {
  return (
    <div className="table-wrap">
      <table>
        <thead>
          <tr>
            {columns.map((column) => (
              <th key={String(column.key)}>{column.label}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((row) => (
            <tr key={row.id}>
              {columns.map((column) => (
                <td key={String(column.key)}>
                  {column.render ? column.render(row) : String(row[column.key as keyof T] ?? "")}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

// Badge highlights operational status values without exposing internal details.
export function Badge({ value }: { value: string }) {
  return <span className={`badge ${value}`}>{value}</span>;
}
