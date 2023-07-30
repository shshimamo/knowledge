import React from 'react';
import { httpClient } from './axios';

export type SignupReqData = {
  email: string;
  password: string;
}

export type SignupResData = {
  token: string;
}

export type SigninReqData = {
  email: string;
  password: string;
}

export type SigninResData = {
  token: string;
}


export const useApiClient = () => {
  const apiClient = React.useMemo(() => createApiClient(), []);

  return apiClient;
};

export const createApiClient  = () => ({
  async signup({ email, password }: SignupReqData): Promise<SignupResData> {
    const response = await httpClient.post('/signup', {
      email,
      password,
    });

    return response.data as SignupResData;
  },

  async signin({ email, password }: SigninReqData): Promise<SigninResData> {
    const response = await httpClient.post('/signin', {
      email,
      password
    });

    return response.data as SigninResData;
  }
});

export type CreateApiClientReturn = ReturnType<typeof createApiClient>;
