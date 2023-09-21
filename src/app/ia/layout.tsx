import { LayoutProps } from "@/interfaces";
import React from "react";

const IALayout = ({ children }: LayoutProps) => {
  return <div className="h-[calc(100vh-8rem)] grid">{children}</div>;
};

export default IALayout;
