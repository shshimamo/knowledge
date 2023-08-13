import { useKnowledgeItemForEdit, useKnowledgeItemForDetail } from '@/usecase/knowledge/reader'
import { Loading } from '@/components/page/global/Loading'
import { KnowledgeEdit } from '@/components/model/knowledge/KnowledgeEdit/KnowledgeEdit'

type KnowledgeEditRootProps = {
  id: string
}

export const KnowledgeEditRoot: React.FC<KnowledgeEditRootProps> = (props) => {
  const { data, error, isLoading } = useKnowledgeItemForEdit({ id: props.id })

  if (error) return <div>failed to load</div>
  if (isLoading) return <Loading />

  return (
    <div>
      {data && data.knowledge && (
        <KnowledgeEdit knowledge={data.knowledge} />
      )}
    </div>
  )
}