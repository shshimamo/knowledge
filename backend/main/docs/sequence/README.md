# Backend/Main ドキュメント

## シーケンス図一覧

このディレクトリには、backend/mainアプリケーションの各ユースケースのシーケンス図が含まれています。

### UserUseCase
- `user_create.puml` - ユーザー作成フロー
- `user_get.puml` - ユーザー取得フロー

### KnowledgeUseCase
- `knowledge_create.puml` - ナレッジ作成フロー
- `knowledge_get.puml` - ナレッジ取得フロー（認可ロジック含む）
- `knowledge_update.puml` - ナレッジ更新フロー
- `knowledge_delete.puml` - ナレッジ削除フロー
- `knowledge_get_list.puml` - ナレッジ一覧取得フロー
- `knowledge_get_my.puml` - 自分のナレッジ取得フロー

### アーキテクチャ
- `overall_architecture.puml` - 全体的なアーキテクチャ構成図

## 自動更新機能

このディレクトリのPlantUMLファイルは、UseCaseファイルが更新された際にGitHub Actionsによって**自動的に更新**されます。

詳細は [GitHub Actions 自動更新ドキュメント](../../.github/README.md) を参照してください。

## PlantUMLの使用方法

これらの.pumlファイルは以下の方法で図として表示できます：

1. **PlantUML拡張機能を使用**（VS Code、IntelliJ等）
2. **オンラインエディタ**: http://www.plantuml.com/plantuml/
3. **ローカル環境でのコンパイル**: PlantUMLをインストールして`plantuml *.puml`

## アーキテクチャ概要

このアプリケーションはClean Architectureパターンに従って実装されており、以下の層で構成されています：

- **Presentation層**: GraphQLハンドラー、ミドルウェア
- **UseCase層**: ビジネスロジックの orchestration
- **Domain層**: エンティティとリポジトリインターフェース
- **Infrastructure層**: データベースアクセス、外部API連携