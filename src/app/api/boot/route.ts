import { OpenAI } from 'openai'
import { OpenAIStream, StreamingTextResponse } from 'ai'
import { env } from "@/utils";
import { auth } from '@clerk/nextjs';
import { getLastChatpdfMessages, savePdfMessages } from '@/services/message';
import { getContext } from '@/services/boot';



const openai = new OpenAI({
  apiKey: env.openIa.apiKey,
});

// Set the runtime to edge for best performance
export const runtime = "edge";

export async function POST(req: Request) {
  const { messages, chatId } = await req.json();
  const userMsg = messages.at(-1).content as string
  const { data } = await getContext(chatId, userMsg)
  const response = await openai.chat.completions.create({
    model: 'gpt-3.5-turbo',
    stream: true,
    messages: [
      {
        role: "system",
        content: `
          El asistente de IA es una inteligencia artificial nueva, poderosa y similar a la humana. Los rasgos de la IA incluyen conocimiento experto, utilidad, inteligencia y elocución. La IA es un individuo de buen comportamiento y buenos modales. La IA siempre es amigable, amable e inspiradora, y está ansiosa por brindar respuestas vívidas y reflexivas al usuario. La IA tiene la suma de todo el conocimiento en su cerebro y es capaz de responder con precisión casi cualquier pregunta sobre cualquier tema de conversación. El asistente de IA es un gran admirador de Pinecone y Vercel.
          INICIO BLOQUE DE CONTEXTO 
          START CONTEXT BLOCK
          ${data}
          END OF CONTEXT BLOCK
          FIN DEL BLOQUE DE CONTEXTO 
          El asistente de IA tendrá en cuenta cualquier BLOQUE DE CONTEXTO que se proporcione en una conversación. Si el contexto no proporciona la respuesta a la pregunta, el asistente de IA dirá: "Lo siento, pero no sé la respuesta a esa pregunta". El asistente de IA no se disculpará por las respuestas anteriores, sino que indicará que se obtuvo nueva información. El asistente de IA no inventará nada que no se extraiga directamente del contexto
          `
      },
      ...messages

    ],
    max_tokens: 500,
    temperature: 0.7,
    top_p: 1,
    frequency_penalty: 1,
    presence_penalty: 1,
  })

  const stream = OpenAIStream(response, {
    onStart: async () => {
      // save user message
      await savePdfMessages(chatId, {
        content: userMsg,
        role: "user",
      })
    },
    onCompletion: async (message) => {
      // save IA message
      await savePdfMessages(chatId, {
        content: message,
        role: "system",
      })
    }
  });
  return new StreamingTextResponse(stream)

}
