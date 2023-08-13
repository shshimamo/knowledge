import React from 'react'
import { useSWRConfig } from "swr"
import { MyKnowledgeListQueryVariables } from '@/gql/__generated__/graphql'

export const knowledgeCacheKeyGenerator = {
  myKnowledgeListKey: ({ variables, includeAll }: { variables: MyKnowledgeListQueryVariables, includeAll?: boolean }) => {
    return ['KNOWLEDGE', 'LIST', 'MyKnowledgeList', ...(includeAll ? [] : [variables])] as const
  },
  knowledgeItemForDetailKey: ({ id }: { id: string })  => {
    return ['KNOWLEDGE', 'ITEM', 'KnowledgeItemForDetail', id] as const
  },
  knowledgeItemForEditKey: ({ id }: { id: string })  => {
    return ['KNOWLEDGE', 'ITEM', 'KnowledgeItemForEdit', id] as const
  },
}

export const useKnowledgeCacheMutator = () => {
  const { mutate } = useSWRConfig()

  return React.useMemo(
    () => ({
      mutateAllKnowledgeList: () => {
        return mutate(
          (key: string[]) => key[0] == 'KNOWLEDGE' && key[1] == 'LIST',
          undefined,
          { revalidate: true }
        )
      },
      mutateAllKnowledgeItem: ({ id }: { id: string }) => {
        return mutate(
          (key: string[]) => key[0] == 'KNOWLEDGE' && key[1] == 'ITEM' && key[key.length - 1] == id,
          undefined,
          { revalidate: true }
        )
      },
    }),
    [mutate],
  )
}