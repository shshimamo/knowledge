import { RecoilRoot } from 'recoil'

import { Signup } from '@/components/page/Signup/Signup'
import { AppInitializer } from '@/components/ui/global/AppInitializer'
import { Header } from '@/components/ui/global/Header'
import { UnsignedInOnly } from '@/components/ui/global/UnsignedInOnly'

export const SignupPage = () => {
  return (
    <RecoilRoot>
      <AppInitializer>
        <UnsignedInOnly>
          <Header />
          <Signup />
        </UnsignedInOnly>
      </AppInitializer>
    </RecoilRoot>
  )
}
