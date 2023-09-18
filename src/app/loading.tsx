"use client";
import { CircularProgress } from "@nextui-org/react";

export default function Loading() {
  return (
    <div className="flex flex-col items-center justify-center h-[80vh] p-8">
      <CircularProgress
        size="lg"
        color="primary"
        label="Cargando contenido..."
      />
    </div>
  );
}
