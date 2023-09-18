"use client";

import { Button } from "@nextui-org/react";
import { useEffect } from "react";

export default function Error({
  error,
  reset,
}: {
  error: Error;
  reset: () => void;
}) {
  useEffect(() => {
    console.error(error);
  }, [error]);

  return (
    <div className="flex flex-col items-center justify-center h-full w-full gap-4 text-center min-h-[80vh]">
      <h2 className="text-2xl font-bold">Oops,Algo salió mal</h2>
      <Button
        aria-label="Intentar de nuevo"
        color="primary"
        onClick={() => reset()}
      >
        Intentar de nuevo
      </Button>
    </div>
  );
}
