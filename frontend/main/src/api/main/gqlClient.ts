import { GraphQLClient } from 'graphql-request'

const endpoint = process.env.NEXT_PUBLIC_MAIN_GQL_URL || (process.env.NODE_ENV === 'production' ? 'http://backend.main.shshimamo.com/query' : 'http://localhost:8080/query')
export const gqlClient = new GraphQLClient(endpoint, {
  credentials: 'include',
});
