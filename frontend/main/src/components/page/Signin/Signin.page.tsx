import { RecoilRoot } from 'recoil'
import { Signin } from '@/components/page/Signin/Signin'
import { Header } from '@/components/page/global/Header'

export const SigninPage = () => {
  return (
    <RecoilRoot>
      <Header/>
      <Signin/>
    </RecoilRoot>
  )
}