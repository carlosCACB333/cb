import NotFound from "@/app/not-found";
import { PDFViewer } from "@/components/chatpdf/PDFViewer";
import { Messages } from "@/components/chatpdf/messages";
import { PageProps } from "@/interfaces";
import { getLastChatpdfMessages } from "@/services/message";
import { Message } from "ai/react";

export default async function ChatpdfDetail({ params }: PageProps) {
  const { data, status } = await getLastChatpdfMessages(params.id);
  if (status !== "success") {
    return <NotFound />;
  }
  const initialMsg: Message[] =
    data?.map((msg) => ({
      id: msg.id,
      content: msg.content,
      role: msg.role,
      createdAt: new Date(msg.createdAt),
    })) || [];

  return (
    <section className="flex-1 h-full lg:flex">
      <div className="h-full flex-1">
        <PDFViewer id={params.id} />
      </div>
      <div className="h-full w-full max-w-md">
        <Messages chatId={params.id} initialMessages={initialMsg} />
      </div>
    </section>
  );
}
