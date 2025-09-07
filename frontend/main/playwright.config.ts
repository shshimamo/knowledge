import { defineConfig, devices } from '@playwright/test';

export default defineConfig({
  testDir: './tests/e2e',
  timeout: 10000,
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  workers: process.env.CI ? 1 : undefined,
  reporter: 'html',
  use: {
    baseURL: process.env.DEVCONTAINER ? 'http://host.docker.internal:3000' : 'http://localhost:3000',
    trace: 'on-first-retry',
    screenshot: 'only-on-failure',
    ...(process.env.DEVCONTAINER && {
      launchOptions: {
        // devcontainer では Chromium から host.docker.internal にアクセスできるようにする
        args: ['--disable-web-security', '--disable-features=VizDisplayCompositor'],
      },
    }),
  },

  projects: [
    {
      name: 'chromium',
      use: { ...devices['Desktop Chrome'] },
    },
  ],

  // Docker Composeで起動されたサーバーを使用するため、webServerは無効化
  // webServer: [
  //   {
  //     command: 'npm run dev',
  //     port: 3000,
  //     reuseExistingServer: true,
  //   },
  // ],
});
