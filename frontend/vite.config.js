import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

const isProduction = process.env.NODE_ENV === "production";

console.log("Current environment:", process.env.NODE_ENV);

export default defineConfig({
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
        target: isProduction ? "http://backend:8081" : "http://localhost:8081",
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
