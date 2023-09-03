import { TypedDocumentNode } from '@graphql-typed-document-node/core'
import { Variables } from 'graphql-request'
import useSWR from 'swr'

import { gqlClient } from '@/api/main/gqlClient'

export function useGraphQL<TResult, TVariables extends Variables | undefined>(
  cacheKey: readonly string[],
  document: TypedDocumentNode<TResult, TVariables>,
  ...[variables]: TVariables extends Record<string, never> ? [] : [TVariables]
) {
  return useSWR<TResult>(
    cacheKey,
    async () =>
      (await gqlClient.request({
        document: document,
        variables: variables,
      })) as Promise<TResult>,
    { revalidateOnFocus: false },
  )
}
