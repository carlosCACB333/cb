import { Card, CardBody, CardFooter, colors } from "@nextui-org/react";
import { UserGitHubCard } from "./user-github-card";
import { FaAws, FaDocker, FaGithub, FaPython, FaReact } from "react-icons/fa";
import { IconCard } from "./Icon-card";
import { SiIbmwatson } from "react-icons/si";
import { BiLogoKubernetes } from "react-icons/bi";
import { GopherIcon, Icon } from "@/components/common/icon";
import python from "@/assets/img/python.png";
import Image from "next/image";

export const FloatingComponents: React.FC<{}> = () => {
  return (
    <div className="hidden md:flex flex-col relative z-20 w-1/2">
      <>
        <FaReact
          size={50}
          className="text-cyan-400 absolute -top-[220px] -right-[40px] animate-[levitate_13s_ease_infinite_1s_reverse]"
        />

        <IconCard
          className="absolute -top-[130px] -right-[120px] animate-[levitate_10s_ease_infinite]"
          color={colors.blue[500]}
          aria-label="ts"
        >
          <span className="font-extrabold">TS</span>
        </IconCard>

        <Card
          isFooterBlurred
          className="absolute -top-[260px] right-[100px] animate-[levitate_12s_ease_infinite_1s] z-0 max-w-fit"
        >
          <CardBody>
            <Image src={python} alt="Python" />
          </CardBody>
        </Card>

        <IconCard
          className="absolute left-[170px] -top-[160px] animate-[levitate_17s_ease_infinite_1s]"
          color={colors.green[500]}
          aria-label="python"
        >
          <FaGithub size={30} />
        </IconCard>

        <UserGitHubCard className="absolute left-[80px] -top-[50px] animate-[levitate_16s_ease_infinite] border-none" />

        <div className="absolute right-[110px] -top-[60px] animate-[levitate_18s_ease_infinite] z-10 max-w-fit border-none">
          <GopherIcon height={100} width={100} />
        </div>

        <div className="absolute z-10 -top-[40px] -right-[230px] animate-[levitate_14s_ease_infinite_1s]">
          <FaDocker size={50} className="text-primary" />
        </div>

        <IconCard
          className="absolute left-[200px] top-[160px] max-w-fit animate-[levitate_14s_ease_infinite_0.5s]"
          color={colors.yellow[500]}
          aria-label="k8s"
        >
          <BiLogoKubernetes size={50} />
        </IconCard>

        <div className="absolute right-[10px] top-[30px] animate-[levitate_16s_ease_infinite] z-10 max-w-fit border-none">
          <SiIbmwatson size={50} />
        </div>

        <Card
          isFooterBlurred
          className="absolute right-[60px] top-[100px] animate-[levitate_12s_ease_infinite_1s] z-0 max-w-fit"
        >
          <CardBody>
            <FaAws aria-label="aws" size={100} />
          </CardBody>
        </Card>
      </>
    </div>
  );
};
