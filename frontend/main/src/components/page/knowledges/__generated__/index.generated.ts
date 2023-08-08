import * as SchemaTypes from "../../../../../graphql/__generated__/graphql-schema-types";

import { GraphQLClient } from "graphql-request";
import { GraphQLClientRequestHeaders } from "graphql-request/build/cjs/types";
import gql from "graphql-tag";
export type MyKnowledgeListQueryVariables = SchemaTypes.Exact<{
  [key: string]: never;
}>;

export type MyKnowledgeListQuery = {
  __typename?: "Query";
  currentUser: {
    __typename?: "CurrentUser";
    knowledgeList: Array<{ __typename?: "Knowledge" }>;
  };
};

export const MyKnowledgeListDocument = gql`
  query MyKnowledgeList {
    currentUser {
      knowledgeList(first: 20) {
        ...KnowledgeItem
      }
    }
  }
  ${KnowledgeItemFragmentDoc}
`;

export type SdkFunctionWrapper = <T>(
  action: (requestHeaders?: Record<string, string>) => Promise<T>,
  operationName: string,
  operationType?: string,
) => Promise<T>;

const defaultWrapper: SdkFunctionWrapper = (
  action,
  _operationName,
  _operationType,
) => action();

export function getSdk(
  client: GraphQLClient,
  withWrapper: SdkFunctionWrapper = defaultWrapper,
) {
  return {
    MyKnowledgeList(
      variables?: MyKnowledgeListQueryVariables,
      requestHeaders?: GraphQLClientRequestHeaders,
    ): Promise<MyKnowledgeListQuery> {
      return withWrapper(
        (wrappedRequestHeaders) =>
          client.request<MyKnowledgeListQuery>(
            MyKnowledgeListDocument,
            variables,
            { ...requestHeaders, ...wrappedRequestHeaders },
          ),
        "MyKnowledgeList",
        "query",
      );
    },
  };
}
export type Sdk = ReturnType<typeof getSdk>;
