import React, { useEffect, useState } from 'react';
import { gqlClient } from '@/api/main/gqlClient'
import { useMyKnowledgeList } from '@/usecase/knowledge/reader'
import { KnowledgeList } from '@/components/model/knowledge/KnowledgeList/KnowledgeList'

export const KnowledgeListTop = () => {
  const { data } = useMyKnowledgeList({ first: 10})
  const knowledgeList = data && data.data?.currentUser?.knowledgeList

  return (
    <div>
      {knowledgeList && (
        <KnowledgeList knowledgeList={knowledgeList} />
      )}
    </div>
  );
};