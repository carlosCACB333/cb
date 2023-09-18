export * from "./contact";
export * from "./chatpdf";

import { Locale, Author as A } from "@/generated/graphql";
import { MDXRemoteSerializeResult } from "next-mdx-remote";
export interface Author extends Omit<A, "bio"> {
  bio: MDXRemoteSerializeResult;
}
export interface PageProps {
  params: {
    locale: Locale;
    [key: string]: string;
  };
  searchParams: { [key: string]: string };
}

export interface LayoutProps {
  params: {
    locale: Locale;
    [key: string]: string;
  };
  children: React.ReactNode;
}

export interface Toc {
  title: string;
  url: string;
  children: { title: string; url: string }[];
}

export interface SearchResultItem {
  content: string;
  objectID: string;
  url: string;
  type: "lvl1" | "lvl2" | "lvl3";
  hierarchy: {
    lvl1: string | null;
    lvl2?: string | null;
    lvl3?: string | null;
  };
}

export interface Route {
  key: string;
  title: string;
  path: string;
}

export interface Response<T> {
  data: T;
  status: 'completed' | 'error'
  message?: string
  [key: string]: any
}