import { test, expect } from '@playwright/test';
import {
  signupUser,
  createKnowledge,
  gotoKnowledgeCreate,
  gotoKnowledgeList,
  editKnowledge,
  deleteKnowledge,
  getKnowledgeIdByTitle,
  gotoKnowledgeDetail,
  checkKnowledgeExistsInList,
  checkKnowledgeNotExistsInList
} from './helpers/test-helper';

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

test.describe('ナレッジ編集機能', () => {
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

  test('ナレッジ編集ページが正しく表示される', async ({ page }) => {
    const knowledgeTitle = `編集テストナレッジ ${Date.now()}`;
    const knowledgeText = 'これは編集テスト用のナレッジです。';

    // ナレッジを作成
    await createKnowledge(page, knowledgeTitle, knowledgeText, false);

    // URLからナレッジIDを取得
    const url = page.url();
    const knowledgeId = url.match(/\/knowledge\/(\d+)/)?.[1] || '';

    // 編集ページに移動
    await page.click('.e2e-model-knowledge-detail-edit-button');

    // 編集フォーム要素が存在することを確認
    await expect(page.locator('.e2e-model-knowledge-edit-input-title')).toBeVisible();
    await expect(page.locator('.e2e-model-knowledge-edit-input-text')).toBeVisible();
    await expect(page.locator('.e2e-model-knowledge-edit-input-public')).toBeVisible();
    await expect(page.locator('.e2e-model-knowledge-edit-save-button')).toBeVisible();
    await expect(page.locator('.e2e-model-knowledge-edit-back-button')).toBeVisible();

    // 既存の値が正しく表示されることを確認
    await expect(page.locator('.e2e-model-knowledge-edit-input-title')).toHaveValue(knowledgeTitle);
    await expect(page.locator('.e2e-model-knowledge-edit-input-text')).toHaveValue(knowledgeText);
  });

  test('ナレッジのタイトルと内容を編集できる', async ({ page }) => {
    const originalTitle = `編集前タイトル ${Date.now()}`;
    const originalText = '編集前の内容です。';
    const newTitle = `編集後タイトル ${Date.now()}`;
    const newText = '編集後の内容です。';

    // ナレッジを作成
    await createKnowledge(page, originalTitle, originalText, false);

    // URLからナレッジIDを取得
    const url = page.url();
    const knowledgeId = url.match(/\/knowledge\/(\d+)/)?.[1] || '';

    // ナレッジを編集
    await editKnowledge(page, knowledgeId, newTitle, newText, false);

    // 編集後の内容が正しく表示されることを確認
    await expect(page.locator('.e2e-model-knowledge-detail-title')).toContainText(newTitle);
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText(newText);
    await expect(page.locator('.e2e-model-knowledge-detail-public')).toContainText('No');
  });

  test('ナレッジの公開設定を変更できる', async ({ page }) => {
    const knowledgeTitle = `公開設定変更テスト ${Date.now()}`;
    const knowledgeText = '公開設定を変更するテストです。';

    // プライベートナレッジを作成
    await createKnowledge(page, knowledgeTitle, knowledgeText, false);

    // URLからナレッジIDを取得
    const url = page.url();
    const knowledgeId = url.match(/\/knowledge\/(\d+)/)?.[1] || '';

    // パブリックに変更
    await editKnowledge(page, knowledgeId, knowledgeTitle, knowledgeText, true);

    // パブリックナレッジに変更されたことを確認
    await expect(page.locator('.e2e-model-knowledge-detail-public')).toContainText('Yes');
  });

  test('編集後、ナレッジ一覧で更新された内容が確認できる', async ({ page }) => {
    const originalTitle = `一覧更新テスト原本 ${Date.now()}`;
    const originalText = '原本の内容です。';
    const newTitle = `一覧更新テスト更新版 ${Date.now()}`;
    const newText = '更新された内容です。';

    // ナレッジを作成
    await createKnowledge(page, originalTitle, originalText, false);

    // URLからナレッジIDを取得
    const url = page.url();
    const knowledgeId = url.match(/\/knowledge\/(\d+)/)?.[1] || '';

    // ナレッジを編集
    await editKnowledge(page, knowledgeId, newTitle, newText, true);

    // ナレッジ一覧ページに移動
    await gotoKnowledgeList(page);

    // 更新されたタイトルが一覧に表示されることを確認
    await expect(page.locator('.e2e-model-knowledge-list-all')).toContainText(newTitle);
    await expect(page.locator('.e2e-model-knowledge-list-all')).toContainText(newText);

    // 古いタイトルは表示されないことを確認
    await expect(page.locator('.e2e-model-knowledge-list-all')).not.toContainText(originalTitle);
  });
});

