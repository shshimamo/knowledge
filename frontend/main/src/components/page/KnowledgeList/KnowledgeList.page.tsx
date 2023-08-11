import { RecoilRoot } from 'recoil'
import { Header } from '@/components/page/global/Header'
import { KnowledgeListTop } from './KnowledgeList'

export const KnowledgeListPage = () => {
  return (
    <RecoilRoot>
      <Header/>
      <KnowledgeListTop/>
    </RecoilRoot>
  )
}