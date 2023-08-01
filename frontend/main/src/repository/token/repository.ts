import React from 'react';
import { useApiClient, ApiClient, SetTokenReqData } from '@/api/main/apiClient'
import { convertTokenToReqData } from '@/repository/token/converter'
import { Token } from '@/components/model/auth/type'

export type TokenRepositoryType = ReturnType<typeof createTokenRepository>

export const useTokenRepository = () => {
  const apiClient = useApiClient();
  return React.useMemo(() => createTokenRepository(apiClient), [apiClient])
}

export const createTokenRepository = (apiClient: ApiClient) => ({
  async setToken(token: Token): Promise<void> {
    const reqData = convertTokenToReqData(token);
    await apiClient.setToken(reqData);
  }
})