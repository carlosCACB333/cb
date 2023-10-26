"use client";
import { searchProjects } from "@/action";
import { sizes } from "@/assets";
import { Pagination } from "@/components/common/Pagination";
import { Searcher } from "@/components/common/Searcher";
import { ProjectCard } from "@/components/project/ProjectCard";
import { Project } from "@/generated/graphql";
import { usePagination } from "@/hooks";
import { env } from "@/utils";

const ProjectPage = () => {
  const { data, onSearch, onChangePage } = usePagination(searchProjects, "", 4);

  return (
    <div className="container mx-auto mt-20 p-6">
      <div className="max-w-lg mx-auto ">
        <Searcher
          size="lg"
          placeholder="Buscar proyectos..."
          setSearch={onSearch}
        />
      </div>
      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 mt-8">
        {data?.edges.map(({ node }) => (
          <div key={node.id}>
            <ProjectCard project={node as Project} sizes={sizes.md} />
          </div>
        ))}
      </div>

      {data?.pageInfo && (
        <Pagination {...data.pageInfo} onChangePage={onChangePage} />
      )}
    </div>
  );
};

export default ProjectPage;

export const revalidate = env.revalidate;
