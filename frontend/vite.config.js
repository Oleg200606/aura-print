import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

const isProduction = process.env.NODE_ENV === "production";

const BACKEND_HOST = isProduction ? "http://backend:8081" : "http://аура-принт.рф:8081";

console.log("Current environment:", process.env.NODE_ENV);

export default defineConfig({
  define: {
    "import.meta.env.BACKEND_HOST": JSON.stringify(BACKEND_HOST),
  },
  plugins: [vue()],
  resolve: {
    alias: {
      "@": "/src",
    },
  },
  server: {
    port: 5173,
    allowedHosts: [process.env.SITE_HOST],
    proxy: {
      "/api": {
        target: BACKEND_HOST,
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
