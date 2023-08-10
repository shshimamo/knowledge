import React from 'react';
import { ApiClient, useApiClient } from '@/api/auth/apiClient';
import {
  convertSigninResDataToToken,
  convertSigninSeedToReqData,
  convertSignupResDataToToken,
  convertSignupSeedToReqData
} from './converter';
import { SigninSeed, SignupSeed, Token } from '@/components/model/auth/type'

export type AuthRepositoryType = ReturnType<typeof createAuthRepository>

// hooks
// 必要なhooksを呼び出してその値を repository factory関数に渡す
export const useAuthRepository = () => {
  const apiClient = useApiClient()
  return React.useMemo(() => createAuthRepository(apiClient), [apiClient])
}

// factory関数
export const createAuthRepository = (apiClient: ApiClient) => ({
  async signup(seed: SignupSeed): Promise<Token> {
    const reqData = convertSignupSeedToReqData(seed)
    const resData = await apiClient.signup(reqData);
    return convertSignupResDataToToken(resData);
  },

  async signin(seed: SigninSeed): Promise<Token> {
    const reqData = convertSigninSeedToReqData(seed)
    const resData = await apiClient.signin(reqData);
    return convertSigninResDataToToken(resData);
  },
});