import React from 'react'
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';

type CurrentUserState = {
  id: string | null;
  name: string | null;
}

const currentUserRecoilState = atom<CurrentUserState>({
  key: 'userState',
  default: { id: null, name: null },
});

export const useCurrentUserState = () => {
  return useRecoilValue(currentUserRecoilState);
}

export const useCurrentUserMutators = () => {
  const setState = useSetRecoilState(currentUserRecoilState);

  const setCurrentUser = React.useCallback(
    (id: string, name: string | null) => setState({ id, name }),
    [setState]
  )

  return { setCurrentUser }
}
export type SetCurrentUser = ReturnType<typeof useCurrentUserMutators>["setCurrentUser"];
