import React, { PropsWithChildren, useEffect } from 'react'

import { useCurrentUserInitializer } from '@/usecase/app_initializer/current_user_initializer/reader'

export const AppInitializer: React.FC<PropsWithChildren> = ({ children }) => {
  const currentUserInitializer = useCurrentUserInitializer()

  useEffect(() => {
    currentUserInitializer.initialize()
  }, [currentUserInitializer])

  return <>{children}</>
}
