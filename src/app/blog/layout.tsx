import { subtitle } from "@/components";
import { IMG } from "@/components/common/IMG";
import { Footer } from "@/components/common/footer";
import { Stage } from "@/generated/graphql";
import { LayoutProps } from "@/interfaces";
import { formatDate } from "@/utils";
import { getSdk } from "@/utils/sdk";
import { Link } from "@nextui-org/link";
import { Metadata } from "next";
import NextLink from "next/link";
import React from "react";

const Layout = async ({ children }: LayoutProps) => {
  const { firstPosts, categories, skills } = await getSdk().blogsLayout({
    stage: Stage.Published,
  });

  return (
    <main className="flex flex-col lg:flex-row gap-8 lg:gap-16 container m-auto">
      <div className="flex-1">
        {children}
        <br />
        <Footer />
      </div>
      <aside className="w-full md:w-96 sticky top-16 scroll pr-2 lg:max-h-[calc(100vh-6rem)] p-6">
        <h3 className={subtitle({ class: "text-foreground" })}>Categorías</h3>
        <ol>
          {categories.map((category) => (
            <li
              key={category.id}
              className="font-semibold ps-2 relative ml-[0.4rem] before:content-[''] before:absolute before:left-[-0.4rem] before:top-[.7rem] before:w-[0.4rem] before:h-[0.4rem] before:bg-primary before:rounded-full "
            >
              <Link
                color="foreground"
                size="sm"
                href={`?category=${category.slug}`}
                as={NextLink}
                aria-label={category.name}
              >
                {category.name}
              </Link>
            </li>
          ))}
        </ol>
        <h3 className={subtitle({ class: "text-foreground" })}>
          Posts recientes
        </h3>
        <div className="flex flex-col gap-4">
          {firstPosts.map((post) => (
            <div key={post.id} className="flex flex-row">
              <div className="w-32 min-h-[4rem]">
                <IMG
                  src={post.banner.url}
                  alt={post.title}
                  sizes="100px"
                  className="rounded-md"
                />
              </div>

              <div className="flex flex-1 flex-col ml-4">
                <Link
                  href={`/blog/${post.slug}`}
                  className="self-start"
                  as={NextLink}
                  aria-label={post.title}
                >
                  {post.title}
                </Link>
                <p className="text-tiny">{formatDate(post.updatedAt)}</p>
              </div>
            </div>
          ))}
        </div>
        <h3 className={subtitle({ class: "text-foreground" })}>Etiquetas</h3>
        <div className="flex flex-row flex-wrap gap-2 max-w-full">
          {skills.map((skill) => (
            <div key={skill.id} className="bg-secondary rounded-lg px-3 py-1">
              {skill.name}
            </div>
          ))}
        </div>
      </aside>
    </main>
  );
};

export default Layout;

export async function generateMetadata(): Promise<Metadata> {
  return {
    title: {
      default: "Blog",
      template: `%s`,
    },
  };
}
