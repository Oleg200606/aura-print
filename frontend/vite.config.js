import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import tailwindcss from "@tailwindcss/vite";

export default ({ mode }) => {
  // eslint-disable-next-line no-undef
  const env = { ...process.env, ...loadEnv(mode, process.cwd()) };
  const isProduction = mode === "production";

  const BACKEND_HOST = isProduction ? "http://backend:8081" : "http://аура-принт.рф:8081";

  console.log("Current environment:", env.NODE_ENV);

  return defineConfig({
    plugins: [vue(), tailwindcss()],
    define: {
      "import.meta.env.BACKEND_HOST": JSON.stringify(BACKEND_HOST),
    },
    resolve: {
      alias: {
        "@": "/src",
      },
    },
    server: {
      port: 5173,
      allowedHosts: [env.SITE_HOST],
      proxy: {
        "/api": {
          target: BACKEND_HOST,
          changeOrigin: true,
          secure: false,
        },
      },
    },
  });
};
