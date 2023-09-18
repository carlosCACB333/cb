import { FC } from "react";
import {
  Card,
  CardHeader,
  Avatar,
  CardBody,
  CardFooter,
  Link,
} from "@nextui-org/react";
import { clsx } from "@nextui-org/shared-utils";
import NextImage from "next/image";
import { useAuthor } from "@/hooks";

interface UserTwitterCardProps {
  className?: string;
}

export const UserGitHubCard: FC<UserTwitterCardProps> = ({ className }) => {
  const { author } = useAuthor();

  return (
    <Card className={clsx("max-w-[300px]", className)}>
      <CardHeader className="justify-between">
        <div className="flex gap-5">
          <Avatar
            className="object-top"
            isBordered
            ImgComponent={NextImage}
            alt={author.firstName + " " + author.lastName}
            style={{
              objectPosition: "top",
            }}
            imgProps={{
              width: 40,
              height: 40,
            }}
            radius="full"
            size="md"
            src={author.photos![0].url}
          />
          <div className="flex flex-col items-start justify-center">
            <h4 className="text-sm font-semibold leading-none ">
              {author.firstName} {author.lastName}
            </h4>
            <Link
              className="text-sm"
              href={author.github?.toString()!}
              target="_blank"
              aria-label="Github"
            >
              @carloscb333
            </Link>
          </div>
        </div>
      </CardHeader>
      <CardBody className="px-3 py-0">
        <p className="text-sm pl-px ">
          Desarrollador Full-stack &nbsp;
          <span aria-label="confetti" role="img">
            🎉
          </span>
        </p>
      </CardBody>
      <CardFooter className="gap-3">
        <div className="flex gap-1">
          <p className="font-semibold text-xs">45</p>
          <p className="text-xs">Respositorios</p>
        </div>
      </CardFooter>
    </Card>
  );
};
