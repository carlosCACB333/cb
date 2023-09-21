"use client";

import { Button, Tooltip, Card, CardBody } from "@nextui-org/react";

import { useState } from "react";

import { title, subtitle, sectionWrapper } from "@/components";
import { Category } from "@/generated/graphql";
import { Icon } from "../common/icon";
import { GradientBox } from "../common/gradient-box";
import { SiIbmwatson } from "react-icons/si";

interface Props {
  categories: Category[];
}

export const SkillSection = ({ categories }: Props) => {
  const [selected, setSelected] = useState(categories[0]);

  return (
    <section className={sectionWrapper({ class: "z-20 mt-16 lg:mt-44" })}>
      <div className="flex flex-col gap-8">
        <div>
          <div className="text-center md:text-start">
            <h1 className={title({ size: "lg" })}>Mis &nbsp;</h1>
            <h1 className={title({ color: "green", size: "lg" })}>
              Habilidades
            </h1>
            <div className="flex flex-col md:flex-row items-center">
              <h1 className={title({ size: "lg" })}>tecnológicas&nbsp;</h1>
              <SiIbmwatson
                className="text-success animate-heartbeat"
                size={50}
                style={{
                  animationDuration: "2.5s",
                }}
              />
            </div>
          </div>
          <p className={subtitle()}>
            Mis habilidades tecnológicas para el desarrollo de software
            <span className="text-green-500 dark:text-green-400">
              {" "}
              Full-Stack{" "}
            </span>
            escalables, robustas y seguras.
          </p>
        </div>
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4 h-min">
            {categories.map((item) => (
              <Button
                size="lg"
                key={item.id}
                onClick={() => setSelected(item)}
                startContent={
                  <Icon
                    name={item.icon as any}
                    style={{
                      height: 30,
                      width: 30,
                    }}
                  />
                }
                className="justify-start"
                color={selected.id === item.id ? "success" : "default"}
                aria-label={item.name}
              >
                {item.name}
              </Button>
            ))}
          </div>

          <GradientBox
            className="items-start flex-wrap gap-4  min-h-[200px] lg:min-h-[390px] p-8 "
            color="green"
            to="right"
          >
            {selected.skills?.map((skill) => (
              <Tooltip
                key={skill.id}
                content={
                  <div className="px-1 py-2 max-w-xs">
                    <div className="text-small font-bold">{skill.name}</div>
                    <div className="text-tiny">{skill.detail}</div>
                  </div>
                }
              >
                <Card className="cursor-pointer">
                  <CardBody>
                    <Icon
                      name={skill.icon as any}
                      className="text-foreground fill-foreground"
                      style={{
                        height: 30,
                        width: 30,
                      }}
                    />
                  </CardBody>
                </Card>
              </Tooltip>
            ))}
          </GradientBox>
        </div>
      </div>
    </section>
  );
};
