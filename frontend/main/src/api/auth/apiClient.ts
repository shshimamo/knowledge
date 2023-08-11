import React from 'react';
import { httpClient } from './axios';

export type SignupAPIReqParams = {
  email: string;
  password: string;
}

export type SignupAPIResponse = {
  token: string;
}

export type SigninAPIReqParams = {
  email: string;
  password: string;
}

export type SigninAPIResponse = {
  token: string;
}


export const useApiClient = () => {
  return React.useMemo(() => createApiClient(), []);
};

const createApiClient  = () => ({
  async signup({ email, password }: SignupAPIReqParams): Promise<SignupAPIResponse> {
    const response = await httpClient.post('/signup', {
      email,
      password,
    });

    return response.data as SignupAPIResponse;
  },

  async signin({ email, password }: SigninAPIReqParams): Promise<SigninAPIResponse> {
    const response = await httpClient.post('/signin', {
      email,
      password
    });

    return response.data as SigninAPIResponse;
  }
});

export type ApiClient = ReturnType<typeof createApiClient>;
