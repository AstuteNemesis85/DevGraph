import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

export default defineConfig({
  plugins: [react()],
  base: '/',   // ✅ FIXED
  server: {
    port: 3000,
    proxy: {
      '/auth': 'http://localhost:8080',
      '/api': 'http://localhost:8080'
    }
  },
  build: {
    outDir: 'dist'
  }
})