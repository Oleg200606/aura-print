import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://149.154.68.194:8081',
        changeOrigin: true,
        secure: false
      }
    }
  }
})