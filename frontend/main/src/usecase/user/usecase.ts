import React from 'react'
import { useAuthRepository, AuthRepositoryType } from '@/repository/auth/repository';
import { SigninSeed, SignupSeed, Token } from '@/components/model/auth/type'
import { useCurrentUserMutators, SetCurrentUserType } from '@/globalStates/currentUserState'
import { TokenRepositoryType, useTokenRepository } from '@/repository/token/repository'
import { graphql } from '@/gql/__generated__'
import { gqlClient } from '@/api/main/gqlClient'

const createUser = graphql(/* GraphQL */`
    mutation createUser($name: String!) {
        createUser(input: { name: $name }) {
            id
            authUserId
            name
        }
    }
`);

const currentUserToSignIn = graphql(/* GraphQL */ `
    query CurrentUserToSignIn {
        currentUser {
            id
            name
        }
    }
`);

export const useAuthUsecase = () => {
  const authRepository = useAuthRepository()
  const tokenRepository = useTokenRepository()
  const { setCurrentUser } = useCurrentUserMutators();

  return React.useMemo(
    () => createAuthUsecase({
      authRepository,
      tokenRepository,
      setCurrentUser,
    }),
    [authRepository, tokenRepository, setCurrentUser],
  )
}

export const createAuthUsecase = ({ authRepository, tokenRepository, setCurrentUser }: {
  authRepository: AuthRepositoryType
  tokenRepository: TokenRepositoryType
  setCurrentUser: SetCurrentUserType
}) => ({
  async signup(seed: SignupSeed, name: string) {
    try {
      // get token
      const token: Token = await authRepository.signup(seed)
      // set cookie to main
      await tokenRepository.setToken(token)
      // create user
      const data = await gqlClient.request(createUser, { name })
      // set state
      setCurrentUser(data.createUser.id, data.createUser.name)
    } catch (error) {
      const errorMessage = 'unexpected error'
      throw new Error(errorMessage)
    }
  },

  async signin(seed: SigninSeed) {
    try {
      // get token
      const token: Token = await authRepository.signin(seed)
      // set cookie to main
      await tokenRepository.setToken(token)
      // get current user
      const data = await gqlClient.request(currentUserToSignIn)
      // set state
      setCurrentUser(data.currentUser.id, data.currentUser.name)
    } catch (error) {
      const errorMessage = 'unexpected error'
      throw new Error(errorMessage)
    }
  }
})
