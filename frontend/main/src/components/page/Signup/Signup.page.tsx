import { RecoilRoot } from 'recoil'

import { Signup } from '@/components/page/Signup/Signup'
import { AppInitializer } from '@/components/ui/global/AppInitializer'
import { Header } from '@/components/ui/global/Header'

export const SignupPage = () => {
  return (
    <RecoilRoot>
      <AppInitializer>
        <Header />
        <Signup />
      </AppInitializer>
    </RecoilRoot>
  )
}
