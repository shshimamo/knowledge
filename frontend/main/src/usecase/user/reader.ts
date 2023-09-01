import { useGraphQL } from '@/api/main/useGraphQL'
import { graphql } from '@/gql/__generated__'
import { userCacheKeyGenerator } from '@/usecase/user/cache'

const currentUser = graphql(/* GraphQL */ `
  query CurrentUser {
    currentUser {
      id
      name
    }
  }
`)

export const useCurrentUser = () => {
  return useGraphQL(userCacheKeyGenerator.currentUserKey(), currentUser)
}
