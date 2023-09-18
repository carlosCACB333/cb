"use server";

import { Contact } from "@/interfaces";
import { env, getMail } from "@/utils";
import { createTransport } from "nodemailer";

export const sendMail = async (data: Contact) => {
  try {
    const { name, email, message, affair, phone } = data;

    const transporter = createTransport({
      host: "smtp.gmail.com",
      port: 465,
      secure: true,
      auth: {
        user: env.author.email,
        pass: env.author.password,
      },
    });

    await transporter.sendMail({
      from: env.author.email,
      to: env.author.email,
      subject: affair,
      html: getMail(name, email, message, phone),
    });

    return {
      ok: true,
      message:
        "Gracias por contactarnos, pronto nos pondremos en contacto contigo",
    };
  } catch (error) {
    console.log(error);
    return {
      ok: false,
      message:
        "Tuvimos un problema al enviar tu mensaje. Por favor, intenta más tarde.",
    };
  }
};
