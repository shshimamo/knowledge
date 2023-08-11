import { TypedDocumentNode } from '@graphql-typed-document-node/core';
import { ASTNode, Kind, OperationDefinitionNode } from 'graphql';
import useSWR from 'swr';
import { gqlClient } from '@/api/main/gqlClient'

const isOperationDefinition = (def: ASTNode): def is OperationDefinitionNode => def.kind === Kind.OPERATION_DEFINITION;

export function useGraphQL<TResult, TVariables>(
  document: TypedDocumentNode<TResult, TVariables>,
  ...[variables]: TVariables extends Record<string, never> ? [] : [TVariables]
) {
  return useSWR(
    [
      document.definitions.find(isOperationDefinition)?.name?.value,
      variables,
    ],
    async ([_key, variables]) =>
      await gqlClient.request({
        document: document,
        variables: variables as any,
      }) as Promise<TResult>
  );
}