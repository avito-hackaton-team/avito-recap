import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';

// В деве фронт и бэк живут на разных портах, поэтому /api проксируется на Go-сервис.
// Так в браузере не возникает CORS, а в проде тот же путь отдаёт nginx.
export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: process.env.VITE_API_PROXY_TARGET ?? 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
  // Тот же прокси для `vite preview`: собранная статика тоже должна видеть API.
  preview: {
    port: 4173,
    proxy: {
      '/api': {
        target: process.env.VITE_API_PROXY_TARGET ?? 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
  build: {
    outDir: 'dist',
    sourcemap: process.env.GENERATE_SOURCEMAP === 'true',
  },
});
