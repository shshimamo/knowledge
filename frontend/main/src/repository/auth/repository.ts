import React from 'react';
import { ApiClient, useApiClient } from '@/api/auth/apiClient';
import {
  convertSigninResponseToToken,
  convertSigninSeedToReqParams,
  convertSignupResponseToToken,
  convertSignupSeedToReqParams
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
    const reqParams = convertSignupSeedToReqParams(seed)
    const res = await apiClient.signup(reqParams);
    return convertSignupResponseToToken(res);
  },

  async signin(seed: SigninSeed): Promise<Token> {
    const reqParams = convertSigninSeedToReqParams(seed)
    const res = await apiClient.signin(reqParams);
    return convertSigninResponseToToken(res);
  },
});