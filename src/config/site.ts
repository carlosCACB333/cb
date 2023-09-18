import { env } from "@/utils";

export type SiteConfig = typeof siteConfig;

export const siteConfig = {
  name: "carloscb - portfolio",
  description: "Portfolio profesional de Carlos Castillo Blas",
  author: "Carlos Castillo Blas",
  email: env.author.email,
  siteUrl: env.site.url,
  creator: "Carlos Castillo Blas",
};
