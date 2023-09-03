import { useRouter } from 'next/router'
import React from 'react'
import { RecoilRoot } from 'recoil'

import { AppInitializer } from '@/components/ui/global/AppInitializer'
import { Header } from '@/components/ui/global/Header'
import { Loading } from '@/components/ui/global/Loading'

import { KnowledgeDetailRoot } from './KnowledgeDetail'

export const KnowledgeDetailPage = () => {
  const router = useRouter()
  const { id } = router.query

  if (typeof id !== 'string') {
    return <Loading />
  }

  return (
    <RecoilRoot>
      <AppInitializer>
        <Header />
        <KnowledgeDetailRoot id={id} />
      </AppInitializer>
    </RecoilRoot>
  )
}
