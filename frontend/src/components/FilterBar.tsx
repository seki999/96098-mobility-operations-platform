"use client";

import { useRouter, useSearchParams } from "next/navigation";
import { FormEvent } from "react";

// FilterBar は status と q を URL query として扱う検索 UI です。
export function FilterBar({ statuses }: { statuses: string[] }) {
  const router = useRouter();
  const params = useSearchParams();

  function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const form = new FormData(event.currentTarget);
    const query = new URLSearchParams();
    const status = String(form.get("status") ?? "");
    const q = String(form.get("q") ?? "");
    if (status) query.set("status", status);
    if (q) query.set("q", q);
    router.push(`?${query.toString()}`);
  }

  return (
    <form className="toolbar" onSubmit={onSubmit}>
      <select name="status" defaultValue={params.get("status") ?? ""} aria-label="status">
        <option value="">All status</option>
        {statuses.map((status) => (
          <option value={status} key={status}>
            {status}
          </option>
        ))}
      </select>
      <input name="q" defaultValue={params.get("q") ?? ""} placeholder="Search code or note" />
      <button type="submit">Search</button>
    </form>
  );
}
