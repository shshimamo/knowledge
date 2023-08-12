import { useKnowledgeDetail } from '@/usecase/knowledge/reader'
import { KnowledgeDetail } from '@/components/model/knowledge/KnowledgeDetail/KnowledgeDetail'
import { Loading } from '@/components/page/global/Loading'

type KnowledgeDetailTopProps = {
  id: string
}

export const KnowledgeDetailTop: React.FC<KnowledgeDetailTopProps> = (props) => {
  const { data, error, isLoading } = useKnowledgeDetail({ id: props.id })

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
