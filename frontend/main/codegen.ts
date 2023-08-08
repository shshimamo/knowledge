import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  schema: 'http://localhost:8080/query',
  generates: {
    'graphql/__generated__/graphql-schema-types.ts': {
      plugins: ['typescript'],
    },
    'src/components/model/': {
      documents: 'src/components/model/**/!(*.generated).tsx',
      preset: 'near-operation-file',
      plugins: ['typescript-operations', 'typescript-graphql-request'],
      presetConfig: {
        baseTypesPath: '../../../graphql/__generated__/graphql-schema-types.ts',
        folder: '__generated__',
        extension: '.generated.ts',
        importTypesNamespace: 'SchemaTypes',
      },
    },
    'src/components/page/': {
      documents: 'src/components/page/**/!(*.generated).tsx',
      preset: 'near-operation-file',
      plugins: ['typescript-operations', 'typescript-graphql-request'],
      presetConfig: {
        baseTypesPath: '../../../graphql/__generated__/graphql-schema-types.ts',
        folder: '__generated__',
        extension: '.generated.ts',
        importTypesNamespace: 'SchemaTypes',
      },
    },
    'src/api/': {
      documents: [
        'src/api/main/mutation/*.mutation.ts',
        'src/api/main/query/*.query.ts',
      ],
      preset: 'near-operation-file',
      plugins: ['typescript-operations', 'typescript-graphql-request'],
      presetConfig: {
        baseTypesPath: '../../graphql/__generated__/graphql-schema-types.ts',
        folder: '__generated__',
        extension: '.generated.ts',
        importTypesNamespace: 'SchemaTypes',
      },
    },
  },
  config: { avoidOptionals: true },
  hooks: { afterAllFileWrite: ['prettier --write'] },
}
export default config