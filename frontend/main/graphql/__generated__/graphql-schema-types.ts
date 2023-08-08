export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
export type MakeEmpty<
  T extends { [key: string]: unknown },
  K extends keyof T,
> = { [_ in K]?: never };
export type Incremental<T> =
  | T
  | {
      [P in keyof T]?: P extends " $fragmentName" | "__typename" ? T[P] : never;
    };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string };
  String: { input: string; output: string };
  Boolean: { input: boolean; output: boolean };
  Int: { input: number; output: number };
  Float: { input: number; output: number };
};

export type CurrentUser = {
  __typename?: "CurrentUser";
  authUserId: Scalars["ID"]["output"];
  id: Scalars["ID"]["output"];
  knowledge: Knowledge;
  knowledgeList: Array<Knowledge>;
  name: Maybe<Scalars["String"]["output"]>;
};

export type CurrentUserKnowledgeArgs = {
  id: Scalars["ID"]["input"];
};

export type CurrentUserKnowledgeListArgs = {
  first: Scalars["Int"]["input"];
};

export type DeleteKnowledgeResult = {
  __typename?: "DeleteKnowledgeResult";
  id: Scalars["ID"]["output"];
  success: Scalars["Boolean"]["output"];
};

export type Knowledge = {
  __typename?: "Knowledge";
  id: Scalars["ID"]["output"];
  isPublic: Scalars["Boolean"]["output"];
  publishedAt: Scalars["String"]["output"];
  text: Scalars["String"]["output"];
  title: Scalars["String"]["output"];
  userId: Scalars["ID"]["output"];
};

export type Mutation = {
  __typename?: "Mutation";
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
  id: Scalars["ID"]["input"];
};

export type MutationUpdateKnowledgeArgs = {
  id: Scalars["ID"]["input"];
  input: InputMaybe<UpdateKnowledge>;
};

export type NewKnowledge = {
  isPublic: Scalars["Boolean"]["input"];
  text: Scalars["String"]["input"];
  title: Scalars["String"]["input"];
};

export type NewUser = {
  name: Scalars["String"]["input"];
};

export type Query = {
  __typename?: "Query";
  currentUser: CurrentUser;
  knowledge: Knowledge;
  user: User;
};

export type QueryKnowledgeArgs = {
  id: Scalars["ID"]["input"];
};

export type QueryUserArgs = {
  id: Scalars["ID"]["input"];
};

export type UpdateKnowledge = {
  isPublic: Scalars["Boolean"]["input"];
  text: Scalars["String"]["input"];
  title: Scalars["String"]["input"];
};

export type User = {
  __typename?: "User";
  authUserId: Scalars["ID"]["output"];
  id: Scalars["ID"]["output"];
  name: Maybe<Scalars["String"]["output"]>;
};
