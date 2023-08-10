/* eslint-disable import/no-extraneous-dependencies */
import { TypedDocumentNode } from '@graphql-typed-document-node/core';
import { ASTNode, ExecutionResult, Kind, OperationDefinitionNode } from 'graphql';
import useSWR from 'swr';
import { gqlClient } from '@/api/main/gqlClient'

const isOperationDefinition = (def: ASTNode): def is OperationDefinitionNode => def.kind === Kind.OPERATION_DEFINITION;

export function useGraphQL<TResult, TVariables>(
  document: TypedDocumentNode<TResult, TVariables>,
  ...[variables]: TVariables extends Record<string, never> ? [] : [TVariables]
) {
  return useSWR(
    [
      // This logic can be customized as desired
      document.definitions.find(isOperationDefinition)?.name,
      variables,
    ] as const,
    async (_key: any, variables: any) =>
      await gqlClient.request({
        document: document,
        variables: variables,
      }) as Promise<ExecutionResult<TResult>>
  );
}