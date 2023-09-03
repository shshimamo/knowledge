import React from 'react'

import {
  SetCurrentUserType,
  useCurrentUserMutators,
} from '@/globalStates/currentUserState'
import {
  SetInitCurrentUserType,
  useInitCurrentUserMutators,
} from '@/globalStates/initCurrentUserState'
import { useCurrentUser, UseCurrentUserType } from '@/usecase/user/reader'

export const useCurrentUserInitializer = () => {
  const currentUserRes = useCurrentUser()
  const { setCurrentUser } = useCurrentUserMutators()
  const { setInitCurrentUser } = useInitCurrentUserMutators()

  return React.useMemo(
    () =>
      createCurrentUserInitializer({
        currentUserRes,
        setCurrentUser,
        setInitCurrentUser,
      }),
    [currentUserRes, setCurrentUser, setInitCurrentUser],
  )
}

export const createCurrentUserInitializer = ({
  currentUserRes,
  setCurrentUser,
  setInitCurrentUser,
}: {
  currentUserRes: UseCurrentUserType
  setCurrentUser: SetCurrentUserType
  setInitCurrentUser: SetInitCurrentUserType
}) => ({
  initialize() {
    if (currentUserRes.isValidating === false) {
      if (currentUserRes.data?.currentUser) {
        setCurrentUser(
          currentUserRes.data?.currentUser.id,
          currentUserRes.data?.currentUser.name,
        )
      }
      setInitCurrentUser({ isInitialized: true })
    }
  },
})
