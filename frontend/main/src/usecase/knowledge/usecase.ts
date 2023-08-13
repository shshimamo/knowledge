import React from 'react'
import { graphql } from '@/gql/__generated__'
import { UpdateKnowledgeInput } from '@/gql/__generated__/graphql'
import { gqlClient } from '@/api/main/gqlClient'
import { useKnowledgeCacheMutator } from '@/usecase/knowledge/cache'

const updateKnowledge = graphql(/* GraphQL */ `
    mutation UpdateKnowledge($id: ID!, $input: UpdateKnowledgeInput!) {
        updateKnowledge(id: $id, input: $input) {
            id
            userId
            title
            text
            isPublic
            publishedAt
        }
    }
`);

const deleteKnowledge = graphql(/* GraphQL */ `
  mutation DeleteKnowledge($id: ID!) {
      deleteKnowledge(id: $id) {
          id
          success
      }
  }
`);

export const useKnowledgeUsecase = () => {
  const mutator = useKnowledgeCacheMutator();

  return React.useMemo(
    () => createKnowledgeUsecase({ mutator }),
    [mutator]
  )
}

const createKnowledgeUsecase = ({ mutator }: {
  mutator: ReturnType<typeof useKnowledgeCacheMutator>
}) => ({
  async updateKnowledge({ id, input }: { id: string, input: UpdateKnowledgeInput }) {
    try {
      await gqlClient.request(updateKnowledge, { id, input })
      await mutator.mutateAllKnowledgeList()
      await mutator.mutateAllKnowledgeItem({ id })
    } catch (error) {
      const errorMessage = 'unexpected error'
      throw new Error(errorMessage)
    }
  },

  async deleteKnowledge(id: string) {
    try {
      await gqlClient.request(deleteKnowledge, { id })
      await mutator.mutateAllKnowledgeList()
      await mutator.mutateAllKnowledgeItem({ id })
    } catch (error) {
      const errorMessage = 'unexpected error'
      throw new Error(errorMessage)
    }
  },
})