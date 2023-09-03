import { RecoilRoot } from 'recoil'

import { Signin } from '@/components/page/Signin/Signin'
import { AppInitializer } from '@/components/ui/global/AppInitializer'
import { Header } from '@/components/ui/global/Header'

export const SigninPage = () => {
  return (
    <RecoilRoot>
      <AppInitializer>
        <Header />
        <Signin />
      </AppInitializer>
    </RecoilRoot>
  )
}
