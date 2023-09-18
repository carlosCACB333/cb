"use client";
import { Post } from "@/generated/graphql";
import React, { FC } from "react";
import { PostCard } from "./PostCard";
interface Props {
  posts: Post[];
}
export const PostList: FC<Props> = ({ posts }) => {
  1;

  return (
    <div className="grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
      {posts.map((item, idx) => (
        <PostCard key={item.id} post={item} idx={idx} />
      ))}
    </div>
  );
};
