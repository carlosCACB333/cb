"use client";

import { Author } from "@/generated/graphql";
import { createContext } from "react";

interface AuthorContextProps {
  author: Author;
}

export const AuthorContext = createContext<AuthorContextProps>(
  {} as AuthorContextProps
);

export const AuthorProvider = ({
  author,
  children,
}: {
  author: Author;
  children: React.ReactNode;
}) => (
  <AuthorContext.Provider value={{ author }}>{children}</AuthorContext.Provider>
);
