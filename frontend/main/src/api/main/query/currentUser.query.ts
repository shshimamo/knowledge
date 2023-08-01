import gql from 'graphql-tag'

gql`
    query CurrentUser {
        currentUser {
            id
            name
        }
    }
`;
