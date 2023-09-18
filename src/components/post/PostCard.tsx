"use client";
import { formatDate } from "@/utils";
import { Card, CardHeader, CardFooter } from "@nextui-org/react";
import clsx from "clsx";
import React from "react";
import { Post } from "@/generated/graphql";
import Link from "next/link";
import { Icon } from "../common/icon";
import { IMG } from "../common/IMG";
import { sizes } from "@/assets";
interface Props {
  post: Post;
  idx?: number;
}
export const PostCard = ({ post, idx }: Props) => {
  return (
    <Card
      as={Link}
      href={`/blog/${post.slug}`}
      className={clsx("relative", {
        "md:col-span-2 md:row-span-2": idx === 0,
      })}
      key={post.id}
      aria-label={post.title}
    >
      <CardHeader className="absolute z-10 top-1 justify-end">
        {
          post.categories?.map((cat) => (
            <div key={cat.id} className="bg-content1/60 rounded-full p-2">
              <Icon
                className="fill-foreground"
                name={cat.icon as any}
                style={{
                  height: 20,
                  width: 20,
                }}
              />
            </div>
          )) as any
        }
      </CardHeader>
      <div className="aspect-square">
        <IMG
          alt={post.title}
          className="z-0 w-full h-full object-cover "
          src={post.banner.url}
          sizes={sizes.sm}
          priority={idx === 0}
        />
      </div>
      <CardFooter className="absolute bg-content1/90 bottom-0 z-10 gap-4">
        <div className="flex flex-col">
          <p className="text-tiny">{formatDate(post.updatedAt)}</p>
          <p className="font-bold">{post.title}</p>
          <p className={clsx("text-sm", `line-clamp-3`)}>{post.summary}</p>
        </div>
      </CardFooter>
    </Card>
  );
};
