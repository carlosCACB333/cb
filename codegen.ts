import type { CodegenConfig } from "@graphql-codegen/cli";
import { env } from "./src/utils/env";

const config: CodegenConfig = {
  overwrite: true,
  schema: [
    {
      [env.cms.url]: {
        headers: {
          Authorization: `Bearer ${env.cms.token}`,
        },
      },
    },
  ],
  documents: "src/graphql/**/*.graphql",
  generates: {
    "src/generated/graphql.ts": {
      plugins: [
        "typescript",
        "typescript-operations",
        "typescript-graphql-request",
      ],
      config: {},
    },
    "src/generated/graphql.schema.json": {
      plugins: ["introspection"],
    },
  },
  hooks: {
    afterOneFileWrite: [
      'sed -i -e"s|graphql-request/dist/types.dom|graphql-request/src/types.dom|g"',
    ],
  },
};

export default config;
