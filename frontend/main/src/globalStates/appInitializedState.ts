import React from 'react'
import { atom, useRecoilValue, useSetRecoilState } from 'recoil'

type AppInitializedState = {
  currentUserInitialized: boolean
}

const appInitializedRecoilState = atom<AppInitializedState>({
  key: 'appInitializedState',
  default: { currentUserInitialized: false },
})

export const useAppInitializedState = () => {
  return useRecoilValue(appInitializedRecoilState)
}

export const useAppInitializedMutators = () => {
  const setState = useSetRecoilState(appInitializedRecoilState)

  const setAppInitialized = React.useCallback(
    (updates: Partial<AppInitializedState>) => {
      setState((prevState) => ({ ...prevState, ...updates }))
    },
    [setState],
  )

  return { setAppInitialized }
}
