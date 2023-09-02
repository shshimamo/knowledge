import { useRouter } from 'next/router'
import React from 'react'

import { FragmentType, graphql, useFragment } from '@/gql/__generated__'
import { useKnowledgeUsecase } from '@/usecase/knowledge/usecase'

import styles from './KnowledgeListItem.module.css'

export const knowledgeListItemFragment = graphql(/* GraphQL */ `
  fragment KnowledgeListItem on Knowledge {
    id
    title
    text
    isPublic
  }
`)

type KnowledgeListItemProps = {
  knowledge: FragmentType<typeof knowledgeListItemFragment>
}

export const KnowledgeListItem: React.FC<KnowledgeListItemProps> = (props) => {
  const knowledge = useFragment(knowledgeListItemFragment, props.knowledge)
  const router = useRouter()
  const knowledgeUsecase = useKnowledgeUsecase()

  const handleEdit = () => {
    router.push(`/knowledge/${knowledge.id}`).catch((error) => {
      console.error(error)
    })
  }

  const handleDelete = () => {
    knowledgeUsecase.deleteKnowledge(knowledge.id).catch((error) => {
      console.error(error)
    })
  }

  return (
    <div className={styles.knowledgeItem}>
      <h3 className={styles.knowledgeTitle}>{knowledge.title}</h3>
      <p className={styles.knowledgeText}>{knowledge.text}</p>
      <div className={styles.buttons}>
        <span className={knowledge.isPublic ? styles.publicButton : styles.privateButton}>
          {knowledge.isPublic ? 'Public' : 'Private'}
        </span>
        <button className={styles.editButton} onClick={handleEdit}>
          Show
        </button>
        <button className={styles.deleteButton} onClick={handleDelete}>
          Delete
        </button>
      </div>
    </div>
  )
}
