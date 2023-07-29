import React from 'react';
import { httpClient } from './axios';
import { Token } from '@/components/model/auth/type'
import * as http from 'http'

export type SetTokenReqData = {
  token: Token;
}

export const useApiClient = () => {
  const apiClient = React.useMemo(() => new ApiClient(), []);

  return apiClient;
};

export class ApiClient {
  async setToken({ token }: SetTokenReqData): Promise<void> {
    const response = await httpClient.post('/set_token', { token });
  }
}