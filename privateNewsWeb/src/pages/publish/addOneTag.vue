<template>
    <div class="content">
        <van-nav-bar title="Tag管理" left-text="" left-arrow @click-left="onClickLeft" />
        <div class="tagList">
            <div v-for="item in tagTypeList" :key="item.id" inset>
                <van-cell-group  inset>
                    <template #title>
                        {{item.tagName}} &nbsp;
                        <van-icon name="edit" class="edit-icon" @click="editOneTag(item.id)" />
                    </template>
                    <div v-if="item.items" v-for="subItem in item.items" :key="subItem.id" >
                        <van-cell :title="subItem.tagName" @click="editOneSubTag(subItem.id)" />
                            
                    </div>
                </van-cell-group>
            </div>



        </div>



    </div>
</template>

<script setup lang="ts">
import { showSuccessToast, Notify } from "vant";
import { ref, reactive, watchEffect, onMounted } from "vue";
import { BulitInTagParams, getBulitInTagList } from "@/api/tag";

const tagTypeList = ref([])


onMounted(async () => {
    initTagData();
});
const initTagData = async () => {
    let param: BulitInTagParams = {
        "tagKey": "aa",
        "tagName": "aa",
        "status": "1"
    }
    const { data } = await getBulitInTagList(param);
    tagTypeList.value = data
};
const editOneSubTag = (id) => {
    console.log(id)
}
const editOneTag = (id) => {
    console.log(id)
}
// 返回上一页
const onClickLeft = () => history.back();
</script>

<style scoped>
.edit-icon {
    font-size: 16px;
    line-height: inherit;
}
</style>