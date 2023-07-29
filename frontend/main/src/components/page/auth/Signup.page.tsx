import { RecoilRoot } from 'recoil'
import { Signup } from '@/components/page/auth/Signup'
import { Header } from '@/components/page/global/header'

export const SignupPage = () => {
  return (
    <RecoilRoot>
      <Header/>
      <Signup/>
    </RecoilRoot>
  )
}