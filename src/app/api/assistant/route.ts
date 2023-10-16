import { OpenAI } from 'openai'
import { OpenAIStream, StreamingTextResponse } from 'ai'
import { env } from "@/utils";
import { savePdfMessages } from '@/services/message';
import { getContext, getContextWithoutAuth } from '@/services/boot';



const openai = new OpenAI({
  apiKey: env.openIa.apiKey,
});

// Set the runtime to edge for best performance
export const runtime = "edge";

export async function POST(req: Request) {

  const { messages } = await req.json();
  const chatId = process.env.ASSISTANT_CHAT_ID!
  const userMsg = messages.at(-1).content as string
  const { data } = await getContextWithoutAuth(chatId, userMsg)
  const response = await openai.chat.completions.create({
    model: 'gpt-3.5-turbo',
    stream: true,
    messages: [
      {
        role: "system",
        content: `
          Eres el asistente de carlos responde las preguntas de acuerdo al contexto que se te proporcione.
          START CONTEXT BLOCK
          ${data}
          END OF CONTEXT BLOCK
          Ten en cuenta el CONTEXT BLOCK que se te proporcione en una conversación. Si el contexto no proporciona la respuesta a la pregunta responde: "Lo siento, pero  no puedo enocntrar información para tu pregunta". No inventes nada que no se extraigas directamente del contexto
          `
      },
      ...messages

    ],
    max_tokens: 100,
    temperature: 0.2,
    top_p: 1,
    frequency_penalty: 1,
    presence_penalty: 1,
  })

  const stream = OpenAIStream(response, {});
  return new StreamingTextResponse(stream)

}
