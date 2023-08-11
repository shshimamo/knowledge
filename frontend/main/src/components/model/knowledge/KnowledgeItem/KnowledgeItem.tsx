import React from 'react'
import { FragmentType, graphql, useFragment } from '@/gql/__generated__'

export const knowledgeItemFragment = graphql(/* GraphQL */ `
    fragment KnowledgeItem on Knowledge {
        title
        text
    }
`);

type KnowledgeItemProps = {
  knowledge: FragmentType<typeof knowledgeItemFragment>
};

export const KnowledgeItem: React.FC<KnowledgeItemProps> = (props) => {
  const knowledge = useFragment(knowledgeItemFragment, props.knowledge)

  return (
      <div>
          <h3>{knowledge.title}</h3>
          <p>{knowledge.text}</p>
      </div>
    );
};
