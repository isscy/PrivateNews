<template>
    <div class="content">
        <van-cell-group title="Tag管理" inset>
            <van-cell title="内置Tag" to="AddOneTag" />
            <van-cell title="自定义Tag"   />
        </van-cell-group>
        <van-cell-group title="内容管理" inset>
            <van-cell title="发布一篇"  to="PublishOneArticle"/>
            <van-cell title="内容编辑"  />
        </van-cell-group>


    </div>
    <tabBar></tabBar>
</template>
  
<script setup lang="ts">
import router from "@/router/index";
import { ref, reactice, onMounted, onBeforeMount } from 'vue'
import tabBar from '@/components/nav/tabBar.vue'
const indexList = ref([1, 2])

</script>
<style lang="scss" scoped></style>
