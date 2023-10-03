export const env = {
  site: {
    url: process.env.SITE_URL || "",
  },

  cms: {
    url: process.env.NEXT_PUBLIC_GRAPHCMS_URL || "",
    token: process.env.NEXT_PUBLIC_GRAPHCMS_TOKEN || "",
    media: process.env.NEXT_PUBLIC_GRAPHCMS_MEDIA || "",
  },
  author: {
    email: process.env.AUTHOR_EMAIL || "",
    password: process.env.AUTHOR_PASSWORD || "",
  },
  revalidate: +(process.env.REVALIDATE || 60),
  mongo: {
    uri: process.env.MONGODB_URI || "",
  },
  openIa: {
    apiKey: process.env.OPENAI_API_KEY || "",
  },
  apiKey: process.env.API_KEY || "",
  back: {
    publicUrl: process.env.NEXT_PUBLIC_BACK_URL || "https://back.carloscb.com/api/v1/public",
    privateUrl: process.env.BACK_URL!,
    apiKey: process.env.BACK_API_KEY!,
    publicApiKey: process.env.NEXT_PUBLIC_BACK_API_KEY || "55wsdw2sedfw",
  },
  ckerk: {
    webHookSecret: process.env.WEBHOOK_SECRET!
  }
};

export const ALLOWED_LOCALES = ["es", "en"];

export const __PROD__ = process.env.NODE_ENV === "production";
