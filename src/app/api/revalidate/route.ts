import { NextResponse } from "next/server";
import fs from "fs";
import { SearchResultItem } from "@/interfaces";
import { getSdk } from "@/utils/sdk";
import { execSync } from "child_process";
import { revalidatePath } from "next/cache";
import { formatDate } from "@/utils";

const defaultSearchResult: SearchResultItem[] = [
  {
    content: "Lista de blogs",
    objectID: "blog",
    type: "lvl1",
    url: "/blog",
    hierarchy: {
      lvl1: "Blog",
    },
  },
  {
    content: "Lista de proyectos",
    objectID: "projects",
    type: "lvl1",
    url: "/projects",
    hierarchy: {
      lvl1: "Projects",
    },
  },
  {
    content: "Inicio",
    objectID: "home",
    type: "lvl1",
    url: "/",
    hierarchy: {
      lvl1: "Home",
    },
  },
];

export async function POST() {
  const { posts, projects } = await getSdk().getSearchMeta({});

  //   Rebuild search.json
  const searchResult: SearchResultItem[] = [...defaultSearchResult];
  posts.forEach((post) => {
    searchResult.push({
      content: post.title,
      objectID: post.id,
      type: "lvl1",
      url: `/blog/${post.slug}`,
      hierarchy: {
        lvl1: "Blog",
      },
    });
  });
  projects.forEach((project) => {
    searchResult.push({
      content: project.title,
      objectID: project.id,
      type: "lvl1",
      url: `/project/${project.slug}`,
      hierarchy: {
        lvl1: "Projects",
      },
    });
  });

  fs.writeFileSync("./public/search-meta.json", JSON.stringify(searchResult));

  //  Rebuild sitemap.xml
  execSync("yarn run postbuild");

  // Revalidate all pages
  revalidatePath("/");

  return NextResponse.json({
    ok: true,
    now: formatDate(new Date().toString()),
  });
}
