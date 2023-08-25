// vue3 框架提供的createAPP方法，可以用来创建应用实例方法
import { createApp } from 'vue'
// 引入的样式文件
//import './style.css'
// 引入清除默认样式
import '@/style/reset.scss'
//Toast 组件是以函数形式提供的，如果项目中使用 unplugin-vue-components 插件来自动引入组件样式，则无法正确识别 Toast 组件，因此需要手动引入 Toast 组件的样式
import 'vant/es/toast/style';


// 引入根组件 App
// import App from './App.vue'
import App from '@/App.vue'
import router from '@/router/index.ts'
import {createPinia} from 'pinia'



// 引入全局组件：顶部、底部都是全局组件
import Header from '@/components/nav/header.vue'
import Footer from '@/components/nav/footer.vue'

// 利用createApp 方法创建应用实例， 且将应用实例挂载到挂载点上
// “挂载点” 位于： index.html 中的 一个id=app 的div
const app = createApp(App)
app.use(router)
app.use(createPinia())
app.component('Header', Header)
app.component('Footer', Footer)
app.mount('#app')
