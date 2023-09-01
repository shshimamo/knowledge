import { RecoilRoot } from 'recoil'

import { Header } from '@/components/page/global/Header'
import { Signin } from '@/components/page/Signin/Signin'

export const SigninPage = () => {
  return (
    <RecoilRoot>
      <Header />
      <Signin />
    </RecoilRoot>
  )
}
