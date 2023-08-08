import gql from 'graphql-tag'
import React, { useEffect, useState } from 'react';
import { gqlClient } from '@/api/main/gqlClient'
import { getSdk } from './__generated__/index.generated'
import { KnowledgeList } from '@/components/model/knowledge/KnowledgeList'
import {
  KnowledgeItemFragment,
  KnowledgeItemFragmentDoc
} from '@/components/model/knowledge/__generated__/KnowledgeItem.generated'

gql`
    query MyKnowledgeList {
        currentUser {
            knowledgeList(first: 20) {
                ...KnowledgeItem
            }
        }
    }
    ${KnowledgeItemFragmentDoc}
`

export const Knowledges = () => {
  const sdk = getSdk(gqlClient);
  const [knowledgeList, setKnowledgeList] = useState<Array<KnowledgeItemFragment>>([]);

  useEffect(() => {
    sdk.MyKnowledgeList().then(response => {
      const list: Array<KnowledgeItemFragment> = response.currentUser?.knowledgeList;
      setKnowledgeList(list || []);
    });
  }, [sdk]);

  return (
    <div>
      <KnowledgeList items={knowledgeList} />
    </div>
  );
};