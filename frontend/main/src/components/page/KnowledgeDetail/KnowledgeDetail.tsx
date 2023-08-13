import { useKnowledgeItemForDetail } from '@/usecase/knowledge/reader'
import { KnowledgeDetail } from '@/components/model/knowledge/KnowledgeDetail/KnowledgeDetail'
import { Loading } from '@/components/page/global/Loading'

type KnowledgeDetailRootProps = {
  id: string
}

export const KnowledgeDetailRoot: React.FC<KnowledgeDetailRootProps> = (props) => {
  const { data, error, isLoading } = useKnowledgeItemForDetail({ id: props.id })

  if (error) return <div>failed to load</div>
  if (isLoading) return <Loading />

  return (
    <div>
      {data && data.knowledge && (
        <KnowledgeDetail knowledge={data.knowledge} />
      )}
    </div>
  )
}
