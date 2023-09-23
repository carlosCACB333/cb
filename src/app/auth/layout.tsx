import { Footer } from "@/components/common/footer";
import { LayoutProps } from "@/interfaces";

export default async function AuthLayout({ children }: LayoutProps) {
  return (
    <>
      <main className="container mx-auto max-w-7xl flex justify-center items-center min-h-[calc(100vh-8rem)] p-6">
        {children}
      </main>
      <Footer />
    </>
  );
}
