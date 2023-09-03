import { useRouter } from 'next/router'
import React, { PropsWithChildren, useEffect, useState } from 'react'

import { useAppInitializedState } from '@/globalStates/appInitializedState'
import { useCurrentUserState } from '@/globalStates/currentUserState'

export const UnsignedInOnly: React.FC<PropsWithChildren> = ({ children }) => {
  const appInitialized = useAppInitializedState()
  const currentUser = useCurrentUserState()
  const router = useRouter()
  const [shouldRedirect, setShouldRedirect] = useState(false)

  useEffect(() => {
    if (appInitialized.currentUserInitialized && currentUser.id) {
      setShouldRedirect(true) // For UI flickering
      void router.push('/knowledge_list')
    }
  }, [currentUser, router])

  if (!appInitialized.currentUserInitialized || shouldRedirect) {
    return <div>loading...</div>
  }

  return <>{children}</>
}
