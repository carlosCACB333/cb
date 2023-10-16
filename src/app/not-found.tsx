"use client";
import React from "react";
import notfound from "@/assets/img/not-found.svg";
import Image from "next/image";
import { Button, Link } from "@nextui-org/react";

export const NotFound = () => {
  return (
    <div className="flex flex-col items-center justify-center min-h-[80vh] w-full">
      <Image src={notfound} alt="not found" />
      <h2>
        <span className="text-2xl font-bold">Oops, este recurso no existe</span>
      </h2>
      <br />
      <Button
        as={Link}
        href="/"
        color="primary"
        aria-label="voler a la pagina de inicio"
      >
        Volver a la página de inicio
      </Button>
    </div>
  );
};

export default NotFound;
