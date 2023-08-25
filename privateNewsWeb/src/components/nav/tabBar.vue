<template>
    <van-tabbar v-model="active" route>
        <van-tabbar-item name="index" to="/" icon="home-o">首页</van-tabbar-item>
        <van-tabbar-item name="explore" to="/explore" icon="guide-o">探索</van-tabbar-item>
        <van-tabbar-item name="publish" to="/publish" icon="link-o">发布</van-tabbar-item>
        <van-tabbar-item name="mine" to="/mine" icon="manager-o">我的</van-tabbar-item>
    </van-tabbar>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';

const active = ref(0);



</script>

<style scoped></style>