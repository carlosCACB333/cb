import "@/styles/globals.css";
import "katex/dist/katex.min.css";
import { Metadata } from "next";
import { clsx } from "clsx";
import { Providers } from "./providers";
import routes from "@/config/routes.json";
import { siteConfig } from "@/config/site";
import { fontRoboto } from "@/config/fonts";
import { __PROD__ } from "@/utils";
import { getAuthor, getCookie } from "@/action";
import { Locale } from "@/generated/graphql";
import { Navbar } from "@/components/common/navbar";
import { Footer } from "@/components/common/footer";
import { Cmdk } from "@/components/common/cmdk";
import { serialize } from "next-mdx-remote/serialize";
import { ClerkProvider } from "@clerk/nextjs";
import { dark } from "@clerk/themes";

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const author = await getAuthor(Locale.Es);
  const bio = await serialize(author?.bio?.toString() || "");
  const defaultTheme = getCookie("theme", "dark");
  const isDark = defaultTheme === "dark";
  return (
    <html suppressHydrationWarning dir="ltr" lang="es">
      <head />
      <body
        className={clsx(
          "scroll overflow-x-clip",
          "min-h-screen bg-background antialiased",
          fontRoboto.className
        )}
      >
        <Providers
          themeProps={{ attribute: "class", defaultTheme }}
          author={{ ...author, bio } as any}
        >
          <ClerkProvider
            appearance={{
              variables: {
                colorPrimary: "#6b9aec",
                colorText: isDark ? "#d4ddfb" : "#000001",
                colorTextSecondary: isDark ? "#bdc8f0" : "#404152",
                colorBackground: isDark ? "#0f121a" : "#eaf5ff",
                colorInputBackground: isDark ? "#0f121a" : "#eaf5ff",
                colorInputText: isDark ? "#d4ddfb" : "#000001",
                colorAlphaShade: "#6b9aec",
              },
              baseTheme: isDark ? dark : undefined,
            }}
          >
            <div className="relative flex flex-col" id="app-container">
              <Navbar
                mobileRoutes={routes.mobileRoutes}
                routes={routes.routes}
              />
              {children}
            </div>
            <Cmdk />
          </ClerkProvider>
        </Providers>
        {/* {__PROD__ && <Analytics />} */}
      </body>
    </html>
  );
}

export async function generateMetadata(): Promise<Metadata> {
  //
  const author = await getAuthor(Locale.Es);
  const authorName = author?.firstName + " " + author?.lastName;
  return {
    title: {
      default: authorName,
      template: `${authorName} | %s`,
    },
    description: author?.bio?.toString(),
    authors: [
      {
        name: authorName,
        url: siteConfig.siteUrl,
      },
    ],
    keywords: author?.keywords || [],
    creator: authorName,

    themeColor: [
      { media: "(prefers-color-scheme: light)", color: "white" },
      { media: "(prefers-color-scheme: dark)", color: "#07090e" },
    ],
    icons: {
      icon: "/favicon.ico",
      shortcut: "/favicon-32x32.png",
      apple: "/apple-touch-icon.png",
    },
    manifest: "/manifest.json",
    viewport:
      "viewport-fit=cover, width=device-width, initial-scale=1, shrink-to-fit=no",
  };
}
