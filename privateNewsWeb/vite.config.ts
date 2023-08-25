import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
// 引用node提供的内置模块path: 可以获取绝对路径
import path from 'path'
import Components from 'unplugin-vue-components/vite'           // 添加
import {VantResolver} from 'unplugin-vue-components/resolvers'  // 添加



// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [VantResolver()],
    })
  ],
  server: {
    host: "192.168.10.88",
    port: 5173,
    proxy: {
      
    }
    

  },
  resolve: { // resolve：路径
    alias: { // alias： 别名
      "@": path.resolve(__dirname, 'src')  // 利用node提供的全局变量 __dirname 去获取src的绝对路径
    }
  }
})
