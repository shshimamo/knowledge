import React from 'react'

import {
  SetCurrentUserType,
  useCurrentUserMutators,
} from '@/globalStates/currentUserState'
import { useCurrentUser, UseCurrentUserType } from '@/usecase/user/reader'

export const useAppInitializer = () => {
  const currentUserRes = useCurrentUser()
  const { setCurrentUser } = useCurrentUserMutators()
  return React.useMemo(
    () =>
      createAppInitializer({
        currentUserRes,
        setCurrentUser,
      }),
    [currentUserRes, setCurrentUser],
  )
}

export const createAppInitializer = ({
  currentUserRes,
  setCurrentUser,
}: {
  currentUserRes: UseCurrentUserType
  setCurrentUser: SetCurrentUserType
}) => ({
  initialize() {
    const currentUser = currentUserRes.data?.currentUser
    if (currentUser) {
      setCurrentUser(currentUser.id, currentUser.name)
    }
  },
})
