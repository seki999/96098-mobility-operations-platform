import "./globals.css";
import Link from "next/link";

export const metadata = {
  title: "96098 Mobility Operations Platform",
  description: "Portfolio project for mobility operation management"
};

const links = [
  ["Overview", "/"],
  ["Fleet", "/vehicles"],
  ["Drivers", "/drivers"],
  ["Tasks", "/operation-tasks"],
  ["Incidents", "/incidents"],
  ["API", "/api-calls"]
];

// RootLayout provides the shared operations console frame.
export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body>
        <div className="shell">
          <aside className="sidebar">
            <div className="brand">
              <span className="brand-mark">96</span>
              <span>
                Mobility Ops
                <small>Case 96098</small>
              </span>
            </div>
            <nav className="nav" aria-label="Primary navigation">
              {links.map(([label, href]) => (
                <Link href={href} key={href}>
                  {label}
                </Link>
              ))}
            </nav>
            <div className="sidebar-panel">
              <span>Environment</span>
              <strong>Local Portfolio</strong>
              <small>No real company or customer data</small>
            </div>
          </aside>
          <main className="main">
            <header className="topbar">
              <div>
                <span className="eyebrow">Operations control center</span>
                <strong>Neutral mobility support platform</strong>
              </div>
              <div className="topbar-actions">
                <span className="sync-dot" />
                <span>API fallback ready</span>
              </div>
            </header>
            {children}
          </main>
        </div>
      </body>
    </html>
  );
}
