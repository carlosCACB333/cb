"use server";

import { ProjectConnection, Stage } from "@/generated/graphql";
import { SearchFunction } from "./types";
import { getSdk } from "@/utils/sdk";

export const searchProjects: SearchFunction<ProjectConnection> = async (
  keyword: string,
  first: number,
  skip: number,
  stage: Stage
) => {
  const { projectsConnection } = await getSdk().searchProjects({
    search: keyword,
    first,
    skip,
    stage,
  });
  return projectsConnection as ProjectConnection;
};
