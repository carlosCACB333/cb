import { ChatCard } from "@/components/chatpdf/chat-card";
import { DropFile } from "@/components/chatpdf/drop-file";
import { Footer } from "@/components/common/footer";
import { LayoutProps } from "@/interfaces";
import { getAllChatpdf } from "@/services";
import { formatDate } from "@/utils";
import clsx from "clsx";
import React from "react";

const IALayout = async ({ children, ...rest }: LayoutProps) => {
  const data = await getAllChatpdf();

  return (
    <main className="lg:h-[calc(100vh-4rem)] flex flex-col gap-2 lg:flex-row">
      <aside
        className={clsx(
          "rounded-lg max-w-sm w-full scroll overflow-y-auto relative",
          "px-4 flex flex-col gap-2"
        )}
      >
        <header className="sticky top-0 bg-background z-10">
          <br />
          <h2 className="text-2xl">Tus pdfs</h2>
          <DropFile />
          <br />
        </header>
        <div className="flex flex-1 flex-col gap-2 h-full">
          {data?.data?.map((chat) => {
            chat.createdAt = formatDate(chat.createdAt);
            return <ChatCard key={chat.id} chat={chat} />;
          })}
        </div>
        <div className="hidden lg:block sticky bottom-0 bg-background z-10 py-4">
          <Footer />
        </div>
      </aside>

      {children}
    </main>
  );
};

export default IALayout;
