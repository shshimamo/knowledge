import { Page } from '@playwright/test';

export const TEST_USER = {
  email: 'test@example.com',
  password: 'testpassword123',
  username: 'testuser'
};

export async function signupUser(page: Page, userData = TEST_USER) {
  await page.goto('/signup');

  // サインアップフォームに入力
  await page.fill('.e2e-page-signup-input-email', userData.email);
  await page.fill('.e2e-page-signup-input-password', userData.password);
  await page.fill('.e2e-page-signup-input-name', userData.username);

  // サインアップボタンをクリック
  await page.click('button[type="submit"]');

  // ページ遷移を待機
  await page.waitForURL('/knowledge_list');
}

export async function loginUser(page: Page, userData = TEST_USER) {
  await page.goto('/signin');

  // ログインフォームに入力
  await page.fill('input[name="email"]', userData.email);
  await page.fill('input[name="password"]', userData.password);

  // ログインボタンをクリック
  await page.click('button[type="submit"]');

  // ページ遷移を待機
  await page.waitForURL('/knowledge_list');
}

// ナレッジ一覧ページに移動
export async function gotoKnowledgeList(page: Page) {
  await page.goto('/knowledge_list');
}

// ナレッジ作成ページに移動
export async function gotoKnowledgeCreate(page: Page){
  await gotoKnowledgeList(page)
  await page.click('.e2e-model-knowledge-list-create-button')
}

export async function createKnowledge(page: Page, title: string, text: string, isPublic = false) {
  // ナレッジ作成ページに移動
  await gotoKnowledgeCreate(page)

  // フォームに入力
  await page.fill('.e2e-model-knowledge-edit-input-title', title);
  await page.fill('.e2e-model-knowledge-edit-input-text', text);

  if (isPublic) {
    await page.check('.e2e-model-knowledge-edit-input-public');
  }

  // 作成ボタンをクリック
  await page.click('button[type="submit"]');

  // ナレッジ詳細ページに遷移するまで待機
  await page.waitForURL(/\/knowledge\/\d+/);
}

export async function clearTestData(page: Page) {
  // テスト用データのクリーンアップ（必要に応じて実装）
  // データベースの初期化処理など
}