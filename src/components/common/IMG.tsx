"use client";
import { env } from "@/utils";
import clsx from "clsx";
import Image, { ImageLoaderProps, ImageProps } from "next/image";
import React, { FC } from "react";

type Fit = "clip" | "crop" | "scale" | "max";
export interface ImgProps extends ImageProps {
  fit?: Fit;
  alt: string;
}

export const IMG: FC<ImgProps> = ({
  fit = "max",
  alt,
  className,
  ...props
}) => {
  const loader = ({ src, width }: ImageLoaderProps) => {
    const id = src.split("/").at(-1)!;
    const url = `${env.cms.media}/resize=fit:${fit},width:${width}/${id}`;
    return url;
  };

  return (
    <div className="relative h-full w-full">
      <Image
        className={clsx("object-cover object-top", className)}
        loader={loader}
        alt={alt}
        fill
        {...props}
      />
    </div>
  );
};
