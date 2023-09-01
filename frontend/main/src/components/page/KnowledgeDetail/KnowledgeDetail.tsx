import React from 'react'

import { KnowledgeDetail } from '@/components/model/knowledge/KnowledgeDetail/KnowledgeDetail'
import { Loading } from '@/components/page/global/Loading'
import { useKnowledgeItemForDetail } from '@/usecase/knowledge/reader'

type KnowledgeDetailRootProps = {
  id: string
}

export const KnowledgeDetailRoot: React.FC<KnowledgeDetailRootProps> = (props) => {
  const { data, isLoading } = useKnowledgeItemForDetail({ id: props.id })

  if (isLoading) return <Loading />

  return (
    <div>{data && data.knowledge && <KnowledgeDetail knowledge={data.knowledge} />}</div>
  )
}
