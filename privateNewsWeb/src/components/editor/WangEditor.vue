/**
  initValue: 父组件传递过来的富文本框初始值，这里会实时监听，更新初始值，放置在弹窗中使用，没有钩子函数的更新。
  getEditorContent() 方法, 父组件通过这个方法获取富文本编辑器的内容, 包括数组格式的和html格式的内容
*/
<template> 
    <!-- <div v-html="valueHtml"></div> -->
    <div style="border: 1px solid #ccc">
        <Toolbar style="border-bottom: 1px solid #ccc" :editor="editorRef" :defaultConfig="toolbarConfig" :mode="mode" />
        <Editor style="height: 300px; overflow-y: hidden; text-align: left;" v-model="valueHtml"
            :defaultConfig="editorConfig" :mode="mode" @onCreated="handleCreated" @onChange="handleChange" />
    </div>
</template>

<script setup lang="ts">
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { onBeforeUnmount, ref, shallowRef, onMounted, nextTick, watch } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import {  } from "@/api/publish";
// 初始值
const props = defineProps({
    initValue: String // 父组件传递过来的富文本框初始值
})
// 父组件通过这个方法获取富文本编辑器的内容
const emits = defineEmits(['getEditorContent'])
const mode = ref('default')
// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef()
// 内容 HTML
const valueHtml = ref('')
// 模拟 ajax 异步获取内容
onMounted(() => {
    nextTick(() => { // 界面节点更新完以后再修改值。
        valueHtml.value = props.initValue
    })
})
// 工具栏配置
const toolbarConfig = {
    toolbarKeys: [
        // 菜单 key
        'headerSelect',
        'bold', // 加粗
        'italic', // 斜体
        'through', // 删除线
        'underline', // 下划线
        'bulletedList', // 无序列表
        'numberedList', // 有序列表
        'color', // 文字颜色
        'insertLink', // 插入链接
        'fontSize', // 字体大小
        'lineHeight', // 行高
        'uploadImage', // 上传图片
        'delIndent', // 缩进
        'indent', // 增进
        'deleteImage',//删除图片
        'divider', // 分割线
        'insertTable', // 插入表格
        'justifyCenter', // 居中对齐
        'justifyJustify', // 两端对齐
        'justifyLeft', // 左对齐
        'justifyRight', // 右对齐
        'undo', // 撤销
        'redo', // 重做
        'clearStyle', // 清除格式
        'fullScreen' // 全屏
    ]
}
const editorConfig = {
    placeholder: '请输入内容...', // 配置默认提示
    // readOnly: true,
    MENU_CONF: {                // 配置上传服务器地址
        uploadImage: {
            // 小于该值就插入 base64 格式（而不上传），默认为 0
            base64LimitSize: 5 * 1024, // 5kb
            // 单个文件的最大体积限制，默认为 2M
            // maxFileSize: 1 * 1024 * 1024, // 1M
            // // 最多可上传几个文件，默认为 100
            // maxNumberOfFiles: 5,
            // 选择文件时的类型限制，默认为 ['image/*'] 。如不想限制，则设置为 []
            allowedFileTypes: ['image/*'],
            // 自定义上传
            async customUpload(file, insertFn) { // 文件上传
                const formData = new FormData();
                formData.set('file', file)

                // 这里根据自己项目封装情况，设置上传地址
                axios.defaults.headers.common['Authorization'] = 'Bearer eyJhbGciOiJIUzUxMiJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX2tleSI6IjVhYzc3MDQxLTE3NGItNDEwZC1hOTJlLTVkYzU0YmRhMjllNiIsInVzZXJuYW1lIjoiYWRtaW4ifQ.GhdthjLmw4NP2WV2owYQS70tacRM-pX7NqI7Mo0_j_hatiqtSrqAr0RJC40osRETiQo_dacZcvNqATLsJHL77A'
                let result = await axios.post('https://zxc.test.cn/file/upload', formData)
                // 插入到富文本编辑器中，主意这里的三个参数都是必填的，要不然控制台报错：typeError: Cannot read properties of undefined (reading 'replace')
                insertFn(result.data.data.url, result.data.data.name, result.data.data.name)
            }
        }
    }
}
// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
    const editor = editorRef.value
    if (editor == null) return
    editor.destroy()
})
const handleCreated = (editor) => {
    editorRef.value = editor // 创建富文本编辑器
}
const handleChange = (info) => {
    // info.children 数组包含了数据类型和内容，valueHtml.value内容的html格式
    emits('getEditorContent', info.children, valueHtml.value)
}
watch(() => props.initValue, (value) => {  // 父组件获取初始值
    valueHtml.value = value
})
</script>

