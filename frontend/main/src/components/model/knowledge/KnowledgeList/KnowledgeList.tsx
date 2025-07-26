import { useRouter } from 'next/router'

import {
  KnowledgeListItem,
  knowledgeListItemFragment,
} from '@/components/model/knowledge/KnowledgeListItem/KnowledgeListItem'
import { FragmentType } from '@/gql/__generated__'
import { useKnowledgeUsecase } from '@/usecase/knowledge/usecase'

import styles from './KnowledgeList.module.css'

type KnowledgeListProps = {
  knowledgeList: Array<FragmentType<typeof knowledgeListItemFragment>>
}

export const KnowledgeList: React.FC<KnowledgeListProps> = (props) => {
  const knowledgeUsecase = useKnowledgeUsecase()
  const router = useRouter()

  const createKnowledge = async () => {
    try {
      const id = await knowledgeUsecase.createKnowledge({
        input: {
          title: 'New Knowledge Title',
          text: 'New Knowledge Text',
          isPublic: false,
        },
      })
      await router.push(`/knowledge/${id}/edit`)
    } catch (error) {
      console.error(error)
    }
  }

  const handleNewKnowledgeClick = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault()
    createKnowledge().catch((error) => {
      console.error(error)
    })
  }

  return (
    <div className={`${styles.knowledgeList} e2e-model-knowledge-list-all`}>
      <div className={styles.header}>
        <button
          className={`${styles.newKnowledgeButton} e2e-model-knowledge-list-create-button`}
          onClick={handleNewKnowledgeClick}>
          New Knowledge
        </button>
      </div>

      {props.knowledgeList.map((knowledge, index) => (
        <KnowledgeListItem key={index} knowledge={knowledge} />
      ))}
    </div>
  )
}
