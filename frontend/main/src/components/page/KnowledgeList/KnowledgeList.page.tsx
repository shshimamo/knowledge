import { RecoilRoot } from 'recoil'
import { Header } from '@/components/page/global/Header'
import { KnowledgeListTop } from './KnowledgeList'
import { SWRConfig } from 'swr'

export const KnowledgeListPage = () => {
  return (
    <RecoilRoot>
      <Header/>
      <KnowledgeListTop/>
    </RecoilRoot>
  )
}