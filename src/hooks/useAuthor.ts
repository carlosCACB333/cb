"use client";

import { AuthorContext } from "@/context";
import { useContext } from "react";

export const useAuthor = () => {
  const author = useContext(AuthorContext);
  return {
    ...author,
  };
};
