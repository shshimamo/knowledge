import React from 'react';
import { useApiClient, ApiClient, SetTokenAPIReqParams } from '@/api/main/apiClient'
import { convertTokenToReqParams } from '@/repository/token/converter'
import { Token } from '@/components/model/auth/type'

export type TokenRepositoryType = ReturnType<typeof createTokenRepository>

export const useTokenRepository = () => {
  const apiClient = useApiClient();
  return React.useMemo(() => createTokenRepository(apiClient), [apiClient])
}

export const createTokenRepository = (apiClient: ApiClient) => ({
  async setToken(token: Token): Promise<void> {
    const reqParams = convertTokenToReqParams(token);
    await apiClient.setToken(reqParams);
  }
})