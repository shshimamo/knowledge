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

// ナレッジ編集ページに移動
export async function gotoKnowledgeEdit(page: Page, knowledgeId: string) {
  await page.goto(`/knowledge/${knowledgeId}/edit`);
}

// ナレッジを編集する
export async function editKnowledge(page: Page, knowledgeId: string, title: string, text: string, isPublic = false) {
  await gotoKnowledgeEdit(page, knowledgeId);
  
  // フォームをクリア
  await page.fill('.e2e-model-knowledge-edit-input-title', '');
  await page.fill('.e2e-model-knowledge-edit-input-text', '');
  
  // 新しい値を入力
  await page.fill('.e2e-model-knowledge-edit-input-title', title);
  await page.fill('.e2e-model-knowledge-edit-input-text', text);
  
  // パブリック設定
  const checkbox = page.locator('.e2e-model-knowledge-edit-input-public');
  if (isPublic) {
    await checkbox.check();
  } else {
    await checkbox.uncheck();
  }
  
  // 保存ボタンをクリック
  await page.click('.e2e-model-knowledge-edit-save-button');
  
  // ナレッジ詳細ページに遷移するまで待機
  await page.waitForURL(`/knowledge/${knowledgeId}`);
}

// ナレッジを削除する
export async function deleteKnowledge(page: Page, knowledgeTitle: string) {
  await gotoKnowledgeList(page);
  
  // 削除対象のナレッジアイテムを見つける
  const knowledgeItems = page.locator('.e2e-model-knowledge-list-item');
  const targetItem = knowledgeItems.filter({ hasText: knowledgeTitle });
  
  // 削除ボタンをクリック
  await targetItem.locator('.e2e-model-knowledge-list-item-delete-button').click();
  
  // 削除確認モーダルの「Yes」ボタンをクリック
  await page.locator('.e2e-ui-modal-button-yes').click();
  
  // モーダルが閉じるまで待機
  await page.waitForTimeout(1000);
}

// ナレッジ詳細ページに移動
export async function gotoKnowledgeDetail(page: Page, knowledgeId: string) {
  await page.goto(`/knowledge/${knowledgeId}`);
}

// 指定されたタイトルのナレッジIDを取得する（一覧ページから）
export async function getKnowledgeIdByTitle(page: Page, title: string): Promise<string> {
  await gotoKnowledgeList(page);
  
  // タイトルを含むナレッジアイテムのShowボタンをクリックしてIDを取得
  const knowledgeItems = page.locator('.e2e-model-knowledge-list-item');
  const targetItem = knowledgeItems.filter({ hasText: title });
  await targetItem.locator('.e2e-model-knowledge-list-item-show-button').click();
  
  // URLからIDを抽出
  await page.waitForURL(/\/knowledge\/\d+/);
  const url = page.url();
  const match = url.match(/\/knowledge\/(\d+)/);
  return match ? match[1] : '';
}

// ナレッジ一覧にナレッジが存在するかチェック
export async function checkKnowledgeExistsInList(page: Page, title: string): Promise<boolean> {
  await gotoKnowledgeList(page);
  const knowledgeList = page.locator('.e2e-model-knowledge-list-all');
  return await knowledgeList.locator('text=' + title).isVisible();
}

// ナレッジ一覧でナレッジが存在しないことをチェック
export async function checkKnowledgeNotExistsInList(page: Page, title: string): Promise<boolean> {
  await gotoKnowledgeList(page);
  const knowledgeList = page.locator('.e2e-model-knowledge-list-all');
  return !(await knowledgeList.locator('text=' + title).isVisible());
}

export async function clearTestData(page: Page) {
  // テスト用データのクリーンアップ（必要に応じて実装）
  // データベースの初期化処理など
}