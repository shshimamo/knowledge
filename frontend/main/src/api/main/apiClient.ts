import React from 'react'

import { Token } from '@/components/model/auth/type'

import { httpClient } from './axios'

export type SetTokenAPIReqParams = {
  token: Token
}

export const useApiClient = () => {
  return React.useMemo(() => createApiClient(), [])
}

const createApiClient = () => ({
  async setToken({ token }: SetTokenAPIReqParams): Promise<void> {
    await httpClient.post('/set_token', { token })
  },
})

export type ApiClient = ReturnType<typeof createApiClient>
