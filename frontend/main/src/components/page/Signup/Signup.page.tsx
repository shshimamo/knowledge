import { RecoilRoot } from 'recoil'

import { Signup } from '@/components/page/Signup/Signup'
import { Header } from '@/components/ui/global/Header'

export const SignupPage = () => {
  return (
    <RecoilRoot>
      <Header />
      <Signup />
    </RecoilRoot>
  )
}
