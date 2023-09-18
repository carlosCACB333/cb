import { Project } from "@/generated/graphql";
import { Button, Card, CardFooter, CardHeader, Link } from "@nextui-org/react";
import React from "react";
import { Icon } from "../common/icon";
import NextLink from "next/link";
import { IMG } from "../common/IMG";
import { sizes as c } from "@/assets";

interface Props {
  project: Project;
  sizes?: string;
}

export const ProjectCard = ({ project, sizes = c.xs }: Props) => {
  return (
    <Card
      isFooterBlurred
      className="relative aspect-video col-span-12 sm:col-span-7"
    >
      <CardHeader className="absolute z-10 top-1 justify-end">
        {
          project.skills?.map((skill) => (
            <div key={skill.id} className="bg-content1/60 rounded-full p-2">
              <Icon
                className="fill-foreground"
                name={skill.icon as any}
                style={{
                  height: 20,
                  width: 20,
                }}
              />
            </div>
          )) as any
        }
      </CardHeader>
      <div className="aspect-video">
        <IMG
          fill
          alt={project.title}
          className="object-cover"
          sizes={sizes}
          src={project.pictures[0].url}
        />
      </div>
      <CardFooter className="absolute bg-warning-50/60 bottom-0 z-10 gap-4">
        <div className="flex flex-grow gap-2 projects-center">
          <div className="flex flex-col">
            <p className="font-bold">{project.title}</p>
            <p className="line-clamp-2">{project.abstract}</p>
          </div>
        </div>
        <div className="flex gap-1">
          {project.gitHub && (
            <Button
              color="warning"
              variant="ghost"
              as={Link}
              href={project.gitHub}
              isIconOnly
              isExternal
              aria-label="Ver en GitHub"
            >
              <Icon
                name="git"
                style={{
                  height: 20,
                  width: 20,
                }}
              />
            </Button>
          )}
          {project.webSide && (
            <Button
              color="warning"
              variant="ghost"
              as={Link}
              href={project.webSide}
              isIconOnly
              isExternal
              aria-label="Ver en la web"
            >
              <Icon
                name="world"
                style={{
                  height: 20,
                  width: 20,
                }}
              />
            </Button>
          )}

          <Button
            color="warning"
            variant="ghost"
            as={NextLink}
            href={`/project/${project.slug}`}
            isIconOnly
            aria-label="Ver detalles"
          >
            <Icon
              name="right"
              style={{
                height: 20,
                width: 20,
              }}
            />
          </Button>
        </div>
      </CardFooter>
    </Card>
  );
};
