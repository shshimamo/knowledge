import React from 'react'
import { KnowledgeListItem, knowledgeListItemFragment } from '@/components/model/knowledge/KnowledgeListItem/KnowledgeListItem'
import { FragmentType } from '@/gql/__generated__'
import styles from './KnowledgeList.module.css';

type KnowledgeListProps = {
  knowledgeList: Array<FragmentType<typeof knowledgeListItemFragment>>;
};

export const KnowledgeList: React.FC<KnowledgeListProps> = (props) => {
  return (
    <div className={styles.knowledgeList}>
      {props.knowledgeList.map((knowledge, index) => (
        <KnowledgeListItem key={index} knowledge={knowledge} />
      ))}
    </div>
  );
};
