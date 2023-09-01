import { useGraphQL } from '@/api/main/useGraphQL'
import { graphql } from '@/gql/__generated__'
import { knowledgeCacheKeyGenerator } from '@/usecase/knowledge/cache'

// prettier-ignore
export const myKnowledgeList = graphql(/* GraphQL */ `
    query MyKnowledgeList($first: Int!) {
        currentUser {
            knowledgeList(first: $first) {
                ...KnowledgeListItem
            }
        }
    }
`)

export const useMyKnowledgeList = ({ first }: { first: number }) => {
  return useGraphQL(
    knowledgeCacheKeyGenerator.myKnowledgeListKey({ variables: { first } }),
    myKnowledgeList,
    { first },
  )
}

// prettier-ignore
const knowledgeItemForDetail = graphql(/* GraphQL */ `
    query KnowledgeItemForDetail($id: ID!) {
        knowledge(id: $id) {
            ...KnowledgeDetail
        }
    }
`)

export const useKnowledgeItemForDetail = ({ id }: { id: string }) => {
  return useGraphQL(
    knowledgeCacheKeyGenerator.knowledgeItemForDetailKey({ id }),
    knowledgeItemForDetail,
    { id },
  )
}

// prettier-ignore
const knowledgeItemForEdit = graphql(/* GraphQL */ `
    query KnowledgeItemForEdit($id: ID!) {
        knowledge(id: $id) {
            ...KnowledgeEdit
        }
    }
`)

export const useKnowledgeItemForEdit = ({ id }: { id: string }) => {
  return useGraphQL(
    knowledgeCacheKeyGenerator.knowledgeItemForEditKey({ id }),
    knowledgeItemForEdit,
    { id },
  )
}
