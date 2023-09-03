import { useRouter } from 'next/router'
import React, { PropsWithChildren, useEffect, useState } from 'react'

import { useCurrentUserState } from '@/globalStates/currentUserState'
import { useInitCurrentUserState } from '@/globalStates/initCurrentUserState'

export const UnsignedInOnly: React.FC<PropsWithChildren> = ({ children }) => {
  const initCurrentUserState = useInitCurrentUserState()
  const currentUser = useCurrentUserState()
  const router = useRouter()
  const [shouldRedirect, setShouldRedirect] = useState(false)

  useEffect(() => {
    if (initCurrentUserState.isInitialized && currentUser.id) {
      setShouldRedirect(true) // For UI flickering
      void router.push('/knowledge_list')
    }
  }, [currentUser, router])

  if (!initCurrentUserState.isInitialized || shouldRedirect) {
    return <div>loading...</div>
  }

  return <>{children}</>
}
