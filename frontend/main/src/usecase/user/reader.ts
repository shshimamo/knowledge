import { graphql } from '@/gql/__generated__'
import { useGraphQL } from '@/hooks/useQuery'

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