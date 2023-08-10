import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  schema: 'http://localhost:8080/query',
  documents: ['src/**/*.tsx', 'src/**/*.ts'],
  ignoreNoDocuments: true, // for better experience with the watcher
  generates: {
    "./src/gql/__generated__/": {
      preset: "client",
      plugins: [
        {
          // Custom Scalar の branded type 定義
          add: {
            content: `export type DateString = string & { readonly __brand: unique symbol }`,
          },
        },
      ],
      config: {
        strictScalars: true,
        useTypeImports: true,
        skipTypename: true,
        arrayInputCoercion: true,
        avoidOptionals: {
          field: true,
          inputValue: false,
          object: true,
          defaultValue: false,
        },
        scalars: {
          Date: "DateString",
        },
        enumsAsTypes: true,
      },
    },
  },
}
export default config