import { graphql } from '@/gql/__generated__'
import { useGraphQL } from '@/hooks/useQuery'
import { gqlClient } from '@/api/main/gqlClient'
import { ExecutionResult } from 'graphql/index'
import useSWR from 'swr'
import React from 'react'
import { createAuthUsecase } from '@/usecase/user/usecase'

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