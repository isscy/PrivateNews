import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

import Index from '@/pages/login/index.vue'
import Login from '@/pages/login/login.vue'
import Mine from '@/pages/login/mine.vue'

// 通过Vite的import.meta.glob()方法实现自动化导入路由
// const mainRouterModules = import.meta.glob('../layout/*.vue')
// const viewRouterModules = import.meta.glob('../views/**/*.vue')


const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        redirect: "/index"
    },
    {
        path: '/index',
        name: 'Index',
        component: Index,
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
    },
    {
        path: '/explore',
        name: 'Explore',
        component: () => import('@/pages/browse/findNews.vue'),
    },
    {
        path: '/publish',
        name: 'Publish',
        component: () => import('@/pages/publish/publish.vue'),
    },
    {
        path: '/addOneTag',
        name: 'AddOneTag',
        component: () => import('@/pages/publish/addOneTag.vue'),
    },
    {
        path: '/publishOneArticle',
        name: 'PublishOneArticle',
        component: () => import('@/pages/publish/publishOneArticle.vue'),
    },
    {
        path: '/mine',
        name: 'Mine',
        component: Mine,
    },

];

const router = createRouter({
    //history: createWebHistory(process.env.BASE_URL),
    history: createWebHistory(),
    routes,
});

export default router;




