import { KnowledgeEdit } from '@/components/model/knowledge/KnowledgeEdit/KnowledgeEdit'
import { Loading } from '@/components/ui/global/Loading'
import { useKnowledgeItemForEdit } from '@/usecase/knowledge/reader'

type KnowledgeEditRootProps = {
  id: string
}

export const KnowledgeEditRoot: React.FC<KnowledgeEditRootProps> = (props) => {
  const { data, isLoading } = useKnowledgeItemForEdit({ id: props.id })

  if (isLoading) return <Loading />

  return (
    <div>{data && data.knowledge && <KnowledgeEdit knowledge={data.knowledge} />}</div>
  )
}
