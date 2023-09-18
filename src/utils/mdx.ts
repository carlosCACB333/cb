import { serialize } from "next-mdx-remote/serialize";
import rehypeSlug from "rehype-slug";
import remarkUnwrapImages from "remark-unwrap-images";
import remarkMath from "remark-math";
import { Toc } from "@/interfaces";
import remarkGfm from "remark-gfm";
import rehypeKatex from "rehype-katex";
export const mdxSerializer = async (source: string) => {
  const toc: Toc[] = [
    {
      title: "Inicio",
      url: "#",
      children: [],
    },
  ];
  const mdx = await serialize(source, {
    parseFrontmatter: true,
    mdxOptions: {
      remarkPlugins: [remarkUnwrapImages, remarkMath, remarkGfm],
      rehypePlugins: [
        rehypeSlug,
        rehypeKatex as any,
        () => {
          return (tree) => {
            tree.children.forEach((node: any) => {
              if (node.type === "element" && node.tagName === "h2") {
                toc.push({
                  title: node.children[0].value as string,
                  url: `#${node.properties.id}`,
                  children: [],
                });
              }
              if (node.type === "element" && node.tagName === "h3") {
                toc[toc.length - 1].children.push({
                  title: node.children[0].value as string,
                  url: `#${node.properties.id}`,
                });
              }
            });
          };
        },
      ],
    },
  });

  return { mdx, toc };
};
