import React, { PropsWithChildren, useEffect } from 'react'

import { useAppInitializer } from '@/usecase/app_initializer/usecase'

export const AppInitializer: React.FC<PropsWithChildren> = ({ children }) => {
  const appInitializer = useAppInitializer()

  useEffect(() => {
    appInitializer.initialize()
  }, [appInitializer])

  return <>{children}</>
}
