import { CategoryList, PostList, PostPaginated } from "@/components/post";
import { Stage } from "@/generated/graphql";
import { getSdk } from "@/utils/sdk";
import { env } from "process";
import React from "react";

const BlogsHome = async () => {
  const { firstPosts, categories } = await getSdk().blogsPage({
    stage: Stage.Published,
  });

  const sortedCategories = [...categories]
    .sort((a, b) => b.posts.length - a.posts.length)
    .slice(0, 4);

  return (
    <div className="p-6">
      <PostList posts={firstPosts as any} />
      <CategoryList categories={sortedCategories as any} />
      <PostPaginated />
    </div>
  );
};

export default BlogsHome;

export const revalidate = env.revalidate;
