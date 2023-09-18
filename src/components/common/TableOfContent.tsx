"use client";
import { Toc } from "@/interfaces";
import { Link } from "@nextui-org/react";
import { Card, CardBody, CardHeader } from "@nextui-org/react";
import React from "react";

interface Props {
  toc: Toc[];
}

export const TableOfContent = ({ toc }: Props) => {
  return (
    <Card className="mx-auto max-w-md mb-4">
      <CardHeader>
        <h3 className="text-center text-2xl font-bold text-primary">
          Contenido
        </h3>
      </CardHeader>
      <CardBody>
        <ul>
          {toc.map((item) => (
            <li key={item.title}>
              <LinkItem title={item.title} url={item.url} />
              <ol className="ps-4">
                {item.children.map((subitem) => (
                  <li key={subitem.title}>
                    <LinkItem title={subitem.title} url={subitem.url} />
                  </li>
                ))}
              </ol>
            </li>
          ))}
        </ul>
      </CardBody>
    </Card>
  );
};

export const LinkItem = ({ title, url }: { title: string; url: string }) => {
  return (
    <Link
      className="font-semibold ps-2 relative ml-[0.4rem] before:content-[''] before:absolute before:left-[-0.4rem] before:top-[.5rem] before:w-[0.4rem] before:h-[0.4rem] before:bg-primary before:rounded-full"
      color="foreground"
      size="sm"
      href={url}
      aria-label={title}
    >
      {title}
    </Link>
  );
};
