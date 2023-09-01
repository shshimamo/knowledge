import { useRouter } from 'next/router'
import React from 'react'
import { RecoilRoot } from 'recoil'

import { Header } from '@/components/page/global/Header'
import { Loading } from '@/components/page/global/Loading'
import { KnowledgeEditRoot } from '@/components/page/KnowledgeEdit/KnowledgeEdit'

export const KnowledgeEditPage = () => {
  const router = useRouter()
  const { id } = router.query

  if (typeof id !== 'string') {
    return <Loading />
  }

  return (
    <RecoilRoot>
      <Header />
      <KnowledgeEditRoot id={id} />
    </RecoilRoot>
  )
}
