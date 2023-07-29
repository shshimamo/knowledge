import { gql } from '@apollo/client';
import { useCreateUserMutation } from '@/api/main/mutation/__generated__/createUser.mutation'

gql`
    mutation CreateUser($name: String!) {
        createUser(input: { name: $name }) {
            id
            authUserId
            name
        }
    }
`

type UseCreateUserMutationReturn = ReturnType<typeof useCreateUserMutation>;
export type CreateUserMutation = UseCreateUserMutationReturn[0];