import { sizes } from "@/assets";
import { title } from "@/components";
import { IMG } from "@/components/common/IMG";
import { TableOfContent } from "@/components/common/TableOfContent";
import { MDXContent } from "@/components/md/MDXContent";
import { Stage } from "@/generated/graphql";
import { PageProps } from "@/interfaces";
import { env, formatDate } from "@/utils";
import { mdxSerializer } from "@/utils/mdx";
import { getSdk } from "@/utils/sdk";
import { Button } from "@nextui-org/button";
import { Metadata, ResolvedMetadata } from "next";
import Link from "next/link";
import { notFound } from "next/navigation";
import React from "react";
import { FaArrowLeft } from "react-icons/fa";

const BlogPage = async ({ params, searchParams }: PageProps) => {
  const { post } = await getSdk().postBySlug({
    slug: params.slug,
    stage: Stage.Published,
  });

  if (!post) return notFound();
  const { mdx, toc } = await mdxSerializer(post.content);
  return (
    <div>
      <div className="relative aspect-square md:aspect-video">
        <IMG src={post.banner.url} alt={post.title} sizes={sizes.lg} priority />
        <div className="absolute bottom-0 left-0 p-4 bg-gradient-to-t from-background dark:from-dark to-transparent w-full h-full flex flex-col justify-end">
          <div className="max-w-4xl mx-auto">
            <p className="text-sm">{formatDate(post.updatedAt)}</p>
            <h1 className={title()}>{post.title}</h1>
            <div className="rounded-lg italic">{post.summary}</div>
            <Button
              color="primary"
              startContent={<FaArrowLeft />}
              href="/blog"
              as={Link}
              aria-label="Volver"
            >
              Volver
            </Button>
          </div>
        </div>
      </div>
      <div className="max-w-4xl mx-auto p-6">
        <TableOfContent toc={toc} />
        <MDXContent {...mdx} />
      </div>
    </div>
  );
};

export default BlogPage;

export async function generateStaticParams() {
  const { posts } = await getSdk().postsSlug({});

  return posts.map(({ slug }) => ({
    slug,
  }));
}

export const revalidate = env.revalidate;

export async function generateMetadata(
  { params }: PageProps,
  parent: Promise<ResolvedMetadata>
): Promise<Metadata> {
  const { post } = await getSdk().postBySlug({
    slug: params.slug,
    stage: Stage.Published,
  });
  return {
    title: post?.title || "Contedido no encontrado",
    description: post?.summary || "",
    keywords: post?.summary.split(" ") || [],
  };
}
