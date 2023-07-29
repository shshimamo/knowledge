import React from 'react'
import { useAuthRepository, AuthRepository } from '@/repository/auth/repository';
import { SignupSeed, Token } from '@/components/model/auth/type'

export const useAuthUsecase = () => {
  const repository = useAuthRepository()

  return React.useMemo(
    () => createAuthUsecase({ repository }),
    [repository],
  )
}

export const createAuthUsecase = ({ repository }: {
  repository: AuthRepository
}) => ({
  async signup(seed: SignupSeed, name: string) {
    try {
      const token: Token = await repository.signup(seed)
      // TODO: call main signup api and send name
      // TODO: set current_user to recoil
    } catch (error) {
      const errorMessage = '予期せぬエラーが発生しました。再度お試しください。'
      throw new Error(errorMessage)
    }
  },
})