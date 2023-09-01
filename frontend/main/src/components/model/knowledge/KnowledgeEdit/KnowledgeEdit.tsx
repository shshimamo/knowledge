/* KnowledgeEdit Model */
import Link from 'next/link'
import { useRouter } from 'next/router'
import React, { useState } from 'react'

import { FragmentType, graphql, useFragment } from '@/gql/__generated__'
import { useKnowledgeUsecase } from '@/usecase/knowledge/usecase'

import styles from './KnowledgeEdit.module.css'

// prettier-ignore
export const knowledgeEditFragment = graphql(/* GraphQL */ `
    fragment KnowledgeEdit on Knowledge {
        id
        userId
        title
        text
        isPublic
        publishedAt
    }
`)

type KnowledgeEditProps = {
  knowledge: FragmentType<typeof knowledgeEditFragment>
}

export const KnowledgeEdit: React.FC<KnowledgeEditProps> = (props) => {
  const router = useRouter()
  const knowledge = useFragment(knowledgeEditFragment, props.knowledge)
  const [title, setTitle] = useState(knowledge.title)
  const [text, setText] = useState(knowledge.text)
  const [isPublic, setIsPublic] = useState(knowledge.isPublic)
  const knowledgeUsecase = useKnowledgeUsecase()

  const updateKnowledge = async () => {
    try {
      await knowledgeUsecase.updateKnowledge({
        id: knowledge.id,
        input: { title, text, isPublic },
      })
      await router.push(`/knowledge/${knowledge.id}`)
    } catch (error) {
      console.error(error)
    }
  }

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    updateKnowledge().catch((error) => {
      console.error(error)
    })
  }

  return (
    <div className={styles.customContainer}>
      <form onSubmit={handleSubmit} className={styles.form}>
        <label className={styles.customLabel}>Title:</label>
        <input
          type='text'
          value={title}
          className={styles.customText}
          onChange={(e) => setTitle(e.target.value)}
        />
        <br />
        <label className={styles.customLabel}>Text:</label>
        <textarea
          className={styles.customTextarea}
          value={text}
          onChange={(e) => setText(e.target.value)}
        />
        <br />
        <label className={styles.customLabel}>
          Is Public:
          <input
            type='checkbox'
            checked={isPublic}
            onChange={(e) => setIsPublic(e.target.checked)}
          />
        </label>
        <br />
        <div className='flex items-center justify-between'>
          <Link href={`/knowledge/${knowledge.id}`} className={styles.link}>
            Back
          </Link>
          <button type='submit' className={styles.button}>
            Save
          </button>
        </div>
      </form>
    </div>
  )
}
