import "./globals.css";
import Link from "next/link";

export const metadata = {
  title: "96098 Mobility Operations Platform",
  description: "Portfolio project for mobility operation management"
};

const links = [
  ["Dashboard", "/"],
  ["Vehicles", "/vehicles"],
  ["Drivers", "/drivers"],
  ["Operation Tasks", "/operation-tasks"],
  ["Incidents", "/incidents"],
  ["API", "/api-calls"]
];

// RootLayout は全画面で共通のナビゲーションを提供します。
export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body>
        <div className="shell">
          <aside className="sidebar">
            <div className="brand">96098 Mobility Operations Platform</div>
            <nav className="nav" aria-label="Primary navigation">
              {links.map(([label, href]) => (
                <Link href={href} key={href}>
                  {label}
                </Link>
              ))}
            </nav>
          </aside>
          <main className="main">{children}</main>
        </div>
      </body>
    </html>
  );
}
