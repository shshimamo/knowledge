/* KnowledgeDetail Model */
import Link from 'next/link'
import React from 'react'

import { FragmentType, graphql, useFragment } from '@/gql/__generated__'

import styles from './KnowledgeDetail.module.css'

export const knowledgeDetailFragment = graphql(/* GraphQL */ `
  fragment KnowledgeDetail on Knowledge {
    id
    userId
    title
    text
    isPublic
    publishedAt
  }
`)

type KnowledgeDetailProps = {
  knowledge: FragmentType<typeof knowledgeDetailFragment>
}

export const KnowledgeDetail: React.FC<KnowledgeDetailProps> = (props) => {
  const knowledge = useFragment(knowledgeDetailFragment, props.knowledge)

  return (
    <div className={styles.container}>
      <Link href='/knowledge_list' className={styles.backButton}>
        &larr; Back to List
      </Link>

      <div className={styles.flame}>
        <h2 className={styles.title}>{knowledge.title}</h2>
        <p className={styles.text}>{knowledge.text}</p>
        <br />
        <div className={styles.info}>
          <p>Public: {knowledge.isPublic ? 'Yes' : 'No'}</p>
          <p>Published At: {knowledge.publishedAt}</p>
        </div>
        <br />
        <Link href={`/knowledge/${knowledge.id}/edit`} className={styles.button}>
          Edit
        </Link>
      </div>
    </div>
  )
}
