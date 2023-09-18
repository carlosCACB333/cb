import { nextui } from "@nextui-org/theme";

const text = {
  title: "#d4ddfb",
  body: "#bdc8f0",
};

/** @type {import("@nextui-org/theme").ColorScale} */
const primary = {
  DEFAULT: "#7390f4",
  foreground: text.title,
  50: "#0f121a",
  100: "#232a3c",
  200: "#374156",
  300: "#4b578f",
  400: "#588aef",
  500: "#6b9aec",
  600: "#7ca6e9",
  700: "#8da2e6",
  800: "#9eaee3",
  900: "#d4ddfb",
};

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        "code-foreground": "rgb(var(--code-foreground) / <alpha-value>)",
        "code-syntax1": "rgb(var(--code-syntax1) / <alpha-value>)",
        "code-syntax2": "rgb(var(--code-syntax2) / <alpha-value>)",
        "code-syntax3": "rgb(var(--code-syntax3) / <alpha-value>)",
        "code-syntax4": "rgb(var(--code-syntax4) / <alpha-value>)",
        "code-syntax5": "rgb(var(--code-syntax5) / <alpha-value>)",
        "code-syntax6": "rgb(var(--code-syntax6) / <alpha-value>)",
        "code-removed": "rgb(var(--code-removed) / <alpha-value>)",
        "code-string": "rgb(var(--code-string) / <alpha-value>)",
        "code-class": "rgb(var(--code-class) / <alpha-value>)",
        "code-punctuation": "rgb(var(--code-punctuation) / <alpha-value>)",
        "code-number": "rgb(var(--code-number) / <alpha-value>)",
        "code-added": "rgb(var(--code-added) / <alpha-value>)",
        "code-line-number": "rgb(var(--code-line-number) / <alpha-value>)",
        "code-faded-line": "rgb(var(--code-faded-line) / <alpha-value>)",
        "code-comment": "rgb(var(--code-comment) / <alpha-value>)",
        "code-keyword": "rgb(var(--code-keyword) / <alpha-value>)",
        "code-function": "rgb(var(--code-function) / <alpha-value>)",
        "code-tag": "rgb(var(--code-tag) / <alpha-value>)",
        "code-attr-name": "rgb(var(--code-attr-name) / <alpha-value>)",
        "code-language-javascript":
          "rgb(var(--code-language-javascript) / <alpha-value>)",
        "code-highlighted-word1-bg":
          "rgb(var(--code-highlighted-word1-bg) / <alpha-value>)",
        "code-highlighted-word1-bg-active":
          "rgb(var(--code-highlighted-word1-bg-active) / <alpha-value>)",
        "code-highlighted-word1-text":
          "rgb(var(--code-highlighted-word1-text) / <alpha-value>)",
        "code-highlighted-word2-bg":
          "rgb(var(--code-highlighted-word2-bg) / <alpha-value>)",
        "code-highlighted-word2-bg-active":
          "rgb(var(--code-highlighted-word2-bg-active) / <alpha-value>)",
        "code-highlighted-word2-text":
          "rgb(var(--code-highlighted-word2-text) / <alpha-value>)",
        "code-highlighted-word3-bg":
          "rgb(var(--code-highlighted-word3-bg) / <alpha-value>)",
        "code-highlighted-word3-bg-active":
          "rgb(var(--code-highlighted-word3-bg-active) / <alpha-value>)",
        "code-highlighted-word3-text":
          "rgb(var(--code-highlighted-word3-text) / <alpha-value>)",
        "shadow-highlighted": "rgb(var(--shadow-highlighted) / <alpha-value>)",
      },
      keyframes: {
        heartbeat: {
          "0%": { transform: "scale(1)" },
          "50%": { transform: "scale(1.2)" },
          "100%": { transform: "scale(1)" },
        },
        levitate: {
          "0%": {
            transform: "translateY(0)",
          },
          "30%": {
            transform: "translateY(-10px)",
          },
          "50%": {
            transform: "translateY(4px)",
          },
          "70%": {
            transform: "translateY(-15px)",
          },
          "100%": {
            transform: "translateY(0)",
          },
        },
        expand: {
          "0%": { transform: "scale(1)" },
          "50%": { transform: "scale(1.2)" },
          "100%": { transform: "scale(1)" },
        },
        "expand-opacity": {
          "0%": {
            opacity: 0,
            transform: "scale(1)",
          },
          "50%": {
            opacity: 1,
            transform: "scale(1.3)",
          },
          "100%": {
            opacity: 0,
            transform: "scale(1.295)",
          },
        },
      },
      animation: {
        heartbeat: "heartbeat 1s ease-in-out infinite",
        levitate: "levitate 5s ease infinite",
        expand: "expand 6s ease-out infinite both",
        "expand-opacity": "expand-opacity 6s linear infinite both",
      },
    },
  },
  darkMode: "class",
  plugins: [
    nextui({
      themes: {
        dark: {
          colors: {
            background: {
              DEFAULT: "#07090e",
            },
            foreground: {
              DEFAULT: text.title,
              500: text.body,
            },
            content1: {
              DEFAULT: "#0f121a",
            },
            content2: {
              DEFAULT: "#0f121a",
            },
            default: {
              DEFAULT: "#0F121A",
            },
            primary,
          },
        },
        light: {
          colors: {
            background: {
              DEFAULT: "#eef6ff",
            },
            foreground: {
              DEFAULT: "#000001",
              500: "#404152",
            },
            content1: {
              DEFAULT: "#eaf5ff",
            },
            content2: {
              DEFAULT: "#0f121a",
            },
            default: {
              DEFAULT: "#eaf5ff",
            },
            primary,
          },
        },
      },
    }),
  ],
};
