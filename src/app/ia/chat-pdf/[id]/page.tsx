import { PDFViewer } from "@/components/chatpdf/PDFViewer";
import { Messages } from "@/components/chatpdf/messages";
import { PageProps } from "@/interfaces";
import { getLastChatpdfMessages } from "@/services/message";
import { Message } from "ai/react";

export default async function ChatpdfDetail({ params }: PageProps) {
  const { data = [] } = await getLastChatpdfMessages(params.id);
  const initialMsg: Message[] = data.map((msg) => ({
    id: msg.id,
    content: msg.content,
    role: msg.role,
    createdAt: new Date(msg.createdAt),
  }));

  return (
    <section className="flex-1 h-[inherit] flex">
      <div className="h-full flex-1">
        <PDFViewer id={params.id + ".pdf"} />
      </div>
      <div className="max-h-full h-[inherit] max-w-md w-full">
        <Messages chatId={params.id} initialMessages={initialMsg} />
      </div>
    </section>
  );
}
