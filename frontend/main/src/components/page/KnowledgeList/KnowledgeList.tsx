import React from 'react';
import { useMyKnowledgeList } from '@/usecase/knowledge/reader'
import { KnowledgeList } from '@/components/model/knowledge/KnowledgeList/KnowledgeList'
import { useSWRConfig } from 'swr'

export const KnowledgeListTop = () => {
  const { data, error, isLoading } = useMyKnowledgeList({ first: 10})

  if (error) return <div>failed to load</div>
  if (isLoading) return <div>loading...</div>

  return (
    <div>
      {data && data.currentUser.knowledgeList && (
        <KnowledgeList knowledgeList={data.currentUser.knowledgeList} />
      )}
    </div>
  );
};