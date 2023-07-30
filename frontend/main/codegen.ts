import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  schema: 'http://localhost:8080/query',
  generates: {
    'graphql/__generated__/graphql-schema-types.ts': {
      plugins: ['typescript'],
    },
    'src/components/page/': {
      documents: 'src/components/page/**/index.tsx',
      preset: 'near-operation-file',
      plugins: ['typescript-operations', 'typescript-react-apollo'],
      presetConfig: {
        baseTypesPath: '../../../graphql/__generated__/graphql-schema-types.ts',
        folder: '__generated__',
        extension: '.generated.tsx',
        importTypesNamespace: 'SchemaTypes',
      },
    },
    'src/api/': {
      documents: 'src/api/main/mutation/*.mutation.ts',
      preset: 'near-operation-file',
      plugins: ['typescript-operations', 'typescript-react-apollo'],
      presetConfig: {
        baseTypesPath: '../../graphql/__generated__/graphql-schema-types.ts',
        folder: '__generated__',
        extension: '.generated.tsx',
        importTypesNamespace: 'SchemaTypes',
      },
    },
  },
  config: { avoidOptionals: true },
  hooks: { afterAllFileWrite: ['prettier --write'] },
}
export default config