"use server";
import { GraphQLClient } from "graphql-request";
import { env } from "./env";
import { getSdk as sdk } from "@/generated/graphql";

export const getSdk = () => {
  const client = new GraphQLClient(env.cms.url, {
    headers: {
      Authorization: `Bearer ${env.cms.token}`,
    },
  });
  return sdk(client);
};
