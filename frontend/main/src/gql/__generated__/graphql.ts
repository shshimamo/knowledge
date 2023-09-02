/* eslint-disable */
import type { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core'
export type Maybe<T> = T | null
export type InputMaybe<T> = Maybe<T>
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] }
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>
}
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>
}
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = {
  [_ in K]?: never
}
export type Incremental<T> =
  | T
  | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never }
export type DateString = string & { readonly __brand: unique symbol }
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string }
  String: { input: string; output: string }
  Boolean: { input: boolean; output: boolean }
  Int: { input: number; output: number }
  Float: { input: number; output: number }
}

export type CreateKnowledgeInput = {
  isPublic: Scalars['Boolean']['input']
  text: Scalars['String']['input']
  title: Scalars['String']['input']
}

export type CurrentUser = {
  authUserId: Scalars['ID']['output']
  id: Scalars['ID']['output']
  knowledge: Knowledge
  knowledgeList: Array<Knowledge>
  name: Maybe<Scalars['String']['output']>
}

export type CurrentUserKnowledgeArgs = {
  id: Scalars['ID']['input']
}

export type CurrentUserKnowledgeListArgs = {
  first: Scalars['Int']['input']
}

export type DeleteKnowledgeResult = {
  id: Scalars['ID']['output']
  success: Scalars['Boolean']['output']
}

export type Knowledge = {
  id: Scalars['ID']['output']
  isPublic: Scalars['Boolean']['output']
  publishedAt: Scalars['String']['output']
  text: Scalars['String']['output']
  title: Scalars['String']['output']
  userId: Scalars['ID']['output']
}

export type Mutation = {
  createKnowledge: Knowledge
  createUser: User
  deleteKnowledge: DeleteKnowledgeResult
  updateKnowledge: Knowledge
}

export type MutationCreateKnowledgeArgs = {
  input: InputMaybe<CreateKnowledgeInput>
}

export type MutationCreateUserArgs = {
  input: NewUser
}

export type MutationDeleteKnowledgeArgs = {
  id: Scalars['ID']['input']
}

export type MutationUpdateKnowledgeArgs = {
  id: Scalars['ID']['input']
  input: InputMaybe<UpdateKnowledgeInput>
}

export type NewUser = {
  name: Scalars['String']['input']
}

export type Query = {
  currentUser: CurrentUser
  knowledge: Knowledge
  user: User
}

export type QueryKnowledgeArgs = {
  id: Scalars['ID']['input']
}

export type QueryUserArgs = {
  id: Scalars['ID']['input']
}

export type UpdateKnowledgeInput = {
  isPublic: Scalars['Boolean']['input']
  text: Scalars['String']['input']
  title: Scalars['String']['input']
}

export type User = {
  authUserId: Scalars['ID']['output']
  id: Scalars['ID']['output']
  name: Maybe<Scalars['String']['output']>
}

export type KnowledgeDetailFragment = {
  id: string
  userId: string
  title: string
  text: string
  isPublic: boolean
  publishedAt: string
} & { ' $fragmentName'?: 'KnowledgeDetailFragment' }

export type KnowledgeEditFragment = {
  id: string
  userId: string
  title: string
  text: string
  isPublic: boolean
  publishedAt: string
} & { ' $fragmentName'?: 'KnowledgeEditFragment' }

export type KnowledgeListItemFragment = { id: string; title: string; text: string } & {
  ' $fragmentName'?: 'KnowledgeListItemFragment'
}

export type MyKnowledgeListQueryVariables = Exact<{
  first: Scalars['Int']['input']
}>

export type MyKnowledgeListQuery = {
  currentUser: {
    knowledgeList: Array<{
      ' $fragmentRefs'?: { KnowledgeListItemFragment: KnowledgeListItemFragment }
    }>
  }
}

export type KnowledgeItemForDetailQueryVariables = Exact<{
  id: Scalars['ID']['input']
}>

export type KnowledgeItemForDetailQuery = {
  knowledge: { ' $fragmentRefs'?: { KnowledgeDetailFragment: KnowledgeDetailFragment } }
}

export type KnowledgeItemForEditQueryVariables = Exact<{
  id: Scalars['ID']['input']
}>

export type KnowledgeItemForEditQuery = {
  knowledge: { ' $fragmentRefs'?: { KnowledgeEditFragment: KnowledgeEditFragment } }
}

export type CreateKnowledgeMutationVariables = Exact<{
  input: CreateKnowledgeInput
}>

export type CreateKnowledgeMutation = { createKnowledge: { id: string } }

export type UpdateKnowledgeMutationVariables = Exact<{
  id: Scalars['ID']['input']
  input: UpdateKnowledgeInput
}>

export type UpdateKnowledgeMutation = {
  updateKnowledge: {
    id: string
    userId: string
    title: string
    text: string
    isPublic: boolean
    publishedAt: string
  }
}

export type DeleteKnowledgeMutationVariables = Exact<{
  id: Scalars['ID']['input']
}>

export type DeleteKnowledgeMutation = {
  deleteKnowledge: { id: string; success: boolean }
}

export type CurrentUserQueryVariables = Exact<{ [key: string]: never }>

