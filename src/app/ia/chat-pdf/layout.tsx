import { subtitle } from "@/components";
import { ChatCard } from "@/components/chatpdf/chat-card";
import { DropFile } from "@/components/chatpdf/drop-file";
import { LayoutProps } from "@/interfaces";
import { getAllChatpdf } from "@/services";
import { formatDate } from "@/utils";
import clsx from "clsx";
import React from "react";

const IALayout = async ({ children, ...rest }: LayoutProps) => {
  const data = await getAllChatpdf();

  return (
    <main className="flex h-[inherit] flex-col lg:flex-row">
      <aside
        className={clsx(
          "rounded-lg max-w-sm w-full h-[inherit] scroll overflow-y-auto relative",
          "px-4"
        )}
      >
        <div className="sticky top-0 bg-background">
          <br />
          <h2 className="text-2xl">Tus pdfs</h2>
          <DropFile />
          <br />
        </div>
        <div className="flex flex-col gap-2">
          {data?.data?.map((chat) => {
            chat.createdAt = formatDate(chat.createdAt);
            return <ChatCard key={chat.id} chat={chat} />;
          })}
        </div>
        <br />
      </aside>

      {children}
    </main>
  );
};

export default IALayout;
