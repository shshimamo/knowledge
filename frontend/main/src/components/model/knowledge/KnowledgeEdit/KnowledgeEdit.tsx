/* KnowledgeEdit Model */
import React, { useState } from 'react'
import Link from 'next/link'
import { useRouter } from 'next/router'
import styles from './KnowledgeEdit.module.css';
import { FragmentType, graphql, useFragment } from '@/gql/__generated__'
import { useKnowledgeUsecase } from '@/usecase/knowledge/usecase'

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
  const router = useRouter();
  const knowledge = useFragment(knowledgeEditFragment, props.knowledge);
  const [title, setTitle] = useState(knowledge.title);
  const [text, setText] = useState(knowledge.text);
  const [isPublic, setIsPublic] = useState(knowledge.isPublic);
  const { updateKnowledge } = useKnowledgeUsecase();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      await updateKnowledge({
        id: knowledge.id,
        input: { title, text, isPublic }
      });
      await router.push(`/knowledge/${knowledge.id}`);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className={styles.customContainer}>
      <form onSubmit={handleSubmit} className={styles.form}>
        <label className={styles.customLabel}>
          Title:
        </label>
        <input type="text" value={title} className={styles.customText}
               onChange={e => setTitle(e.target.value)}/>
        <br />
        <label className={styles.customLabel}>
          Text:
        </label>
        <textarea
          className={styles.customTextarea}
          value={text}
          onChange={e => setText(e.target.value)}
        />
        <br />
        <label className={styles.customLabel}>
          Is Public:
          <input type="checkbox" checked={isPublic} onChange={e => setIsPublic(e.target.checked)} />
        </label>
        <br />
        <div className="flex items-center justify-between">
          <Link href={`/knowledge/${knowledge.id}`} className={styles.link}>
            Back
          </Link>
          <button type="submit" className={styles.button}>Save</button>
        </div>
      </form>
    </div>
  );
};