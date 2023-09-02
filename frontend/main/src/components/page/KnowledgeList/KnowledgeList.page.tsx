import { RecoilRoot } from 'recoil'

import { Header } from '@/components/ui/global/Header'

import { KnowledgeListTop } from './KnowledgeList'

export const KnowledgeListPage = () => {
  return (
    <RecoilRoot>
      <Header />
      <KnowledgeListTop />
    </RecoilRoot>
  )
}
