import React from 'react'
import { FragmentType, graphql, useFragment } from '@/gql/__generated__'
import styles from './KnowledgeItem.module.css';

export const knowledgeItemFragment = graphql(/* GraphQL */ `
    fragment KnowledgeItem on Knowledge {
        title
        text
    }
`);

type KnowledgeItemProps = {
  knowledge: FragmentType<typeof knowledgeItemFragment>
};

export const KnowledgeItem: React.FC<KnowledgeItemProps> = (props) => {
  const knowledge = useFragment(knowledgeItemFragment, props.knowledge)

  return (
      <div className={styles.knowledgeItem}>
        <h3 className={styles.knowledgeTitle}>{knowledge.title}</h3>
        <p className={styles.knowledgeText}>{knowledge.text}</p>
      </div>
    );
};
