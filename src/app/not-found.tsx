"use client";
import React from "react";
import notfound from "@/assets/img/not-found.svg";
import Image from "next/image";
import { Button, Link } from "@nextui-org/react";

export const notFound = () => {
  return (
    <div className="flex flex-col items-center justify-center min-h-[80vh] ">
      <Image src={notfound} alt="not found" />

      <h2>
        <span className="text-2xl font-bold">Oops,Algo salió mal</span>
      </h2>
      <br />
      <Button as={Link} href="/" color="primary" aria-label="ir a inicio">
        Volver a la página de inicio
      </Button>
    </div>
  );
};

export default notFound;
