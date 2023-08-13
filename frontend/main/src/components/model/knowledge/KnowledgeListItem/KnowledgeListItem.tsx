import React from 'react'
import { useRouter } from 'next/router';
import styles from './KnowledgeListItem.module.css';
import { FragmentType, graphql, useFragment } from '@/gql/__generated__'
import { useKnowledgeUsecase } from '@/usecase/knowledge/usecase'
import { useKnowledgeCacheMutator } from '@/usecase/knowledge/cache'

export const knowledgeListItemFragment = graphql(/* GraphQL */ `
    fragment KnowledgeListItem on Knowledge {
        id
        title
        text
    }
`);

type KnowledgeListItemProps = {
  knowledge: FragmentType<typeof knowledgeListItemFragment>
};

export const KnowledgeListItem: React.FC<KnowledgeListItemProps> = (props) => {
  const knowledge = useFragment(knowledgeListItemFragment, props.knowledge)
  const router = useRouter();
  const { deleteKnowledge } = useKnowledgeUsecase();
  const { mutateAllKnowledgeList, mutateAllKnowledgeItem } = useKnowledgeCacheMutator();

  const handleEdit = async () => {
    await router.push(`/knowledge/${knowledge.id}`);
  };

  const handleDelete = async () => {
    await deleteKnowledge(knowledge.id)
    await mutateAllKnowledgeList()
    await mutateAllKnowledgeItem({ id: knowledge.id })
  }

  return (
      <div className={styles.knowledgeItem}>
        <h3 className={styles.knowledgeTitle}>{knowledge.title}</h3>
        <p className={styles.knowledgeText}>{knowledge.text}</p>
        <div className={styles.buttons}>
          <button className={styles.editButton} onClick={handleEdit}>Show</button>
          <button className={styles.deleteButton} onClick={handleDelete}>Delete</button>
        </div>
      </div>
    );
};
