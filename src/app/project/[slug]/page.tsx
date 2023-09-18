import { PostContent } from "@/components/post/PostContent";
import { Project, Stage } from "@/generated/graphql";
import { PageProps } from "@/interfaces";
import { env } from "@/utils";
import { Metadata, ResolvedMetadata } from "next";
import { notFound } from "next/navigation";
import React from "react";
import { mdxSerializer } from "@/utils/mdx";
import { ProjectCarrousel } from "@/components/project/ProjectCarrousel";
import { getSdk } from "@/utils/sdk";

const ProjectPage = async ({ params, searchParams }: PageProps) => {
  const { project } = await getSdk().projectBySlug({
    slug: params.slug,
    stage: Stage.Published,
  });

  if (!project) {
    notFound();
  }
  const { mdx, toc } = await mdxSerializer(project.detail);
  await new Promise((resolve) => setTimeout(resolve, 10000));
  return (
    <div>
      <div className="relative ">
        <ProjectCarrousel project={project as Project} />
      </div>
      <div className="max-w-4xl mx-auto p-6">
        <PostContent content={mdx} />
      </div>
    </div>
  );
};

export default ProjectPage;

export async function generateStaticParams() {
  const { projects } = await getSdk().ProjectsSlug({});

  return projects.map(({ slug }) => ({
    slug,
  }));
}

export const revalidate = env.revalidate;

export async function generateMetadata(
  { params }: PageProps,
  parent: Promise<ResolvedMetadata>
): Promise<Metadata> {
  const { project } = await getSdk().projectBySlug({
    slug: params.slug,
    stage: Stage.Published,
  });
  return {
    title: project?.title || "Proyecto",
    description: project?.abstract || "",
    keywords: project?.abstract.split(" ") || [],
  };
}
