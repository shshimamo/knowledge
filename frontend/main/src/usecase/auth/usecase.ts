import React from 'react'
import { useAuthRepository, AuthRepositoryType } from '@/repository/auth/repository';
import { SigninSeed, SignupSeed, Token } from '@/components/model/auth/type'
import { CreateUserMutation } from '@/api/main/mutation/createUser.mutation'
import { useCurrentUserMutators, SetCurrentUserType } from '@/globalStates/currentUserState'
import { TokenRepositoryType, useTokenRepository } from '@/repository/token/repository'
import { getSdk as getCurrentUserSdk, Sdk as CurrentUserSdk } from '@/api/main/query/__generated__/currentUser.query.generated'
import { getSdk as getCreateUserSdk, Sdk as CreateUserSdk } from '@/api/main/mutation/__generated__/createUser.mutation.generated'
import { gqlClient } from '@/api/main/gqlClient'

export const useAuthUsecase = () => {
  const authRepository = useAuthRepository()
  const tokenRepository = useTokenRepository()
  const { setCurrentUser } = useCurrentUserMutators();
  const currentUserSdk = getCurrentUserSdk(gqlClient);
  const createUserSdk = getCreateUserSdk(gqlClient);

  return React.useMemo(
    () => createAuthUsecase({
      authRepository,
      tokenRepository,
      setCurrentUser,
      currentUserSdk,
      createUserSdk,
    }),
    [authRepository, tokenRepository, setCurrentUser, currentUserSdk, createUserSdk],
  )
}

export const createAuthUsecase = ({ authRepository, tokenRepository, setCurrentUser, currentUserSdk, createUserSdk }: {
  authRepository: AuthRepositoryType
  tokenRepository: TokenRepositoryType
  setCurrentUser: SetCurrentUserType
  currentUserSdk: CurrentUserSdk
  createUserSdk: CreateUserSdk
}) => ({
  async signup(seed: SignupSeed, name: string) {
    try {
      // get token
      const token: Token = await authRepository.signup(seed)
      // set cookie
      await tokenRepository.setToken(token)
      // create user
      const data = await createUserSdk.CreateUser({ name })
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
      // set cookie
      await tokenRepository.setToken(token)
      // get current user
      const data = await currentUserSdk.CurrentUser()
      // set state
      setCurrentUser(data.currentUser.id, data.currentUser.name)
    } catch (error) {
      const errorMessage = 'unexpected error'
      throw new Error(errorMessage)
    }
  }
})
