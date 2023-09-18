import { Configuration, OpenAIApi } from "openai-edge";
import { OpenAIStream, StreamingTextResponse } from "ai";

import { env } from "@/utils";

// Create an OpenAI API client (that's edge friendly!)
const openAIConfig = new Configuration({
  apiKey: env.openIa.apiKey,
});
const openai = new OpenAIApi(openAIConfig);

// Set the runtime to edge for best performance
export const runtime = "edge";

export async function POST(req: Request) {
  const { prompt } = await req.json();

  // Ask OpenAI for a streaming completion given the prompt
  const response = await openai.createCompletion({
    model: "gpt-3.5-turbo",
    stream: true,
    temperature: 0.6,
    prompt: ["This is a test"],
  });
  // Convert the response into a friendly text-stream
  const stream = OpenAIStream(response);
  // Respond with the stream
  return new StreamingTextResponse(stream);
}
