/* eslint-disable */
import * as types from './graphql';
import type { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n    fragment KnowledgeDetail on Knowledge {\n        id\n        userId\n        title\n        text\n        isPublic\n        publishedAt\n    }\n": types.KnowledgeDetailFragmentDoc,
    "\n    fragment KnowledgeEdit on Knowledge {\n        id\n        userId\n        title\n        text\n        isPublic\n        publishedAt\n    }\n": types.KnowledgeEditFragmentDoc,
    "\n    fragment KnowledgeListItem on Knowledge {\n        id\n        title\n        text\n    }\n": types.KnowledgeListItemFragmentDoc,
    "\n    query MyKnowledgeList($first: Int!) {\n        currentUser {\n            knowledgeList(first: $first) {\n                ...KnowledgeListItem\n            }\n        }\n    }\n": types.MyKnowledgeListDocument,
    "\n    query KnowledgeItemForDetail($id: ID!) {\n        knowledge(id: $id) {\n            ...KnowledgeDetail\n        }\n    }\n": types.KnowledgeItemForDetailDocument,
    "\n    query KnowledgeItemForEdit($id: ID!) {\n        knowledge(id: $id) {\n            ...KnowledgeEdit\n        }\n    }\n": types.KnowledgeItemForEditDocument,
    "\n    mutation UpdateKnowledge($id: ID!, $input: UpdateKnowledgeInput!) {\n        updateKnowledge(id: $id, input: $input) {\n            id\n            userId\n            title\n            text\n            isPublic\n            publishedAt\n        }\n    }\n": types.UpdateKnowledgeDocument,
    "\n  mutation DeleteKnowledge($id: ID!) {\n      deleteKnowledge(id: $id) {\n          id\n          success\n      }\n  }\n": types.DeleteKnowledgeDocument,
    "\n    query CurrentUser {\n        currentUser {\n            id\n            name\n        }\n    }\n": types.CurrentUserDocument,
    "\n    mutation createUser($name: String!) {\n        createUser(input: { name: $name }) {\n            id\n            authUserId\n            name\n        }\n    }\n": types.CreateUserDocument,
    "\n    query CurrentUserToSignIn {\n        currentUser {\n            id\n            name\n        }\n    }\n": types.CurrentUserToSignInDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    fragment KnowledgeDetail on Knowledge {\n        id\n        userId\n        title\n        text\n        isPublic\n        publishedAt\n    }\n"): (typeof documents)["\n    fragment KnowledgeDetail on Knowledge {\n        id\n        userId\n        title\n        text\n        isPublic\n        publishedAt\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    fragment KnowledgeEdit on Knowledge {\n        id\n        userId\n        title\n        text\n        isPublic\n        publishedAt\n    }\n"): (typeof documents)["\n    fragment KnowledgeEdit on Knowledge {\n        id\n        userId\n        title\n        text\n        isPublic\n        publishedAt\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    fragment KnowledgeListItem on Knowledge {\n        id\n        title\n        text\n    }\n"): (typeof documents)["\n    fragment KnowledgeListItem on Knowledge {\n        id\n        title\n        text\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query MyKnowledgeList($first: Int!) {\n        currentUser {\n            knowledgeList(first: $first) {\n                ...KnowledgeListItem\n            }\n        }\n    }\n"): (typeof documents)["\n    query MyKnowledgeList($first: Int!) {\n        currentUser {\n            knowledgeList(first: $first) {\n                ...KnowledgeListItem\n            }\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query KnowledgeItemForDetail($id: ID!) {\n        knowledge(id: $id) {\n            ...KnowledgeDetail\n        }\n    }\n"): (typeof documents)["\n    query KnowledgeItemForDetail($id: ID!) {\n        knowledge(id: $id) {\n            ...KnowledgeDetail\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query KnowledgeItemForEdit($id: ID!) {\n        knowledge(id: $id) {\n            ...KnowledgeEdit\n        }\n    }\n"): (typeof documents)["\n    query KnowledgeItemForEdit($id: ID!) {\n        knowledge(id: $id) {\n            ...KnowledgeEdit\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation UpdateKnowledge($id: ID!, $input: UpdateKnowledgeInput!) {\n        updateKnowledge(id: $id, input: $input) {\n            id\n            userId\n            title\n            text\n            isPublic\n            publishedAt\n        }\n    }\n"): (typeof documents)["\n    mutation UpdateKnowledge($id: ID!, $input: UpdateKnowledgeInput!) {\n        updateKnowledge(id: $id, input: $input) {\n            id\n            userId\n            title\n            text\n            isPublic\n            publishedAt\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation DeleteKnowledge($id: ID!) {\n      deleteKnowledge(id: $id) {\n          id\n          success\n      }\n  }\n"): (typeof documents)["\n  mutation DeleteKnowledge($id: ID!) {\n      deleteKnowledge(id: $id) {\n          id\n          success\n      }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query CurrentUser {\n        currentUser {\n            id\n            name\n        }\n    }\n"): (typeof documents)["\n    query CurrentUser {\n        currentUser {\n            id\n            name\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createUser($name: String!) {\n        createUser(input: { name: $name }) {\n            id\n            authUserId\n            name\n        }\n    }\n"): (typeof documents)["\n    mutation createUser($name: String!) {\n        createUser(input: { name: $name }) {\n            id\n            authUserId\n            name\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query CurrentUserToSignIn {\n        currentUser {\n            id\n            name\n        }\n    }\n"): (typeof documents)["\n    query CurrentUserToSignIn {\n        currentUser {\n            id\n            name\n        }\n    }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;