import React from 'react';
import { httpClient } from './axios';
import { Token } from '@/components/model/auth/type'

export type SetTokenAPIReqParams = {
  token: Token;
}

export const useApiClient = () => {
  return React.useMemo(() => createApiClient(), []);
};

const createApiClient = () => ({
  async setToken({ token }: SetTokenAPIReqParams): Promise<void> {
    const response = await httpClient.post('/set_token', { token });
  }
});

export type ApiClient = ReturnType<typeof createApiClient>;
