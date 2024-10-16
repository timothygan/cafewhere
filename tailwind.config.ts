import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
        onyx: "#403F4C",
        raisin: "#2C2B3C",
        gunmetal: "#1B2432",
        richblack: "#121420",
        indianred: "#B76D68",
        "indianred-100": "#EBD7D6",
        "indianred-200": "#DDBCBB",
        "indianred-900": "#95504B",
        "indianred-700": "#AD625C",
        "indianred-400": "#C18985"
      },
      spacing: {
        half: "50%",
        fourth: "25%",
      },
      fontFamily: {
        'ivyora-display': ['ivyora-display', 'sans-serif'],
        'ivyora-text': ['ivyora-text', 'sans-serif'],
      },
    },
  },
  plugins: [],
};
export default config;
