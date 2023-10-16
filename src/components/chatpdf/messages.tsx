"use client";

import React, { useEffect } from "react";
import { Message, useChat } from "ai/react";
import { Input } from "@nextui-org/react";
import { MessageItem } from "./message-item";

interface Props {
  chatId: string;
  initialMessages: Message[];
}
export const Messages = ({ chatId, initialMessages }: Props) => {
  const { input, handleInputChange, handleSubmit, messages } = useChat({
    initialMessages,
    api: "/api/boot",
    body: {
      chatId,
    },
  });

  const containerRef = React.useRef<HTMLDivElement>(null);
  useEffect(() => {
    containerRef.current?.scrollTo({
      top: containerRef.current.scrollHeight,
      behavior: "smooth",
    });
  }, [messages]);

  return (
    <div
      ref={containerRef}
      className="h-[inherit] scroll overflow-y-auto relative px-4 flex flex-col gap-2"
    >
      <header className="sticky top-0 bg-background">
        <h2 className="text-lg font-bold text-center">Tus conversaciones</h2>
      </header>

      <main className="flex-1">
        {messages?.map((message) => {
          return <MessageItem key={message.id} message={message} />;
        })}
      </main>

      <footer className="sticky bottom-0 bg-background">
        <form onSubmit={handleSubmit}>
          <Input
            type="text"
            size="lg"
            value={input}
            onChange={handleInputChange}
            variant="underlined"
            placeholder="¿Qué quieres saber...?"
            aria-label="input de mensaje para el chatbot"
          />
        </form>
      </footer>
    </div>
  );
};
