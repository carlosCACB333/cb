import { DropFile } from "@/components/chatpdf/drop-file";
import { getAllChatpdf } from "@/services";
import { formatDate, subtitle } from "@/utils";

export default async function IAPage() {
  const data = await getAllChatpdf();
  console.log(data);
  return (
    <main className="flex h-full">
      <aside className="w-xs bg-content1 rounded-lg p-4">
        <h2 className={subtitle({})}>Tus chats</h2>
        <DropFile />
        <br />
        <div className="flex flex-col gap-2">
          {data.data.map((d) => (
            <div
              key={d.id}
              className="bg-background rounded-lg p-4 shadow-lg hover:opacity-80 cursor-pointer"
            >
              <h3 className="font-bold">{d.name}</h3>
              <p className="text-tiny">{formatDate(d.createdAt)}</p>
            </div>
          ))}
        </div>
      </aside>

      <section className="flex-1 grid grid-cols-2">
        <div>pdf</div>
        <div>chat</div>
      </section>
    </main>
  );
}
