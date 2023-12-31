"use client";
import { Button, Input } from "@nextui-org/react";
import { Message } from "ai";
import { useChat } from "ai/react";
import clsx from "clsx";
import React, { useEffect, useState } from "react";
import { BiSend, BiUser } from "react-icons/bi";
import { MdMessage } from "react-icons/md";

import { motion } from "framer-motion";
import { FaRobot } from "react-icons/fa";

export const ChatBoot = () => {
  const [isOpenTooltip, setIsopenTooltip] = useState(false);
  const [isOpen, setIsOpen] = useState(false);
  const { input, handleInputChange, handleSubmit, messages, isLoading } =
    useChat({
      api: "/api/assistant",
    });

  const containerRef = React.useRef<HTMLDivElement>(null);
  useEffect(() => {
    containerRef.current?.scrollTo({
      top: containerRef.current.scrollHeight,
      behavior: "smooth",
    });
  }, [messages]);

  useEffect(() => {
    setTimeout(() => {
      setIsopenTooltip(true);
    }, 4000);
  }, []);

  return (
    <aside className="fixed bottom-[1rem] right-[1rem] z-40 ">
      <div className="relative">
        <Button
          className="z-50 shadow-2xl"
          isIconOnly
          color="primary"
          radius="full"
          size="lg"
          onClick={() => {
            setIsopenTooltip(false);
            setIsOpen(!isOpen);
          }}
          aria-label="Abrir chatbot"
        >
          <MdMessage size={24} />
        </Button>
        {isOpenTooltip && (
          <div
            className={clsx(
              "absolute top-0.5 right-full mr-3 bg-primary rounded-lg p-2 whitespace-nowrap animate-levitate",
              "before:content-[''] before:absolute before:bottom-[0.8rem] before:left-[calc(100%-0.5rem)] before:bg-primary before:w-3 before:h-3 before:rotate-45"
            )}
          >
            ¿Qué quieres saber de mí...?
          </div>
        )}

        <motion.div
          animate={isOpen ? "open" : "closed"}
          initial="closed"
          variants={{
            open: {
              opacity: 1,
              display: "block",
              y: 0,
              transition: {
                duration: 0.3,
              },
            },
            closed: {
              opacity: 0,
              display: "none",
              y: 100,
              transition: {
                duration: 0.3,
              },
            },
          }}
          className="absolute bg-content1 bottom-full right-0 mb-2 rounded-md shadow-2xl"
        >
          <div className="h-8 bg-content1 rounded-md"></div>
          <div
            className="p-4 scroll overflow-y-auto w-[90vw] sm:w-96 h-[60vh] sm:h-[34rem]"
            ref={containerRef}
          >
            <MessageItem
              message={{
                content:
                  "Hola, soy el asistente virtual de Carlos, ¿en qué puedo ayudarte?",
                id: "1u#s",
                role: "assistant",
              }}
            />
            {messages.map((m, i) => (
              <MessageItem
                key={m.id}
                message={m}
                isTyping={isLoading && i === messages.length - 1}
              />
            ))}
          </div>
          <footer>
            <form className="m-4" onSubmit={handleSubmit}>
              <Input
                aria-label="Input de asistente virtual"
                value={input}
                onChange={handleInputChange}
                size="lg"
                placeholder="¿Qué quieres saber...?"
                width="100%"
                variant="underlined"
                autoComplete="off"
                endContent={
                  <Button
                    isIconOnly
                    radius="full"
                    className="text-xl text-foreground"
                    size="sm"
                    type="submit"
                    aria-label="Enviar mensaje"
                  >
                    <BiSend />
                  </Button>
                }
              />
            </form>
          </footer>
        </motion.div>
      </div>
    </aside>
  );
};

const MessageItem = ({
  message,
  isTyping = false,
}: {
  message: Message;
  isTyping?: boolean;
}) => {
  return (
    <div
      key={message.id}
      className={clsx(
        "flex gap-1 items-end justify-end my-2",
        { "flex-row ml-6": message.role === "user" },
        { "flex-row-reverse mr-6": message.role !== "user" }
      )}
    >
      <span
        className={clsx("px-4 py-3 rounded-xl inline-block ", {
          "rounded-br-none bg-primary": message.role === "user",
          "rounded-bl-none bg-primary-900 dark:bg-primary-100":
            message.role !== "user",
        })}
        style={{
          overflowWrap: "anywhere",
        }}
      >
        {message.content}
      </span>

      <div className="text-2xl">
        {message.role === "user" ? (
          <BiUser className="text-primary-500" />
        ) : (
          <FaRobot
            className={clsx("text-primary-500", {
              "animate-bounce": isTyping,
            })}
          />
        )}
      </div>
    </div>
  );
};
