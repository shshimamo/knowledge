import { graphql } from '@/gql/__generated__'
import { useGraphQL } from '@/api/main/useGraphQL'

const currentUser = graphql(/* GraphQL */ `
    query CurrentUser {
        currentUser {
            id
            name
        }
    }
`);

export const useCurrentUser = () => {
  return useGraphQL(currentUser)
}