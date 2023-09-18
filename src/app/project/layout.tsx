import { LayoutProps } from "@/interfaces";
import { Metadata } from "next";

const Layout = async ({ children }: LayoutProps) => {
  return <>{children}</>;
};

export default Layout;

export async function generateMetadata(): Promise<Metadata> {
  return {
    title: "Proyectos",
  };
}
