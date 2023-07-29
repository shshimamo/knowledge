import { ApolloClient, createHttpLink, InMemoryCache } from "@apollo/client";
import { setContext } from '@apollo/client/link/context';

const httpLink = createHttpLink({
  uri: process.env.NODE_ENV === 'production' ? 'Production URL' : 'http://localhost:8080/query',
  credentials: 'include',
});

const requestLink = setContext((_, { headers }) => {
  return {
    headers: {
      ...headers,
    }
  };
});

const client = new ApolloClient({
  link: requestLink.concat(httpLink),
  cache: new InMemoryCache(),
});

export default client;
