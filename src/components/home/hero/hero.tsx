"use client";

import NextLink from "next/link";
import { Button, Link } from "@nextui-org/react";
import dynamic from "next/dynamic";
import { FloatingComponents } from "./floating-components";
import { title, subtitle } from "@/components";
import { useAuthor } from "@/hooks";
import { AiOutlineHeart, AiOutlinePaperClip } from "react-icons/ai";
import { Feature, FeaturesGrid } from "@/components/common/features-grid";
import { FC } from "react";

const BgLooper = dynamic(
  () => import("./bg-looper").then((mod) => mod.BgLooper),
  {
    ssr: false,
  }
);

interface Props {
  features: Feature[];
}

export const Hero: FC<Props> = ({ features }) => {
  const { author } = useAuthor();

  const [
    firstWord = "",
    secondWord = "",
    thirdWord = "",
    fourthWord = "",
    fifthWord = "",
  ] = author.profession!.split(" ");
  return (
    <>
      <section className="relative w-full flex flex-col gap-4 justify-evenly min-h-[calc(100vh-4rem)]">
        <div className="flex flex-nowrap justify-between items-center h-full w-full">
          <div className="flex relative z-20 flex-col gap-6 w-full md:w-1/2 xl:mt-10">
            <h1 className={title({ size: "lg" })}>
              {firstWord} {secondWord} &nbsp;
              <span className={title({ color: "blue", size: "lg" })}>
                {thirdWord} &nbsp;
              </span>
              {fourthWord} &nbsp;
              <span className={title({ color: "blue", size: "lg" })}>
                {fifthWord}
              </span>
              <span className="animate-heartbeat">😎</span>
            </h1>
            <h2
              className={subtitle({
                fullWidth: true,
              })}
            >
              {author.detail}
            </h2>
            <div className="flex flex-col md:flex-row items-center gap-4">
              <Button
                as={NextLink}
                className="w-full md:w-auto"
                color="primary"
                aria-label="boton sobre mí"
                endContent={
                  <AiOutlineHeart
                    className="group-data-[hover=true]:translate-x-0.5 outline-none transition-transform"
                    strokeWidth={2}
                  />
                }
                href="#home-about-me"
                size="lg"
              >
                Sobre mí
              </Button>

              <Button
                fullWidth
                isExternal
                as={Link}
                className="w-full md:w-auto"
                href={author.cv?.url}
                size="lg"
                aria-label="boton descargar cv"
                startContent={
                  <AiOutlinePaperClip
                    className="group-data-[hover=true]:translate-x-0.5 outline-none transition-transform"
                    strokeWidth={2}
                  />
                }
                variant="bordered"
              >
                Descargar CV
              </Button>
            </div>
          </div>
          <FloatingComponents />
        </div>
        <BgLooper />
        <FeaturesGrid features={features} />
      </section>
    </>
  );
};