test.describe('ナレッジ削除機能', () => {
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

  test('削除確認モーダルが正しく表示される', async ({ page }) => {
    const knowledgeTitle = `削除モーダルテスト ${Date.now()}`;
    const knowledgeText = '削除モーダルのテスト用ナレッジです。';

    // ナレッジを作成
    await createKnowledge(page, knowledgeTitle, knowledgeText, false);

    // ナレッジ一覧ページに移動
    await gotoKnowledgeList(page);

    // 作成したナレッジの削除ボタンをクリック
    const knowledgeItems = page.locator('.e2e-model-knowledge-list-item');
    const targetItem = knowledgeItems.filter({ hasText: knowledgeTitle });
    await targetItem.locator('.e2e-model-knowledge-list-item-delete-button').click();

    // 削除確認モーダルが表示されることを確認
    await expect(page.locator('.e2e-ui-modal')).toBeVisible();
    await expect(page.locator('.e2e-ui-modal-text')).toContainText('Delete?');
    await expect(page.locator('.e2e-ui-modal-button-yes')).toBeVisible();
    await expect(page.locator('.e2e-ui-modal-button-no')).toBeVisible();
  });

  test('削除確認モーダルでNoを選択すると削除がキャンセルされる', async ({ page }) => {
    const knowledgeTitle = `削除キャンセルテスト ${Date.now()}`;
    const knowledgeText = '削除キャンセルのテスト用ナレッジです。';

    // ナレッジを作成
    await createKnowledge(page, knowledgeTitle, knowledgeText, false);

    // ナレッジ一覧ページに移動
    await gotoKnowledgeList(page);

    // 削除ボタンをクリック
    const knowledgeItems = page.locator('.e2e-model-knowledge-list-item');
    const targetItem = knowledgeItems.filter({ hasText: knowledgeTitle });
    await targetItem.locator('.e2e-model-knowledge-list-item-delete-button').click();

    // モーダルの「No」ボタンをクリック
    await page.locator('.e2e-ui-modal-button-no').click();

    // モーダルが閉じることを確認
    await page.waitForTimeout(500);
    await expect(page.locator('.e2e-ui-modal')).not.toBeVisible();

    // ナレッジが削除されていないことを確認
    await expect(page.locator('.e2e-model-knowledge-list-all')).toContainText(knowledgeTitle);
  });

  test('ナレッジを削除できる', async ({ page }) => {
    const knowledgeTitle = `削除実行テスト ${Date.now()}`;
    const knowledgeText = '削除実行のテスト用ナレッジです。';

    // ナレッジを作成
    await createKnowledge(page, knowledgeTitle, knowledgeText, false);

    // ナレッジ削除を実行
    await deleteKnowledge(page, knowledgeTitle);

    // ナレッジが一覧から削除されていることを確認
    const exists = await checkKnowledgeExistsInList(page, knowledgeTitle);
    expect(exists).toBe(false);
  });

  test('削除後に一覧ページで該当ナレッジが表示されない', async ({ page }) => {
    const knowledgeTitle1 = `削除テスト1 ${Date.now()}`;
    const knowledgeTitle2 = `削除テスト2 ${Date.now()}`;
    const knowledgeText = '削除テスト用のナレッジです。';

    // 複数のナレッジを作成
    await createKnowledge(page, knowledgeTitle1, knowledgeText, false);
    await gotoKnowledgeList(page);
    await createKnowledge(page, knowledgeTitle2, knowledgeText, false);

    // 1つ目のナレッジを削除
    await deleteKnowledge(page, knowledgeTitle1);

    // 削除されたナレッジが一覧にないことを確認
    await expect(page.locator('.e2e-model-knowledge-list-all')).not.toContainText(knowledgeTitle1);
    // 削除されていないナレッジは残っていることを確認
    await expect(page.locator('.e2e-model-knowledge-list-all')).toContainText(knowledgeTitle2);
  });
});

test.describe('ナレッジ詳細表示機能', () => {
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

  test('ナレッジ詳細ページが正しく表示される', async ({ page }) => {
    const knowledgeTitle = `詳細表示テスト ${Date.now()}`;
    const knowledgeText = 'これは詳細表示のテスト用ナレッジです。';

    // ナレッジを作成
    await createKnowledge(page, knowledgeTitle, knowledgeText, true);

    // 詳細ページの要素が正しく表示されることを確認
    await expect(page.locator('.e2e-model-knowledge-detail-title')).toContainText(knowledgeTitle);
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText(knowledgeText);
    await expect(page.locator('.e2e-model-knowledge-detail-public')).toContainText('Yes');
    await expect(page.locator('.e2e-model-knowledge-detail-edit-button')).toBeVisible();
  });

  test('一覧からナレッジ詳細ページに遷移できる', async ({ page }) => {
    const knowledgeTitle = `一覧から詳細遷移テスト ${Date.now()}`;
    const knowledgeText = 'この内容が詳細ページで確認できるはずです。';

    // ナレッジを作成
    await createKnowledge(page, knowledgeTitle, knowledgeText, false);

    // 一覧ページに移動
    await gotoKnowledgeList(page);

    // 作成したナレッジのShowボタンをクリック
    const knowledgeItems = page.locator('.e2e-model-knowledge-list-item');
    const targetItem = knowledgeItems.filter({ hasText: knowledgeTitle });
    await targetItem.locator('.e2e-model-knowledge-list-item-show-button').click();

    // 詳細ページに遷移していることを確認
    await expect(page).toHaveURL(/\/knowledge\/\d+/);
    await expect(page.locator('.e2e-model-knowledge-detail-title')).toContainText(knowledgeTitle);
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText(knowledgeText);
  });
});

