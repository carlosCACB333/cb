import { z } from "zod";

export const ContacSchema = z.object({
  name: z.string().min(3).max(50),
  email: z.string().email(),
  phone: z.string().min(9).max(11),
  affair: z.string().min(3).max(50),
  message: z.string().min(3).max(500),
});
