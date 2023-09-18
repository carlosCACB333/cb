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
  isProd: process.env.NODE_ENV === "production",
  mongo: {
    uri: process.env.MONGODB_URI || "",
  },
  openIa: {
    apiKey: process.env.OPENAI_API_KEY || "",
  },
  apiKey: process.env.API_KEY || "",
  back: {
    publicUrl: process.env.NEXT_PUBLIC_BACK_URL || "",
    privateUrl: `http://${process.env.BACK_HOST}:${process.env.BACK_PORT}/api/v1`
  },
  ckerk: {
    webHookSecret: process.env.WEBHOOK_SECRET!
  }
};

export const ALLOWED_LOCALES = ["es", "en"];

export const __PROD__ = process.env.NODE_ENV === "production";
export const __DEV__ = process.env.NODE_ENV !== "production";
export const __TEST__ = process.env.NODE_ENV === "test";
export const __PREVIEW__ = process.env.IS_PREVIEW === "true";
