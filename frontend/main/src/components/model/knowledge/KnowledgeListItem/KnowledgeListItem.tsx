import { useRouter } from 'next/router'
import React, { useState } from 'react'

import { Modal } from '@/components/ui/modal/Modal'
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
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState(false)

  const handleEdit = () => {
    router.push(`/knowledge/${knowledge.id}`).catch((error) => {
      console.error(error)
    })
  }

  const showDeleteModal = () => {
    setIsDeleteModalOpen(true)
  }

  const onDeleteModalNo = () => {
    setIsDeleteModalOpen(false)
  }

  const onDeleteModalYes = () => {
    knowledgeUsecase.deleteKnowledge(knowledge.id).catch((error) => {
      console.error(error)
      setIsDeleteModalOpen(false)
    })
  }

  return (
    <div className={`${styles.knowledgeItem} e2e-model-knowledge-list-item`}>
      <h3 className={`${styles.knowledgeTitle} e2e-model-knowledge-list-item-title`}>{knowledge.title}</h3>
      <p className={`${styles.knowledgeText} e2e-model-knowledge-list-item-text`}>{knowledge.text}</p>
      <div className={styles.buttons}>
        <span className={`${knowledge.isPublic ? styles.publicButton : styles.privateButton} e2e-model-knowledge-list-item-public`}>
          {knowledge.isPublic ? 'Public' : 'Private'}
        </span>
        <button className={`${styles.editButton} e2e-model-knowledge-list-item-show-button`} onClick={handleEdit}>
          Show
        </button>
        <button className={`${styles.deleteButton} e2e-model-knowledge-list-item-delete-button`} onClick={showDeleteModal}>
          Delete
        </button>
        {isDeleteModalOpen && (
          <Modal text={'Delete?'} onNo={onDeleteModalNo} onYes={onDeleteModalYes} />
        )}
      </div>
    </div>
  )
}
