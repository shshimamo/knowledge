import React from 'react'
import { useAuthRepository, AuthRepository } from '@/repository/auth/repository';
import { SignupSeed, Token } from '@/components/model/auth/type'
import { CreateUserMutation } from '@/api/main/mutation/createUser.mutation'
import { useCurrentUserMutators, SetCurrentUser } from '@/globalStates/currentUserState'
import { TokenRepository, useTokenRepository } from '@/repository/token/repository'

export const useAuthUsecase = (createUserMutation: CreateUserMutation) => {
  const authRepository = useAuthRepository()
  const tokenRepository = useTokenRepository()
  const { setCurrentUser } = useCurrentUserMutators();

  return React.useMemo(
    () => createAuthUsecase({
      authRepository,
      tokenRepository,
      setCurrentUser,
      createUserMutation,
    }),
    [authRepository, setCurrentUser, createUserMutation],
  )
}

export const createAuthUsecase = ({ authRepository, tokenRepository, setCurrentUser, createUserMutation }: {
  authRepository: AuthRepository
  tokenRepository: TokenRepository
  setCurrentUser: SetCurrentUser
  createUserMutation: CreateUserMutation
}) => ({
  async signup(seed: SignupSeed, name: string) {
    try {
      const token: Token = await authRepository.signup(seed)

      await tokenRepository.setToken(token)

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
})
