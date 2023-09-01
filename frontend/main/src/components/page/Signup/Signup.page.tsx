import { RecoilRoot } from 'recoil'

import { Header } from '@/components/page/global/Header'
import { Signup } from '@/components/page/Signup/Signup'

export const SignupPage = () => {
  return (
    <RecoilRoot>
      <Header />
      <Signup />
    </RecoilRoot>
  )
}
