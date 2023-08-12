import { useRouter } from 'next/router';
import { RecoilRoot } from 'recoil'
import { Header } from '@/components/page/global/Header'
import { KnowledgeDetailTop } from './KnowledgeDetail'
import React from 'react'
import { Loading } from '@/components/page/global/Loading'

export const KnowledgeDetailPage = () => {
  const router = useRouter();
  const { id } = router.query;

  if (typeof id !== "string") {
    return <Loading/>;
  }

  return (
    <RecoilRoot>
      <Header/>
      <KnowledgeDetailTop id={id}/>
    </RecoilRoot>
  )
}