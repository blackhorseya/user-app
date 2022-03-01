import {defineConfig, loadEnv} from 'vite';
import react from '@vitejs/plugin-react';
import * as path from 'path';

// https://vitejs.dev/config/
export default ({mode}) => {
  process.env = {...process.env, ...loadEnv(mode, process.cwd())};
  const {VITE_PORT, VITE_API_ENDPOINT} = process.env;
  return defineConfig({
    server: {
      port: Number(VITE_PORT),
      proxy: {
        '/api': {
          target: VITE_API_ENDPOINT,
          changeOrigin: true,
        },
      },
    },
    resolve: {
      alias: {
        '@': path.resolve(__dirname, './src'),
      },
    },
    plugins: [react()],
    build: {
      target: 'esnext',
    },
  });
};
