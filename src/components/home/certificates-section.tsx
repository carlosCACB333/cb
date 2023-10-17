"use client";
import "swiper/css";
import "swiper/css/pagination";
import { Button } from "@nextui-org/react";
import { ArrowRightIcon } from "@nextui-org/shared-icons";
import NextLink from "next/link";
import { Certification } from "@/generated/graphql";
import { Autoplay, EffectCreative, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/react";
import { AiFillSafetyCertificate } from "react-icons/ai";
import { IMG } from "../common/IMG";
import { sizes } from "@/assets";
import { sectionWrapper, subtitle, title } from "..";
interface Props {
  certifications: Certification[];
}
export const CertificateSection = ({ certifications }: Props) => {
  return (
    <section
      className={sectionWrapper({
        isBlurred: true,
        class:
          "border-t border-b border-divider px-8 flex justify-center items-center mt-16 lg:mt-44",
      })}
      id="home-certifications"
    >
      <div className="w-full mx-8 py-10 grid grid-cols-12 gap-6 md:gap-0 z-20 items-center">
        <div className="flex flex-col gap-2 col-span-12 md:col-span-5">
          <div className="text-center lg:text-start">
            <h1 className={title({ size: "lg" })}>Mis &nbsp;</h1>
            <div className="flex flex-col items-center lg:flex-row">
              <h1 className={title({ color: "violet", size: "lg" })}>
                certificaciones&nbsp;
              </h1>
              <AiFillSafetyCertificate
                className="text-secondary animate-heartbeat"
                size={50}
                style={{
                  animationDuration: "2.5s",
                }}
              />
            </div>
          </div>
          <p className={subtitle({ class: "md:w-full text-base lg:text-lg" })}>
            Aquí encontrarás las certificaciones que he obtenido a lo largo de
            mi carrera profesional.
          </p>
          <div className="flex flex-row gap-3 justify-start">
            <Button
              as={NextLink}
              className="text-sm"
              color="secondary"
              endContent={
                <ArrowRightIcon
                  className="group-data-[hover=true]:translate-x-0.5 outline-none transition-transform"
                  strokeWidth={2}
                  aria-label="Ver más certificaciones"
                />
              }
              aria-label="Ver más certificaciones"
              href="/certificate"
              size="lg"
            >
              Ver más
            </Button>
          </div>
        </div>
        <div className="col-span-12 md:col-span-7">
          <Swiper
            slidesPerView={"auto"}
            effect="creative"
            spaceBetween={30}
            autoplay={{
              delay: 2500,
              disableOnInteraction: false,
            }}
            pagination={{
              clickable: true,
            }}
            modules={[Pagination, Autoplay, EffectCreative]}
          >
            {certifications.map(({ id, picture, name }, idx) => (
              <SwiperSlide key={id}>
                <div className="aspect-[4/3] ">
                  <IMG src={picture.url} alt={name} sizes={sizes.md} />
                </div>
              </SwiperSlide>
            ))}
          </Swiper>
        </div>
      </div>
    </section>
  );
};
