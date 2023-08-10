/* eslint-disable */
import type { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type DateString = string & { readonly __brand: unique symbol }
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type CurrentUser = {
  authUserId: Scalars['ID']['output'];
  id: Scalars['ID']['output'];
  knowledge: Knowledge;
  knowledgeList: Array<Knowledge>;
  name: Maybe<Scalars['String']['output']>;
};


export type CurrentUserKnowledgeArgs = {
  id: Scalars['ID']['input'];
};


export type CurrentUserKnowledgeListArgs = {
  first: Scalars['Int']['input'];
};

export type DeleteKnowledgeResult = {
  id: Scalars['ID']['output'];
  success: Scalars['Boolean']['output'];
};

export type Knowledge = {
  id: Scalars['ID']['output'];
  isPublic: Scalars['Boolean']['output'];
  publishedAt: Scalars['String']['output'];
  text: Scalars['String']['output'];
  title: Scalars['String']['output'];
  userId: Scalars['ID']['output'];
};

export type Mutation = {
  createKnowledge: Knowledge;
  createUser: User;
  deleteKnowledge: DeleteKnowledgeResult;
  updateKnowledge: Knowledge;
};


export type MutationCreateKnowledgeArgs = {
  input: InputMaybe<NewKnowledge>;
};


export type MutationCreateUserArgs = {
  input: NewUser;
};


export type MutationDeleteKnowledgeArgs = {
  id: Scalars['ID']['input'];
};


export type MutationUpdateKnowledgeArgs = {
  id: Scalars['ID']['input'];
  input: InputMaybe<UpdateKnowledge>;
};

export type NewKnowledge = {
  isPublic: Scalars['Boolean']['input'];
  text: Scalars['String']['input'];
  title: Scalars['String']['input'];
};

export type NewUser = {
  name: Scalars['String']['input'];
};

export type Query = {
  currentUser: CurrentUser;
  knowledge: Knowledge;
  user: User;
};


export type QueryKnowledgeArgs = {
  id: Scalars['ID']['input'];
};


export type QueryUserArgs = {
  id: Scalars['ID']['input'];
};

export type UpdateKnowledge = {
  isPublic: Scalars['Boolean']['input'];
  text: Scalars['String']['input'];
  title: Scalars['String']['input'];
};

export type User = {
  authUserId: Scalars['ID']['output'];
  id: Scalars['ID']['output'];
  name: Maybe<Scalars['String']['output']>;
};

export type KnowledgeItemFragment = { title: string, text: string } & { ' $fragmentName'?: 'KnowledgeItemFragment' };

export type MyKnowledgeListQueryVariables = Exact<{ [key: string]: never; }>;


export type MyKnowledgeListQuery = { currentUser: { knowledgeList: Array<{ ' $fragmentRefs'?: { 'KnowledgeItemFragment': KnowledgeItemFragment } }> } };

export type CurrentUserQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentUserQuery = { currentUser: { id: string, name: string | null } };

export type CreateUserMutationVariables = Exact<{
  name: Scalars['String']['input'];
}>;


export type CreateUserMutation = { createUser: { id: string, authUserId: string, name: string | null } };

export type CurrentUserToSignInQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentUserToSignInQuery = { currentUser: { id: string, name: string | null } };

export const KnowledgeItemFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"KnowledgeItem"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Knowledge"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"text"}}]}}]} as unknown as DocumentNode<KnowledgeItemFragment, unknown>;
export const MyKnowledgeListDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"MyKnowledgeList"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currentUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"knowledgeList"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"first"},"value":{"kind":"IntValue","value":"20"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"KnowledgeItem"}}]}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"KnowledgeItem"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Knowledge"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"text"}}]}}]} as unknown as DocumentNode<MyKnowledgeListQuery, MyKnowledgeListQueryVariables>;
export const CurrentUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CurrentUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currentUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<CurrentUserQuery, CurrentUserQueryVariables>;
export const CreateUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"authUserId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<CreateUserMutation, CreateUserMutationVariables>;
export const CurrentUserToSignInDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CurrentUserToSignIn"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currentUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<CurrentUserToSignInQuery, CurrentUserToSignInQueryVariables>;