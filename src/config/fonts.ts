import { Roboto } from "next/font/google";

export const fontRoboto = Roboto({
  subsets: ["latin"],
  variable: "--font-roboto",
  style: ["normal"],
  weight: ["300", "400", "500", "700", "900"], // light, normal, medium, bold, black
});
