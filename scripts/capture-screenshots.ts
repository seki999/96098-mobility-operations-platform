import { chromium } from "playwright";

const routes = [
  ["/", "home.png"],
  ["/vehicles", "vehicles.png"],
  ["/drivers", "drivers.png"],
  ["/operation-tasks", "operation-tasks.png"],
  ["/", "dashboard.png"]
] as const;

// capture は README 用スクリーンショットを指定パスに保存します。
async function capture() {
  const baseURL = process.env.FRONTEND_URL ?? "http://localhost:3000";
  const browser = await chromium.launch();
  const page = await browser.newPage({ viewport: { width: 1440, height: 960 } });

  for (const [route, fileName] of routes) {
    await page.goto(`${baseURL}${route}`, { waitUntil: "networkidle" });
    await page.screenshot({ path: `screenshots/${fileName}`, fullPage: true });
  }

  await browser.close();
}

capture().catch((error) => {
  console.error(error);
  process.exit(1);
});
