import { useRouter } from 'next/router'
import React from 'react'
import { RecoilRoot } from 'recoil'

import { KnowledgeEditRoot } from '@/components/page/KnowledgeEdit/KnowledgeEdit'
import { AppInitializer } from '@/components/ui/global/AppInitializer'
import { Header } from '@/components/ui/global/Header'
import { Loading } from '@/components/ui/global/Loading'

export const KnowledgeEditPage = () => {
  const router = useRouter()
  const { id } = router.query

  if (typeof id !== 'string') {
    return <Loading />
  }

  return (
    <RecoilRoot>
      <AppInitializer>
        <Header />
        <KnowledgeEditRoot id={id} />
      </AppInitializer>
    </RecoilRoot>
  )
}
