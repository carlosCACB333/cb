import { formatDate } from "@/utils";
import { Message } from "ai/react";
import clsx from "clsx";
import React from "react";

interface Props {
  message: Message;
}

export const MessageItem = ({ message }: Props) => {
  // create a message item
  return (
    <div
      className={clsx("flex", {
        "justify-end": message.role === "user",
      })}
    >
      <div
        className={clsx("py-2 px-4 rounded-lg ", {
          "bg-primary text-primary-foreground": message.role === "user",
          "bg-content1": message.role === "system",
        })}
      >
        <p>{message.content}</p>
        <p className="text-tiny">{formatDate(message.createdAt!)}</p>
      </div>
    </div>
  );
};
