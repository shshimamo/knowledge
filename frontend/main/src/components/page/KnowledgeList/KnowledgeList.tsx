import React from 'react'

import { KnowledgeList } from '@/components/model/knowledge/KnowledgeList/KnowledgeList'
import { useMyKnowledgeList } from '@/usecase/knowledge/reader'

export const KnowledgeListTop = () => {
  const { data, isLoading } = useMyKnowledgeList({ first: 10 })

  if (isLoading) return <div>loading...</div>

  return (
    <div>
      {data && data.currentUser.knowledgeList && (
        <KnowledgeList knowledgeList={data.currentUser.knowledgeList} />
      )}
    </div>
  )
}
