import { defineConfig } from 'vite';
import FullReload from 'vite-plugin-full-reload';

export default defineConfig({
  plugins: [
    FullReload(['templates/**/*.html']),
  ],
  build: {
    outDir: 'static/dist',
    manifest: true,
    rollupOptions: {
      input: 'src/main.js',
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name].js`,
        assetFileNames: `assets/[name].[ext]`
      }
    },
  },
  server: {
    origin: 'http://localhost:5173',
    strictPort: true,
  },
});
