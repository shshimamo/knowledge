import gql from 'graphql-tag'
import { KnowledgeItemFragment } from './__generated__/KnowledgeItem.generated'
import React from 'react'

gql`
    fragment KnowledgeItem on Knowledge {
        title
        text
    }
`;

type KnowledgeItemProps = {
    knowledge: KnowledgeItemFragment;
};

export const KnowledgeItem: React.FC<KnowledgeItemProps> = ({ knowledge }) => {
    return (
      <div>
          <h3>{knowledge.title}</h3>
          <p>{knowledge.text}</p>
      </div>
    );
};
