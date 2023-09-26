"use client";
import { Chatpdf } from "@/interfaces";
import { deleteChatpdf } from "@/services";
import { Button, Card, CardBody, CircularProgress } from "@nextui-org/react";
import clsx from "clsx";
import { usePathname, useRouter } from "next/navigation";
import React, { FC, useState } from "react";
import { MdClose } from "react-icons/md";
import { toast } from "react-toastify";

interface Props {
  chat: Chatpdf;
}
export const ChatCard: FC<Props> = ({ chat }) => {
  const pathname = usePathname();
  const id = pathname.split("/").pop();
  const { push, refresh, replace, forward } = useRouter();
  const [isDeleting, setIsDeleting] = useState(false);

  const handleDelete = async () => {
    setIsDeleting(true);
    const res = await deleteChatpdf(chat.id);
    setIsDeleting(false);
    toast(res.message, {
      type: res.status,
    });
    refresh();
    if (res.status === "success" && id === chat.id) {
      replace(`/ia/chat-pdf`, {});
    }
  };
  return (
    <Card
      key={chat.id}
      className={clsx("hover:opacity-80 cursor-pointer", {
        "bg-primary-900 dark:bg-primary-200": id === chat.id,
      })}
    >
      <CardBody
        className="relative"
        onClick={() => {
          push(`/ia/chat-pdf/${chat.id}`);
        }}
      >
        <h3 className="font-bold">{chat.name}</h3>
        <p className="text-tiny">{chat.createdAt}</p>

        <Button
          className="absolute top-0 right-0"
          isIconOnly
          variant="light"
          onClick={handleDelete}
          disabled={isDeleting}
        >
          {isDeleting ? <CircularProgress size="sm" /> : <MdClose />}
        </Button>
      </CardBody>
    </Card>
  );
};
