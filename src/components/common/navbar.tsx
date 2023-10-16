"use client";

import { useState, FC, ReactNode } from "react";
import {
  link,
  Navbar as NextUINavbar,
  NavbarContent,
  NavbarMenu,
  NavbarMenuToggle,
  NavbarBrand,
  NavbarItem,
  Link,
  Button,
  Kbd,
  NavbarMenuItem,
} from "@nextui-org/react";
import { isAppleDevice } from "@react-aria/utils";
import { clsx } from "@nextui-org/shared-utils";
import NextLink from "next/link";
import { usePathname } from "next/navigation";
import { includes } from "lodash";
import { useEffect } from "react";
import { AiOutlineSearch } from "react-icons/ai";
import { ThemeSwitch } from "./theme-switch";
import { Icon } from "./icon";
import { useAuthor } from "@/hooks";
import { useCmdkStore } from "./cmdk";
import { Route } from "@/interfaces";
import { FiLogIn } from "react-icons/fi";
import { UserButton, useAuth } from "@clerk/nextjs";

export interface NavbarProps {
  routes: Route[];
  mobileRoutes: Route[];
  children?: ReactNode;
}

export const Navbar: FC<NavbarProps> = ({
  children,
  routes,
  mobileRoutes = [],
}) => {
  const [isMenuOpen, setIsMenuOpen] = useState<boolean | undefined>(false);
  const [commandKey, setCommandKey] = useState<"ctrl" | "command">("command");
  const { author } = useAuthor();
  const pathname = usePathname();
  const cmdkStore = useCmdkStore();
  const { userId } = useAuth();

  const isAuth = !!userId;
  useEffect(() => {
    setIsMenuOpen(false);
  }, [pathname]);

  useEffect(() => {
    setCommandKey(isAppleDevice() ? "command" : "ctrl");
  }, []);

  const searchButton = (
    <Button
      aria-label="Buscar"
      className="text-sm font-normal"
      endContent={
        <Kbd className="py-0.5 px-2 bg-background" keys={commandKey}>
          K
        </Kbd>
      }
      startContent={
        <AiOutlineSearch
          className="text-base pointer-events-none flex-shrink-0"
          size={18}
          strokeWidth={2}
        />
      }
      onPress={() => cmdkStore.onOpen()}
    >
      Buscar...
    </Button>
  );

  const navLinkClasses = clsx(
    link({ color: "foreground" }),
    "data-[active=true]:text-primary"
  );

  return (
    <NextUINavbar
      className={clsx({
        "z-[100001]": isMenuOpen,
      })}
      isMenuOpen={isMenuOpen}
      maxWidth="xl"
      position="sticky"
      onMenuOpenChange={setIsMenuOpen}
    >
      <NavbarContent className="basis-1/5 sm:basis-full" justify="start">
        <NavbarBrand as="li" className="gap-3 max-w-fit">
          <NextLink
            aria-label="Home"
            className="flex justify-start items-center gap-2 tap-highlight-transparent transition-opacity active:opacity-50"
            href="/"
          >
            <Icon name="logo" className="fill-primary" height={40} width={40} />
          </NextLink>
        </NavbarBrand>
        <ol className="hidden sm:flex gap-4 justify-start items-center">
          {routes.map((route) => (
            <NavbarItem key={route.key}>
              <NextLink
                className={navLinkClasses}
                color="foreground"
                data-active={includes(pathname, route.path)}
                href={route.path}
                aria-label={route.title}
              >
                {route.title}
              </NextLink>
            </NavbarItem>
          ))}
        </ol>
      </NavbarContent>

      <NavbarContent className="flex w-full gap-2" justify="end">
        {author.github && (
          <NavbarItem className="flex h-full items-center">
            <Link
              isExternal
              aria-label="Github"
              className="p-1"
              href={author.github}
            >
              <Icon name="git" className="text-foreground" />
            </Link>
          </NavbarItem>
        )}
        {author.linkedin && (
          <NavbarItem className="flex h-full items-center">
            <Link
              isExternal
              aria-label="Linkedin"
              className="p-1"
              href={author.linkedin}
            >
              <Icon name="linkedin" className="text-foreground" />
            </Link>
          </NavbarItem>
        )}
        <NavbarItem className="flex h-full items-center">
          <ThemeSwitch />
        </NavbarItem>
        <NavbarItem className="flex h-full items-center sm:hidden">
          <Button
            size="sm"
            isIconOnly
            className="bg-transparent"
            aria-label="Buscar"
            onPress={() => cmdkStore.onOpen()}
          >
            <AiOutlineSearch className="mt-px text-foreground" size={20} />
          </Button>
        </NavbarItem>

        <NavbarItem className="hidden sm:flex">{searchButton}</NavbarItem>
        {!isAuth && (
          <Button
            endContent={<FiLogIn />}
            color="primary"
            className="hidden sm:flex"
            href="/auth/sign-in/"
            as={NextLink}
            aria-label="botón de login"
          >
            Login
          </Button>
        )}
        <UserButton userProfileUrl="/me" />
        <NavbarItem className="w-10 h-full sm:hidden">
          <NavbarMenuToggle
            aria-label={isMenuOpen ? "Close menu" : "Open menu"}
            className="w-full h-full pt-1"
          />
        </NavbarItem>
      </NavbarContent>

      <NavbarMenu>
        <div className="mx-4 mt-2 flex flex-col gap-2">
          {mobileRoutes.map((item) => (
            <NavbarMenuItem key={item.key}>
              <NextLink
                href={item.path}
                data-active={includes(pathname, item.key)}
                className={clsx(
                  link({ color: "foreground" }),
                  "data-[active=true]:text-primary"
                )}
                aria-label={item.title}
              >
                {item.title}
              </NextLink>
            </NavbarMenuItem>
          ))}
        </div>
        {children}
      </NavbarMenu>
    </NextUINavbar>
  );
};
