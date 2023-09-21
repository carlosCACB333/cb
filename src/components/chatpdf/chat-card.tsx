"use client";
import { Chatpdf } from "@/interfaces";
import clsx from "clsx";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React, { FC } from "react";

interface Props {
  chat: Chatpdf;
}
export const ChatCard: FC<Props> = ({ chat }) => {
  const pathname = usePathname();
  const id = pathname.split("/").pop();

  return (
    <Link
      href={`/ia/chat-pdf/${chat.id}`}
      key={chat.id}
      className={clsx(
        "rounded-lg p-4 shadow-lg hover:opacity-80 cursor-pointer",
        id === chat.id && "bg-primary-900 dark:bg-primary-200",
        id !== chat.id && "bg-content1"
      )}
    >
      <h3 className="font-bold">{chat.name}</h3>
      <p className="text-tiny">{chat.createdAt}</p>
    </Link>
  );
};
