import { test, expect } from '@playwright/test';
import { signupUser, createKnowledge, gotoKnowledgeCreate, gotoKnowledgeList } from './helpers/test-helper';

test.describe('ナレッジ作成機能', () => {
  test.beforeEach(async ({ page }) => {
    // 各テスト前にユニークなユーザーでサインアップ
    const uniqueUser = {
      email: `test${Date.now()}@example.com`,
      password: 'testpassword123',
      username: `testuser${Date.now()}`
    };

    await signupUser(page, uniqueUser);
    await expect(page).toHaveURL('/knowledge_list');
  });

  test('ナレッジ作成ページが正しく表示される', async ({ page }) => {
    // ナレッジ作成ページに移動
    await gotoKnowledgeCreate(page)

    // フォーム要素が存在することを確認
    await expect(page.locator('.e2e-model-knowledge-edit-input-title')).toBeVisible();
    await expect(page.locator('.e2e-model-knowledge-edit-input-text')).toBeVisible();
    await expect(page.locator('.e2e-model-knowledge-edit-input-public')).toBeVisible();
    await expect(page.locator('button[type="submit"]')).toBeVisible();
  });

  test('プライベートナレッジを作成できる', async ({ page }) => {
    const knowledgeTitle = `テストナレッジ ${Date.now()}`;
    const knowledgeText = 'これはテスト用のナレッジ内容です。';

    // ナレッジ作成
    await createKnowledge(page, knowledgeTitle, knowledgeText, false);

    // ナレッジ詳細ページに遷移していることを確認
    await expect(page).toHaveURL(/\/knowledge\/\d+/);

    // ナレッジ内容が正しく表示されることを確認
    await expect(page.locator('.e2e-model-knowledge-detail-title')).toContainText(knowledgeTitle);
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText(knowledgeText);

    // プライベートナレッジであることを確認
    await expect(page.locator('.e2e-model-knowledge-detail-public')).toContainText('No');
  });

  test('パブリックナレッジを作成できる', async ({ page }) => {
    const knowledgeTitle = `パブリックナレッジ ${Date.now()}`;
    const knowledgeText = 'これは公開ナレッジです。';

    // ナレッジ作成
    await createKnowledge(page, knowledgeTitle, knowledgeText, true);

    // ナレッジ詳細ページに遷移していることを確認
    await expect(page).toHaveURL(/\/knowledge\/\d+/);

    // ナレッジ内容が正しく表示されることを確認
    await expect(page.locator('.e2e-model-knowledge-detail-title')).toContainText(knowledgeTitle);
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText(knowledgeText);

    // パブリックナレッジであることを確認
    await expect(page.locator('.e2e-model-knowledge-detail-public')).toContainText('Yes');
  });

  test('ナレッジ作成後、一覧ページで確認できる', async ({ page }) => {
    const knowledgeTitle = `一覧確認ナレッジ ${Date.now()}`;
    const knowledgeText = 'このナレッジが一覧に表示されるかテストします。';

    // ナレッジ作成
    await createKnowledge(page, knowledgeTitle, knowledgeText, false);

    // ナレッジ一覧ページに移動
    await gotoKnowledgeList(page)

    // 作成したナレッジが一覧に表示されることを確認
    await expect(page.locator(`.e2e-model-knowledge-list-all`)).toContainText(knowledgeTitle);
  });
});