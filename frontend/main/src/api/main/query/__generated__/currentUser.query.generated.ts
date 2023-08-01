import * as SchemaTypes from "../../../../../graphql/__generated__/graphql-schema-types";

import { GraphQLClient } from "graphql-request";
import { GraphQLClientRequestHeaders } from "graphql-request/build/cjs/types";
import gql from "graphql-tag";
export type CurrentUserQueryVariables = SchemaTypes.Exact<{
  [key: string]: never;
}>;

export type CurrentUserQuery = {
  __typename?: "Query";
  currentUser: { __typename?: "User"; id: string; name: string | null };
};

export const CurrentUserDocument = gql`
  query CurrentUser {
    currentUser {
      id
      name
    }
  }
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
    CurrentUser(
      variables?: CurrentUserQueryVariables,
      requestHeaders?: GraphQLClientRequestHeaders,
    ): Promise<CurrentUserQuery> {
      return withWrapper(
        (wrappedRequestHeaders) =>
          client.request<CurrentUserQuery>(CurrentUserDocument, variables, {
            ...requestHeaders,
            ...wrappedRequestHeaders,
          }),
        "CurrentUser",
        "query",
      );
    },
  };
}
export type Sdk = ReturnType<typeof getSdk>;
