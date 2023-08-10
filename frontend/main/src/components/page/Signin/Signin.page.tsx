import { RecoilRoot } from 'recoil'
import { Signin } from '@/components/page/Signin/index'
import { Header } from '@/components/page/global/header'

export const SigninPage = () => {
  return (
    <RecoilRoot>
      <Header/>
      <Signin/>
    </RecoilRoot>
  )
}