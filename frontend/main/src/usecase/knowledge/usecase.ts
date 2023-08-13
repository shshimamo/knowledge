import React from 'react'
import { graphql } from '@/gql/__generated__'
import { UpdateKnowledgeInput } from '@/gql/__generated__/graphql'
import { gqlClient } from '@/api/main/gqlClient'

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
  return React.useMemo(
    () => createKnowledgeUsecase(),
    []
  )
}

const createKnowledgeUsecase = () => ({
  async updateKnowledge({ id, input }: { id: string, input: UpdateKnowledgeInput }) {
    try {
      const data = await gqlClient.request(updateKnowledge, { id, input })
      // TODO: mutate SWR
    } catch (error) {
      const errorMessage = 'unexpected error'
      throw new Error(errorMessage)
    }
  },

  async deleteKnowledge(id: string) {
    try {
      const data = await gqlClient.request(deleteKnowledge, { id })
      // TODO: mutate SWR
    } catch (error) {
      const errorMessage = 'unexpected error'
      throw new Error(errorMessage)
    }
  },
})