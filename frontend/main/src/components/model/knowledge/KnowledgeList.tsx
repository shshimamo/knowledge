import { KnowledgeItemFragment } from './__generated__/KnowledgeItem.generated'
import { gqlClient } from '@/api/main/gqlClient'
import React, { useEffect, useState } from 'react'
import { KnowledgeItem } from './KnowledgeItem'

type KnowledgeListProps = {
  items: Array<KnowledgeItemFragment>;
};

export const KnowledgeList: React.FC<KnowledgeListProps> = ({ items }) => {
  return (
    <div>
      {items.map((knowledge, index) => (
        <KnowledgeItem key={index} knowledge={knowledge} />
      ))}
    </div>
  );
};
