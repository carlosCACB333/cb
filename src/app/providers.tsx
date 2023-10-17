"use client";
import "react-toastify/dist/ReactToastify.css";
import { NextUIProvider } from "@nextui-org/react";
import { ThemeProvider as NextThemesProvider, useTheme } from "next-themes";
import { ThemeProviderProps } from "next-themes/dist/types";
import { AuthorProvider } from "@/context";
import { ToastContainer } from "react-toastify";
import { ReactNode } from "react";
import { Author } from "@/generated/graphql";

export interface ProvidersProps {
  children: ReactNode;
  themeProps?: ThemeProviderProps;
  author: Author;
}

export function Providers({ children, themeProps, author }: ProvidersProps) {
  return (
    <NextUIProvider>
      <AuthorProvider author={author as any}>
        <NextThemesProvider {...themeProps}>
          <ProvidersChild>{children}</ProvidersChild>
        </NextThemesProvider>
      </AuthorProvider>
    </NextUIProvider>
  );
}

const ProvidersChild = ({ children }: { children: ReactNode }) => {
  const { theme } = useTheme();
  return (
    <>
      <ToastContainer
        theme={theme as any}
        toastStyle={{
          background: theme === "dark" ? "#0f121a" : "#eaf5ff",
        }}
      />
      {children}
    </>
  );
};
