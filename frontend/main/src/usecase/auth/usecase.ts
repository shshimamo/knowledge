import React from 'react'
import { useAuthRepository, AuthRepository } from '@/repository/auth/repository';
import { SigninSeed, SignupSeed, Token } from '@/components/model/auth/type'
import { CreateUserMutation } from '@/api/main/mutation/createUser.mutation'
import { useCurrentUserMutators, SetCurrentUser } from '@/globalStates/currentUserState'
import { TokenRepository, useTokenRepository } from '@/repository/token/repository'
import {
  CurrentUserLazyQueryHookResult,
  CurrentUserQueryHookResult
} from '@/components/page/signin/__generated__/index.generated'

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
    [authRepository, setCurrentUser],
  )
}

export const createAuthUsecase = ({ authRepository, tokenRepository, setCurrentUser, createUserMutation }: {
  authRepository: AuthRepository
  tokenRepository: TokenRepository
  setCurrentUser: SetCurrentUser
}) => ({
  async signup(seed: SignupSeed, name: string, createUserMutation: CreateUserMutation) {
    try {
      // get token
      const token: Token = await authRepository.signup(seed)
      // set cookie
      await tokenRepository.setToken(token)
      // create user
      const { data } = await createUserMutation({ variables: { name } })
      if (!data?.createUser) {
        // TODO: only auth_users created and not users
        throw new Error('unexpected error')
      }

      setCurrentUser(data.createUser.id, data.createUser.name)
    } catch (error) {
      const errorMessage = 'unexpected error'
      throw new Error(errorMessage)
    }
  },

  async signin(seed: SigninSeed, currentUser: CurrentUserLazyQueryHookResult[0]) {
    try {
      // get token
      const token: Token = await authRepository.signin(seed)
      // set cookie
      await tokenRepository.setToken(token)
      // get current user
      const { data } = await currentUser()
      if (!data?.currentUser) {
        // TODO: only auth_users created and not users
        throw new Error('unexpected error')
      }

      setCurrentUser(data.currentUser.id, data.currentUser.name)
    } catch (error) {
      const errorMessage = 'unexpected error'
      throw new Error(errorMessage)
    }
  }
})
