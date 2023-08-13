import React from 'react'
import { useSWRConfig } from "swr"

export const userCacheKeyGenerator = {
  currentUserKey: () =>  ['USER', 'CurrentUser'] as const,
}

export const useUserCacheMutator = () => {
  const { mutate } = useSWRConfig()

  return React.useMemo(
    () => ({
      mutateCurrentUser: () => {
        return mutate(
          (key: string[]) => key[0] == 'USER' && key[1] == 'CurrentUser',
          undefined,
          { revalidate: true }
        )
      }
    }),
    [mutate],
  )
}