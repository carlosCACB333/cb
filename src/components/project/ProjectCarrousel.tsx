"use client";
import "swiper/css";
import { Project } from "@/generated/graphql";
import React, { FC } from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import { Autoplay, EffectCreative } from "swiper/modules";
import { sizes } from "@/assets";
import { Button } from "@nextui-org/button";
import { FaArrowLeft } from "react-icons/fa";
import { useRouter } from "next/navigation";
import { IMG } from "../common/IMG";
import { Link } from "@nextui-org/react";
import { Icon } from "../common/icon";
import clsx from "clsx";

interface Props {
  project: Project;
}
export const ProjectCarrousel: FC<Props> = ({ project }) => {
  const router = useRouter();

  return (
    <div className="relative">
      <Swiper
        slidesPerView={1}
        effect="creative"
        autoplay={{
          delay: 2500,
        }}
        modules={[Autoplay, EffectCreative]}
        centeredSlides={true}
        loop={true}
      >
        {project.pictures.map(({ id, url }, idx) => (
          <SwiperSlide key={id}>
            <div
              className={clsx(
                "flex flex-col items-center justify-center aspect-square md:aspect-video relative"
              )}
            >
              <IMG
                src={url}
                alt={project.title}
                fill
                sizes={sizes.xl}
                priority={idx === 0}
              />
              <div
                className={clsx(
                  "absolute h-1/2 w-full bottom-0",
                  "bg-gradient-to-t from-background dark:from-dark to-transparent"
                )}
              />
            </div>
          </SwiperSlide>
        ))}
      </Swiper>
      <div className="-mt-60 z-10 p-4 w-full h-full flex flex-col justify-end relative">
        <div className="p-4 max-w-4xl mx-auto">
          <h1 className="text-2xl font-bold mb-4">{project.title}</h1>
          <div className="rounded-lg my-4 italic">{project.abstract}</div>
          <div className="flex  flex-col md:flex-row gap-1">
            <Button
              color="primary"
              startContent={<FaArrowLeft />}
              onClick={() => {
                router.back();
              }}
              aria-label="Volver"
            >
              Volver
            </Button>
            {project.gitHub && (
              <Button
                color="primary"
                variant="ghost"
                as={Link}
                href={project.gitHub}
                isExternal
                aria-label="Ver en GitHub"
                startContent={
                  <Icon
                    name="git"
                    style={{
                      height: 20,
                      width: 20,
                    }}
                  />
                }
              >
                Ver en GitHub
              </Button>
            )}
            {project.webSide && (
              <Button
                color="primary"
                variant="ghost"
                as={Link}
                href={project.webSide}
                isExternal
                aria-label="Ver en la web"
                startContent={
                  <Icon
                    name="world"
                    style={{
                      height: 20,
                      width: 20,
                    }}
                  />
                }
              >
                Ver en la web
              </Button>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};
