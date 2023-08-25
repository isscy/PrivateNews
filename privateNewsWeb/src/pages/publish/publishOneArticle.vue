<template>
    <div class="content">
        <van-nav-bar title="PublishNew" left-text="" left-arrow @click-left="onClickLeft" />
        <div class="editorContent">
            <van-form @submit="onSubmit">
                <van-cell-group title="主题" inset>
                    <van-field v-model="newData.title" name="title" label="标题" placeholder="标题"
                        :rules="[{ required: true, message: '请填写标题' }]" />
                    <van-field v-model="newData.topic" name="topic" label="主题" placeholder="主题"
                        :rules="[{ required: true, message: '请填写主题' }]" />
                    <van-field name="tagList" label="标签">
                        <template #input>
                            <div v-if="newData.tagList" v-for="tag in newData.tagList" :key="tag">
                                <van-tag round color="#ffe1e1" text-color="#ad0000" closeable @close="removeOneTag(tag)">{{
                                    tag }}</van-tag>
                                &nbsp; &nbsp;
                            </div>
                            <van-button plain icon="plus" type="primary" size="mini" @click="showAddTagBtn = true" />
                        </template>
                    </van-field>
                </van-cell-group>
                <van-cell-group title="内容" inset>
                    <wang-editor :initValue="newData.initContent" @getEditorContent="onEditorChange"></wang-editor>
                </van-cell-group>
                <van-cell-group title="内置标签" inset>
                    <van-field name="international" label="国域">
                        <template #input>
                            <van-radio-group v-model="newData.international" direction="horizontal">
                                <van-radio :name="1001">国内</van-radio>
                                <van-radio :name="1002">国外</van-radio>
                            </van-radio-group>
                        </template>
                    </van-field>
                    <van-field v-model="newData.regionName" is-link readonly name="region" label="所处地域"
                        placeholder="点击选择所处地域" @click="showRegionPicker = true" />

                    <van-field v-model="newData.eventTypeName" is-link readonly name="region" label="类型"
                        placeholder="点击选择类型" @click="showEnvntTypePicker = true" />


                </van-cell-group>

                <div style="margin: 16px;">
                    <van-button round block type="primary" native-type="submit">
                        提交
                    </van-button>
                </div>
            </van-form>


        </div>

        <van-popup v-model:show="showAddTagBtn" position="bottom" :style="{ height: '30%', padding: '9%' }" closeable>
            <!-- @click-overlay="onClickOverlay" @click-close-icon="onClickCloseIcon" -->
            <van-cell-group inset>
                <van-field v-model="addTagName" center clearable label="Tag名称" placeholder="Tag名称">
                    <template #button>
                        <van-button size="small" type="primary" @click="addOneTag">保存</van-button>
                    </template>
                </van-field>
            </van-cell-group>
        </van-popup>
        <van-popup v-model:show="showRegionPicker" position="bottom">
            <van-picker :columns="regionPickerList" @confirm="onRegionPickerConfirm" @cancel="showRegionPicker = false" />
        </van-popup>
        <van-popup v-model:show="showEnvntTypePicker" position="bottom">
            <van-picker :columns="envntTypePickerList" @confirm="onEnvntTypeConfirm"
                @cancel="showEnvntTypePicker = false" />
        </van-popup>

    </div>
</template>

<script setup lang="ts">
import router from "@/router/index";
import { showSuccessToast, showFailToast, Notify } from "vant";
import { onBeforeUnmount, ref, shallowRef, onMounted, nextTick } from 'vue'
import { BulitInTagParams, getBulitInTagList, BasePickerEntity } from "@/api/tag";
import { NewsDataEntity, addArticle } from "@/api/publish";
import WangEditor from '@/components/editor/WangEditor.vue'

const newData = ref<NewsDataEntity>({ id: 0, title: '', topic: '', tags: '',tagList:[], initContent: '', content: '', region: '', regionName: '', international: '', internationalName: '', eventType: '', eventTypeName: '' })
const addTagName = ref("")
const tagTypeList = ref([])
const showAddTagBtn = ref(false);
const showRegionPicker = ref(false);
const showEnvntTypePicker = ref(false);
//const regionPickerList  = ref<BasePickerEntity[]>([])
let regionPickerList: BasePickerEntity[] = []
let envntTypePickerList: BasePickerEntity[] = []

const onEditorChange = (arr, html) => {
    console.log(arr, html)
    newData.value.content = html;
}


const onSubmit = async () => {
    if(newData.value.tagList){
    }
    try {
        const { data, msg } = await addArticle(newData.value);
        console.log("data = ", data, "   msg = ", msg)
        showSuccessToast("发布成功！ id=" + data?.id)
        router.go(-1)
    } catch (error) {
        showFailToast(error?.message);
    }


}



const initTagData = async () => {
    let param: BulitInTagParams = {
        "tagKey": "aa",
        "tagName": "aa",
        "status": "1"
    }
    const { data } = await getBulitInTagList(param);
    tagTypeList.value = data
    regionPickerList = selectOneTagData("region")
    envntTypePickerList = selectOneTagData("eventType")
};

onMounted(async () => {
    initTagData();
});

const onRegionPickerConfirm = ({ selectedOptions }) => {
    newData.value.regionName = selectedOptions[0]?.text;
    newData.value.region = selectedOptions[0]?.value;
    showRegionPicker.value = false;
};
const onEnvntTypeConfirm = ({ selectedOptions }) => {
    newData.value.eventTypeName = selectedOptions[0]?.text;
    newData.value.eventType = selectedOptions[0]?.value;
    showEnvntTypePicker.value = false;
};
const addOneTag = () => {
    if (addTagName.value) {
        if (newData.value.tagList.includes(addTagName.value)) {
            showFailToast("重复了");
            addTagName.value = ""
            return
        }
        newData.value.tagList.push(addTagName.value)
        addTagName.value = ""
        showAddTagBtn.value = false
    }
}
const removeOneTag = (tag: string) => {
    let index = newData.value.tagList.indexOf(tag)
    if (index !== -1) {
        newData.value.tagList.splice(index, 1);
    }
}

// 返回上一页
const onClickLeft = () => history.back();
const selectOneTagData = (tagKey: string) => {
    let list: BasePickerEntity[] = []
    if (tagTypeList.value) {
        tagTypeList.value.forEach(element => {
            if (element.tagKey == tagKey) {
                if (element.items) {
                    element.items.forEach(e => {
                        let pickerEntity: BasePickerEntity = {
                            "text": e.tagName,
                            "value": e.tagKey
                        }
                        list.push(pickerEntity)
                    })

                }
            }
        });
    }
    return list;
}

</script>

<style scoped></style>