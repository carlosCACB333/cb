import { Footer } from "@/components/common/footer";
import { LayoutProps } from "@/interfaces";
import { Metadata } from "next";

const Layout = async ({ children }: LayoutProps) => {
  return (
    <>
      {children}
      <br />
      <Footer />
    </>
  );
};

export default Layout;

export async function generateMetadata(): Promise<Metadata> {
  return {
    title: "Certificados",
  };
}
