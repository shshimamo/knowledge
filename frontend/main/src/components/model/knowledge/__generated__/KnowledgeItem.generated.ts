import * as SchemaTypes from "../../../../../graphql/__generated__/graphql-schema-types";

import { GraphQLClient } from "graphql-request";
import { GraphQLClientRequestHeaders } from "graphql-request/build/cjs/types";
import gql from "graphql-tag";
export type KnowledgeItemFragment = {
  __typename?: "Knowledge";
  title: string;
  text: string;
};

export const KnowledgeItemFragmentDoc = gql`
  fragment KnowledgeItem on Knowledge {
    title
    text
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
  return {};
}
export type Sdk = ReturnType<typeof getSdk>;
