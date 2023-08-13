import { graphql } from '@/gql/__generated__'
import { useGraphQL } from '@/api/main/useGraphQL'
import { userCacheKeyGenerator } from '@/usecase/user/cache'

const currentUser = graphql(/* GraphQL */ `
    query CurrentUser {
        currentUser {
            id
            name
        }
    }
`);

export const useCurrentUser = () => {
  return useGraphQL(userCacheKeyGenerator.currentUserKey(), currentUser)
}