export type CurrentUserQuery = { currentUser: { id: string; name: string | null } }

export type CreateUserMutationVariables = Exact<{
  name: Scalars['String']['input']
}>

export type CreateUserMutation = {
  createUser: { id: string; authUserId: string; name: string | null }
}

export type CurrentUserToSignInQueryVariables = Exact<{ [key: string]: never }>

export type CurrentUserToSignInQuery = {
  currentUser: { id: string; name: string | null }
}

export const KnowledgeDetailFragmentDoc = {
  kind: 'Document',
  definitions: [
    {
      kind: 'FragmentDefinition',
      name: { kind: 'Name', value: 'KnowledgeDetail' },
      typeCondition: { kind: 'NamedType', name: { kind: 'Name', value: 'Knowledge' } },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          { kind: 'Field', name: { kind: 'Name', value: 'id' } },
          { kind: 'Field', name: { kind: 'Name', value: 'userId' } },
          { kind: 'Field', name: { kind: 'Name', value: 'title' } },
          { kind: 'Field', name: { kind: 'Name', value: 'text' } },
          { kind: 'Field', name: { kind: 'Name', value: 'isPublic' } },
          { kind: 'Field', name: { kind: 'Name', value: 'publishedAt' } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<KnowledgeDetailFragment, unknown>
export const KnowledgeEditFragmentDoc = {
  kind: 'Document',
  definitions: [
    {
      kind: 'FragmentDefinition',
      name: { kind: 'Name', value: 'KnowledgeEdit' },
      typeCondition: { kind: 'NamedType', name: { kind: 'Name', value: 'Knowledge' } },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          { kind: 'Field', name: { kind: 'Name', value: 'id' } },
          { kind: 'Field', name: { kind: 'Name', value: 'userId' } },
          { kind: 'Field', name: { kind: 'Name', value: 'title' } },
          { kind: 'Field', name: { kind: 'Name', value: 'text' } },
          { kind: 'Field', name: { kind: 'Name', value: 'isPublic' } },
          { kind: 'Field', name: { kind: 'Name', value: 'publishedAt' } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<KnowledgeEditFragment, unknown>
export const KnowledgeListItemFragmentDoc = {
  kind: 'Document',
  definitions: [
    {
      kind: 'FragmentDefinition',
      name: { kind: 'Name', value: 'KnowledgeListItem' },
      typeCondition: { kind: 'NamedType', name: { kind: 'Name', value: 'Knowledge' } },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          { kind: 'Field', name: { kind: 'Name', value: 'id' } },
          { kind: 'Field', name: { kind: 'Name', value: 'title' } },
          { kind: 'Field', name: { kind: 'Name', value: 'text' } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<KnowledgeListItemFragment, unknown>
export const MyKnowledgeListDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'query',
      name: { kind: 'Name', value: 'MyKnowledgeList' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'first' } },
          type: {
            kind: 'NonNullType',
            type: { kind: 'NamedType', name: { kind: 'Name', value: 'Int' } },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'currentUser' },
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                {
                  kind: 'Field',
                  name: { kind: 'Name', value: 'knowledgeList' },
                  arguments: [
                    {
                      kind: 'Argument',
                      name: { kind: 'Name', value: 'first' },
                      value: { kind: 'Variable', name: { kind: 'Name', value: 'first' } },
                    },
                  ],
                  selectionSet: {
                    kind: 'SelectionSet',
                    selections: [
                      {
                        kind: 'FragmentSpread',
                        name: { kind: 'Name', value: 'KnowledgeListItem' },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: 'FragmentDefinition',
      name: { kind: 'Name', value: 'KnowledgeListItem' },
      typeCondition: { kind: 'NamedType', name: { kind: 'Name', value: 'Knowledge' } },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          { kind: 'Field', name: { kind: 'Name', value: 'id' } },
          { kind: 'Field', name: { kind: 'Name', value: 'title' } },
          { kind: 'Field', name: { kind: 'Name', value: 'text' } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<MyKnowledgeListQuery, MyKnowledgeListQueryVariables>
export const KnowledgeItemForDetailDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'query',
      name: { kind: 'Name', value: 'KnowledgeItemForDetail' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
          type: {
            kind: 'NonNullType',
            type: { kind: 'NamedType', name: { kind: 'Name', value: 'ID' } },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'knowledge' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'id' },
                value: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                {
                  kind: 'FragmentSpread',
                  name: { kind: 'Name', value: 'KnowledgeDetail' },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: 'FragmentDefinition',
      name: { kind: 'Name', value: 'KnowledgeDetail' },
      typeCondition: { kind: 'NamedType', name: { kind: 'Name', value: 'Knowledge' } },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          { kind: 'Field', name: { kind: 'Name', value: 'id' } },
          { kind: 'Field', name: { kind: 'Name', value: 'userId' } },
          { kind: 'Field', name: { kind: 'Name', value: 'title' } },
          { kind: 'Field', name: { kind: 'Name', value: 'text' } },
          { kind: 'Field', name: { kind: 'Name', value: 'isPublic' } },
          { kind: 'Field', name: { kind: 'Name', value: 'publishedAt' } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  KnowledgeItemForDetailQuery,
  KnowledgeItemForDetailQueryVariables
>
export const KnowledgeItemForEditDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'query',
      name: { kind: 'Name', value: 'KnowledgeItemForEdit' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
          type: {
            kind: 'NonNullType',
            type: { kind: 'NamedType', name: { kind: 'Name', value: 'ID' } },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'knowledge' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'id' },
                value: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                {
                  kind: 'FragmentSpread',
                  name: { kind: 'Name', value: 'KnowledgeEdit' },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: 'FragmentDefinition',
      name: { kind: 'Name', value: 'KnowledgeEdit' },
      typeCondition: { kind: 'NamedType', name: { kind: 'Name', value: 'Knowledge' } },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          { kind: 'Field', name: { kind: 'Name', value: 'id' } },
          { kind: 'Field', name: { kind: 'Name', value: 'userId' } },
          { kind: 'Field', name: { kind: 'Name', value: 'title' } },
          { kind: 'Field', name: { kind: 'Name', value: 'text' } },
          { kind: 'Field', name: { kind: 'Name', value: 'isPublic' } },
          { kind: 'Field', name: { kind: 'Name', value: 'publishedAt' } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  KnowledgeItemForEditQuery,
  KnowledgeItemForEditQueryVariables
>
export const CreateKnowledgeDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'mutation',
      name: { kind: 'Name', value: 'CreateKnowledge' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'input' } },
          type: {
            kind: 'NonNullType',
            type: {
              kind: 'NamedType',
              name: { kind: 'Name', value: 'CreateKnowledgeInput' },
            },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'createKnowledge' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'input' },
                value: { kind: 'Variable', name: { kind: 'Name', value: 'input' } },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [{ kind: 'Field', name: { kind: 'Name', value: 'id' } }],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<CreateKnowledgeMutation, CreateKnowledgeMutationVariables>
export const UpdateKnowledgeDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'mutation',
      name: { kind: 'Name', value: 'UpdateKnowledge' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
          type: {
            kind: 'NonNullType',
            type: { kind: 'NamedType', name: { kind: 'Name', value: 'ID' } },
          },
        },
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'input' } },
          type: {
            kind: 'NonNullType',
            type: {
              kind: 'NamedType',
              name: { kind: 'Name', value: 'UpdateKnowledgeInput' },
            },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'updateKnowledge' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'id' },
                value: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
              },
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'input' },
                value: { kind: 'Variable', name: { kind: 'Name', value: 'input' } },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                { kind: 'Field', name: { kind: 'Name', value: 'id' } },
                { kind: 'Field', name: { kind: 'Name', value: 'userId' } },
                { kind: 'Field', name: { kind: 'Name', value: 'title' } },
                { kind: 'Field', name: { kind: 'Name', value: 'text' } },
                { kind: 'Field', name: { kind: 'Name', value: 'isPublic' } },
                { kind: 'Field', name: { kind: 'Name', value: 'publishedAt' } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<UpdateKnowledgeMutation, UpdateKnowledgeMutationVariables>
export const DeleteKnowledgeDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'mutation',
      name: { kind: 'Name', value: 'DeleteKnowledge' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
          type: {
            kind: 'NonNullType',
            type: { kind: 'NamedType', name: { kind: 'Name', value: 'ID' } },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'deleteKnowledge' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'id' },
                value: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                { kind: 'Field', name: { kind: 'Name', value: 'id' } },
                { kind: 'Field', name: { kind: 'Name', value: 'success' } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<DeleteKnowledgeMutation, DeleteKnowledgeMutationVariables>
export const CurrentUserDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'query',
      name: { kind: 'Name', value: 'CurrentUser' },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'currentUser' },
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                { kind: 'Field', name: { kind: 'Name', value: 'id' } },
                { kind: 'Field', name: { kind: 'Name', value: 'name' } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<CurrentUserQuery, CurrentUserQueryVariables>
export const CreateUserDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'mutation',
      name: { kind: 'Name', value: 'createUser' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'name' } },
          type: {
            kind: 'NonNullType',
            type: { kind: 'NamedType', name: { kind: 'Name', value: 'String' } },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'createUser' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'input' },
                value: {
                  kind: 'ObjectValue',
                  fields: [
                    {
                      kind: 'ObjectField',
                      name: { kind: 'Name', value: 'name' },
                      value: { kind: 'Variable', name: { kind: 'Name', value: 'name' } },
                    },
                  ],
                },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                { kind: 'Field', name: { kind: 'Name', value: 'id' } },
                { kind: 'Field', name: { kind: 'Name', value: 'authUserId' } },
                { kind: 'Field', name: { kind: 'Name', value: 'name' } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<CreateUserMutation, CreateUserMutationVariables>
export const CurrentUserToSignInDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'query',
      name: { kind: 'Name', value: 'CurrentUserToSignIn' },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'currentUser' },
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                { kind: 'Field', name: { kind: 'Name', value: 'id' } },
                { kind: 'Field', name: { kind: 'Name', value: 'name' } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<CurrentUserToSignInQuery, CurrentUserToSignInQueryVariables>
