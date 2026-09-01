import { defineConfig, type UserConfig, type ConfigEnv } from 'vite';
import vue from '@vitejs/plugin-vue';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig(({ mode }: ConfigEnv): UserConfig => {
  const isProd = mode === 'production';
  const isTest = mode === 'testbuild';

  return {
    plugins: [
      vue(),
      tailwindcss(),
    ],

    resolve: {
      alias: [
        { find: '@', replacement: '/src' }
      ],
    },

    server: {
      host: isProd || isTest ? '0.0.0.0' : true,
      port: isTest ? 8087 : isProd ? 5000 : 3000,
      hmr: isTest
        ? { host: '192.168.15.4', port: 8087 }
        : isProd
        ? { host: 'contasuper.com', port: 443 }
        : { clientPort: 8087 },
    },

    build: {
      target: 'esnext',
      chunkSizeWarningLimit: 2000,
    },
  };
});