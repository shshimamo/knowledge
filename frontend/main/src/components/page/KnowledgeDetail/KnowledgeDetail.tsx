import { useRouter } from 'next/router'
import React, { useEffect } from 'react'

import { KnowledgeDetail } from '@/components/model/knowledge/KnowledgeDetail/KnowledgeDetail'
import { Loading } from '@/components/ui/global/Loading'
import { useKnowledgeItemForDetail } from '@/usecase/knowledge/reader'

type KnowledgeDetailRootProps = {
  id: string
}

export const KnowledgeDetailRoot: React.FC<KnowledgeDetailRootProps> = (props) => {
  const { data, isLoading } = useKnowledgeItemForDetail({ id: props.id })
  const router = useRouter()

  useEffect(() => {
    if (!data && !isLoading) {
      router.push('/knowledge_list').catch((error) => {
        console.error('Failed to redirect:', error)
      })
    }
  }, [data, isLoading, router])

  if (isLoading) return <Loading />

  return (
    <div>{data && data.knowledge && <KnowledgeDetail knowledge={data.knowledge} />}</div>
  )
}
