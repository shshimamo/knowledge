import { graphql } from '@/gql/__generated__'
import { useGraphQL } from '@/api/main/useGraphQL'

const myKnowledgeList = graphql(/* GraphQL */ `
    query MyKnowledgeList($first: Int!) {
        currentUser {
            knowledgeList(first: $first) {
                ...KnowledgeItem
            }
        }
    }
`)

export const useMyKnowledgeList = ({ first }: { first: number }) => {
    return useGraphQL(myKnowledgeList, { first })
}

const knowledgeDetail = graphql(/* GraphQL */ `
    query KnowledgeDetail($id: ID!) {
        knowledge(id: $id) {
            ...KnowledgeDetail
        }
    }
`)

export const useKnowledgeDetail = ({ id }: { id: string }) => {
    return useGraphQL(knowledgeDetail, { id })
}

const getKnowledgeForEdit = graphql(/* GraphQL */ `
    query GetKnowledgeForEdit($id: ID!) {
        knowledge(id: $id) {
            ...KnowledgeEdit
        }
    }
`)

export const useGetKnowledgeForEdit = ({ id }: { id: string }) => {
    return useGraphQL(getKnowledgeForEdit, { id })
}
