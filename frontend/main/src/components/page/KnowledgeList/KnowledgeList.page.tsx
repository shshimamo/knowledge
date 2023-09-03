import { RecoilRoot } from 'recoil'

import { AppInitializer } from '@/components/ui/global/AppInitializer'
import { Header } from '@/components/ui/global/Header'

import { KnowledgeListTop } from './KnowledgeList'

export const KnowledgeListPage = () => {
  return (
    <RecoilRoot>
      <AppInitializer>
        <Header />
        <KnowledgeListTop />
      </AppInitializer>
    </RecoilRoot>
  )
}
