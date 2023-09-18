"use client";

import { clsx } from "@nextui-org/shared-utils";

import { useIsMounted } from "@/hooks/use-is-mounted";

export const BgLooper = () => {
  const isMounted = useIsMounted();

  return (
    <div
      className={clsx(
        "absolute -top-20 left-0 w-screen h-[calc(100%+5rem)]",
        "transition-opacity opacity-50 -z-1",
        "bg-left bg-no-repeat bg-[url('/gradients/looper-pattern.svg')]"
      )}
      data-mounted={isMounted}
    />
  );
};
