"use client";

import { MDXRemoteSerializeResult } from "next-mdx-remote";
import React, { FC } from "react";
import { MDXContent } from "../md/MDXContent";

interface Props {
  content: MDXRemoteSerializeResult;
}
export const PostContent: FC<Props> = ({ content }) => {
  return (
    <div>
      <MDXContent {...content} />
    </div>
  );
};
