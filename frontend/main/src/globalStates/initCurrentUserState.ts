import React from 'react'
import { atom, useRecoilValue, useSetRecoilState } from 'recoil'

type InitCurrentUserState = {
  isInitialized: boolean
}

const initCurrentUserRecoilState = atom<InitCurrentUserState>({
  key: 'initCurrentUserState',
  default: { isInitialized: false },
})

export const useInitCurrentUserState = () => {
  return useRecoilValue(initCurrentUserRecoilState)
}

export const useInitCurrentUserMutators = () => {
  const setState = useSetRecoilState(initCurrentUserRecoilState)

  const setInitCurrentUser = React.useCallback(
    (updates: Partial<InitCurrentUserState>) => {
      setState((prevState) => ({ ...prevState, ...updates }))
    },
    [setState],
  )

  return { setInitCurrentUser }
}

export type SetInitCurrentUserType = ReturnType<
  typeof useInitCurrentUserMutators
>['setInitCurrentUser']
