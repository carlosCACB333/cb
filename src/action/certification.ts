"use server";

import { CertificationConnection, Stage } from "@/generated/graphql";
import { SearchFunction } from "./types";
import { getSdk } from "@/utils/sdk";

export const searchCertification: SearchFunction<
  CertificationConnection
> = async (keyword: string, first: number, skip: number, stage: Stage) => {
  const { certificationsConnection } = await getSdk().searchCertifications({
    search: keyword,
    first,
    skip,
    stage,
  });
  return certificationsConnection as CertificationConnection;
};
