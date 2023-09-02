import { RecoilRoot } from 'recoil'

import { Signin } from '@/components/page/Signin/Signin'
import { Header } from '@/components/ui/global/Header'

export const SigninPage = () => {
  return (
    <RecoilRoot>
      <Header />
      <Signin />
    </RecoilRoot>
  )
}
