import { ApolloProvider } from '@apollo/client'
import client from "@/apolloClient";
import { RecoilRoot } from 'recoil'
import { Signup } from '@/components/page/signup/index'
import { Header } from '@/components/page/global/header'

export const SignupPage = () => {
  return (
    <ApolloProvider client={client}>
      <RecoilRoot>
        <Header/>
        <Signup/>
      </RecoilRoot>
    </ApolloProvider>
  )
}