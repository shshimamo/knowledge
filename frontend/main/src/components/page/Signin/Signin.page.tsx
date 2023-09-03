import { RecoilRoot } from 'recoil'

import { Signin } from '@/components/page/Signin/Signin'
import { AppInitializer } from '@/components/ui/global/AppInitializer'
import { Header } from '@/components/ui/global/Header'
import { UnsignedInOnly } from '@/components/ui/global/UnsignedInOnly'

export const SigninPage = () => {
  return (
    <RecoilRoot>
      <AppInitializer>
        <UnsignedInOnly>
          <Header />
          <Signin />
        </UnsignedInOnly>
      </AppInitializer>
    </RecoilRoot>
  )
}
