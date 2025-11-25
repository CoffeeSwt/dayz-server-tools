import { defineConfig, loadEnv } from 'vite'
import * as path from 'path'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'

export default (mode: string) => {
  const env = loadEnv(mode, process.cwd())
  return defineConfig({
    plugins: [vue(),tailwindcss()],
    resolve: {
      //设置别名
      alias: {
        '@': path.resolve(__dirname, './src'),
      },
    },
    server: {
      proxy: {
        [env.VITE_BASE_API]: {
          target: `${env.VITE_BASE_URL}:${env.VITE_SERVER_PORT}`,
          changeOrigin: true,
          rewrite: path => path.replace(new RegExp('^' + env.VITE_BASE_API), ''),
        },
      }
    }
  })
};


