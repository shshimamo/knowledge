import { useRouter } from 'next/router'
import React from 'react'
import { RecoilRoot } from 'recoil'

import { Header } from '@/components/page/global/Header'
import { Loading } from '@/components/page/global/Loading'

import { KnowledgeDetailRoot } from './KnowledgeDetail'

export const KnowledgeDetailPage = () => {
  const router = useRouter()
  const { id } = router.query

  if (typeof id !== 'string') {
    return <Loading />
  }

  return (
    <RecoilRoot>
      <Header />
      <KnowledgeDetailRoot id={id} />
    </RecoilRoot>
  )
}
