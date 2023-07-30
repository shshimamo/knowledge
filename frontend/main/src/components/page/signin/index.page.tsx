import { ApolloProvider } from '@apollo/client'
import client from "@/apolloClient";
import { RecoilRoot } from 'recoil'
import { Signin } from '@/components/page/signin/index'
import { Header } from '@/components/page/global/header'

export const SigninPage = () => {
  return (
    <ApolloProvider client={client}>
      <RecoilRoot>
        <Header/>
        <Signin/>
      </RecoilRoot>
    </ApolloProvider>
  )
}