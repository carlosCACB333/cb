"use client";

import {
  Table,
  TableBody,
  TableCell,
  TableRow,
  TableHeader,
  TableColumn,
} from "@nextui-org/react";

import { title, subtitle, sectionWrapper } from "@/components";
import { useAuthor } from "@/hooks";
import { MDXContent } from "@/components/md/MDXContent";
import { AiFillHeart } from "react-icons/ai";
import { IMG } from "@/components/common/IMG";
import { sizes } from "@/assets";

export const AboutSection = () => {
  const { author } = useAuthor();
  const fullName = `${author.firstName} ${author.lastName}`;
  const age = new Date().getFullYear() - new Date(author.birth).getFullYear();
  const photo = author.photos.at(-1)!;

  return (
    <section
      className={sectionWrapper({ class: "mt-24 lg:mt-56" })}
      id="home-about-me"
    >
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div className="row-span-2  md:row-start-1 bg-gradient-to-br from-blue-200 to-blue-900 rounded-lg min-h-[400px]">
          <IMG src={photo.url} alt={author.firstName} sizes={sizes.sm} />
        </div>
        <div className="row-start-1 md:row-start-2 ">
          <div className="text-center md:text-start">
            <h1 className={title({ size: "lg" })}>Conoce &nbsp;</h1>
            <div className="flex flex-col items-center md:flex-row">
              <h1 className={title({ color: "blue", size: "lg" })}>
                más de mí&nbsp;
              </h1>
              <AiFillHeart
                className="text-primary animate-heartbeat"
                size={50}
                style={{
                  animationDuration: "2.5s",
                }}
              />
            </div>
          </div>
          <div className={subtitle({ fullWidth: true })}>
            <MDXContent>{author.bio?.toString()}</MDXContent>
          </div>
          <Table
            hideHeader
            removeWrapper
            isCompact
            aria-label="Tabla de información personal"
          >
            <TableHeader>
              <TableColumn> </TableColumn>
              <TableColumn> </TableColumn>
            </TableHeader>

            <TableBody>
              <TableRow key={1}>
                <TableCell>Nombre</TableCell>
                <TableCell>{fullName}</TableCell>
              </TableRow>
              <TableRow key={2}>
                <TableCell>Correo</TableCell>
                <TableCell>{author.email}</TableCell>
              </TableRow>
              <TableRow key={3}>
                <TableCell>Edad</TableCell>
                <TableCell>{age}</TableCell>
              </TableRow>
              <TableRow key={5}>
                <TableCell>Profesión</TableCell>
                <TableCell>{author.profession?.toString()}</TableCell>
              </TableRow>
              <TableRow key={6}>
                <TableCell>Universidad</TableCell>
                <TableCell>{author.university}</TableCell>
              </TableRow>
              <TableRow key={7}>
                <TableCell>Ciclo</TableCell>
                <TableCell>{author.cycle?.toString()}</TableCell>
              </TableRow>
              <TableRow key={8}>
                <TableCell>Dirección</TableCell>
                <TableCell>{author.address}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </div>
      </div>
    </section>
  );
};
