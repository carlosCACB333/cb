"use server";

import { Locale, Stage } from "@/generated/graphql";
import { env } from "@/utils";
import { getSdk } from "@/utils/sdk";

export const getAuthor = async (locale: Locale) => {
  try {
    const { author } = await getSdk().getAuthor({
      email: env.author.email,
      locales: [locale],
      stage: Stage.Published,
    });

    return author;
  } catch (error) {
    console.error(error);
    return undefined;
  }
};
