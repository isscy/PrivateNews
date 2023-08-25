<template>
    <div class="container">
        <!-- 顶部全局组件 -->
        <!-- <van-row>
            <Header></Header>
         </van-row> -->
        <!-- 展示路由组件的区域 -->
        <!-- <van-row>
            <div class="content">xxxx</div>
         </van-row> -->
        <router-view></router-view>
        <!-- <tabBar></tabBar> -->



    </div>
</template>

<script setup lang="ts">
// import tabBar from '@/components/nav/tabBar.vue'

// 禁止手机端扩大和缩小手势
window.onload = function () {
    document.addEventListener('touchstart', function (event) {
        if (event.touches.length > 1) {
            event.preventDefault()
        }
    })
    document.addEventListener('gesturestart', function (event) {
        event.preventDefault()
    })
}
</script>

<style scoped></style>