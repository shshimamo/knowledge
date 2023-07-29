import React from 'react';
import { httpClient } from './axios';

export type SignupReqData = {
  email: string;
  password: string;
}

export type SignupResData = {
  token: string;
}


export const useApiClient = () => {
  const apiClient = React.useMemo(() => new ApiClient(), []);

  return apiClient;
};

export class ApiClient {
  async signup({ email, password }: SignupReqData): Promise<SignupResData> {
    const response = await httpClient.post('/signup', {
      email,
      password,
    });

    return response.data as SignupResData;
  }
}