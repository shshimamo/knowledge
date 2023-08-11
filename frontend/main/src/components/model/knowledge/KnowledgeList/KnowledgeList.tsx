import React from 'react'
import { KnowledgeItem, knowledgeItemFragment } from '../KnowledgeItem/KnowledgeItem'
import { FragmentType } from '@/gql/__generated__'

type KnowledgeListProps = {
  knowledgeList: Array<FragmentType<typeof knowledgeItemFragment>>;
};

export const KnowledgeList: React.FC<KnowledgeListProps> = (props) => {
  return (
    <div>
      {props.knowledgeList.map((knowledge, index) => (
        <KnowledgeItem key={index} knowledge={knowledge} />
      ))}
    </div>
  );
};