test.describe('ナレッジエッジケース・エラーハンドリング', () => {
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

  test('非常に長いタイトルでナレッジを作成できる', async ({ page }) => {
    const longTitle = `${'非常に長いタイトルのテスト'.repeat(10)} ${Date.now()}`;
    const knowledgeText = '長いタイトルのテスト内容です。';

    // 長いタイトルでナレッジを作成
    await createKnowledge(page, longTitle, knowledgeText, false);

    // 作成されたことを確認
    await expect(page.locator('.e2e-model-knowledge-detail-title')).toContainText(longTitle);
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText(knowledgeText);
  });

  test('非常に長い内容でナレッジを作成できる', async ({ page }) => {
    const knowledgeTitle = `長い内容テスト ${Date.now()}`;
    const longText = `${'これは非常に長い内容のテストです。'.repeat(50)}`;

    // 長い内容でナレッジを作成
    await createKnowledge(page, knowledgeTitle, longText, false);

    // 作成されたことを確認
    await expect(page.locator('.e2e-model-knowledge-detail-title')).toContainText(knowledgeTitle);
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText(longText);
  });

  test('特殊文字を含むタイトルでナレッジを作成できる', async ({ page }) => {
    const specialTitle = `特殊文字テスト!@#$%^&*()_+-=[]{}|;:'"<>,.?/ ${Date.now()}`;
    const knowledgeText = '特殊文字を含むタイトルのテスト内容です。';

    // 特殊文字を含むタイトルでナレッジを作成
    await createKnowledge(page, specialTitle, knowledgeText, false);

    // 作成されたことを確認
    await expect(page.locator('.e2e-model-knowledge-detail-title')).toContainText(specialTitle);
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText(knowledgeText);
  });

  test('改行を含む内容でナレッジを作成できる', async ({ page }) => {
    const knowledgeTitle = `改行テスト ${Date.now()}`;
    const multilineText = `1行目の内容です。
2行目の内容です。
3行目の内容です。

空行を含む5行目の内容です。`;

    // 改行を含む内容でナレッジを作成
    await createKnowledge(page, knowledgeTitle, multilineText, false);

    // 作成されたことを確認
    await expect(page.locator('.e2e-model-knowledge-detail-title')).toContainText(knowledgeTitle);
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText('1行目の内容です。');
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText('5行目の内容です。');
  });

  test('編集キャンセル機能が正しく動作する', async ({ page }) => {
    const originalTitle = `編集キャンセルテスト ${Date.now()}`;
    const originalText = '元の内容です。';

    // ナレッジを作成
    await createKnowledge(page, originalTitle, originalText, false);

    // URLからナレッジIDを取得
    const url = page.url();
    const knowledgeId = url.match(/\/knowledge\/(\d+)/)?.[1] || '';

    // 編集ページに移動
    await page.click('.e2e-model-knowledge-detail-edit-button');

    // 内容を変更するが保存しない
    await page.fill('.e2e-model-knowledge-edit-input-title', '変更されたタイトル');
    await page.fill('.e2e-model-knowledge-edit-input-text', '変更された内容');

    // Backボタンをクリックして詳細ページに戻る
    await page.click('.e2e-model-knowledge-edit-back-button');

    // 元の内容が保持されていることを確認
    await expect(page.locator('.e2e-model-knowledge-detail-title')).toContainText(originalTitle);
    await expect(page.locator('.e2e-model-knowledge-detail-text')).toContainText(originalText);
  });

  test('存在しないナレッジページへのアクセス処理', async ({ page }) => {
    // 存在しない可能性の高いIDでアクセス
    const nonExistentId = '999999';
    await page.goto(`/knowledge/${nonExistentId}`);

    // エラーページや適切な処理が行われることを確認
    // （実際の実装に依存するため、適切な確認方法を使用）
    await page.waitForTimeout(2000);

    // 一覧ページにリダイレクトされることを確認
    const currentUrl = page.url();
    expect(currentUrl.includes('/knowledge_list') || currentUrl.includes('404') || currentUrl.includes('error')).toBe(true);
  });
});
