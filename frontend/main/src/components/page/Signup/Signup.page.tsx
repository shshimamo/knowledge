import { RecoilRoot } from 'recoil'
import { Signup } from '@/components/page/Signup/Signup'
import { Header } from '@/components/page/global/Header'

export const SignupPage = () => {
  return (
    <RecoilRoot>
      <Header/>
      <Signup/>
    </RecoilRoot>
  )
}