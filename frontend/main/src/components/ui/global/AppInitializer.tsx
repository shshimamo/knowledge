import React, { PropsWithChildren, useEffect } from 'react'

import { useAppInitializedMutators } from '@/globalStates/appInitializedState'
import { useCurrentUserMutators } from '@/globalStates/currentUserState'
import { useCurrentUser } from '@/usecase/user/reader'

export const AppInitializer: React.FC<PropsWithChildren> = ({ children }) => {
  const currentUserRes = useCurrentUser()
  const { setCurrentUser } = useCurrentUserMutators()
  const { setAppInitialized } = useAppInitializedMutators()

  useEffect(() => {
    if (currentUserRes.isValidating === false) {
      if (currentUserRes.data?.currentUser) {
        setCurrentUser(
          currentUserRes.data?.currentUser.id,
          currentUserRes.data?.currentUser.name,
        )
      }
      setAppInitialized({ currentUserInitialized: true })
    }
  }, [currentUserRes, setCurrentUser])

  return <>{children}</>
}
