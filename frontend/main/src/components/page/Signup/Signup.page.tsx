import { RecoilRoot } from 'recoil'
import { Signup } from '@/components/page/Signup/index'
import { Header } from '@/components/page/global/header'

export const SignupPage = () => {
  return (
    <RecoilRoot>
      <Header/>
      <Signup/>
    </RecoilRoot>
  )
}