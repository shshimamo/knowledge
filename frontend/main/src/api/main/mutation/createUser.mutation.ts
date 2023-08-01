import gql from 'graphql-tag'

gql`
    mutation CreateUser($name: String!) {
        createUser(input: { name: $name }) {
            id
            authUserId
            name
        }
    }
`