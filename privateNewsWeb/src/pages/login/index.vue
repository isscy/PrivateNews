<template>
    <div class="header">
        <!-- 顶部搜索栏 -->
        <van-row>
            <van-col span="24">
                <form action="/">
                    <van-search v-model="mainSearchValue" show-action placeholder="搜索关键字 / Tag / Topic"
                        @search="onMainSearching" @cancel="onMainSearchCancel" />
                </form>
            </van-col>
        </van-row>
        <div class="content">
            <div class="articleList">
                <!-- <van-list v-model:loading="articleListLoading" :finished="articleListFinished" finished-text="没有更多了"
                    @load="articleListOnLoad">
                    <van-cell v-for="item in articleList" :key="item.id" class="van-clearfix">
                        <template #default class="van-clearfix">
                            <div style="width:100%;min-height:100px;text-align: left;" class="van-clearfix">
                                <p><b>{{ item.title }}</b></p>
                                <p><b># {{ item.topic }}</b></p>
                            </div>
                        </template>
                    </van-cell>
                </van-list> -->
                <ul class="post-list">
                    <li class="post-item" v-for="post in articleList" :key="post.id">
                        <div class="post-item-container" @click="goDetail(post.content)">
                            <h4 class="post-item-container-title">{{ post.title }}</h4>
                            <h4 class="post-item-container-topic"># {{ post.topic }}</h4>
                            <h4 class="post-item-container-bulitInTag">
                                <van-tag round color="#ffe1e1" text-color="#ad0000">{{ post.internationalName }}</van-tag>
                                <van-tag round color="#ffe1e1" text-color="#ad0000">{{ post.regionName }}</van-tag>
                                <van-tag round color="#ffe1e1" text-color="#ad0000">{{ post.eventTypeName }}</van-tag>
                                <van-tag round v-for="tagItem in post.tagList" :key="tagItem">{{ tagItem }}</van-tag>
                            </h4>
                        </div>
                    </li>
                    <van-pagination v-model="currentPage" :page-count="pageCount" mode="simple"
                        @change="onPaginationChange" />
                </ul>

            </div>


        </div>
        <van-popup v-model:show="showPopup" position="bottom" :style="{ height: '80%' }" v-html="contentPopup" />

        <tabBar></tabBar>
    </div>
</template>

<script setup lang="ts">
import { ref, reactice, onMounted, onBeforeMount } from 'vue'
import router from "@/router/index";
import { showToast, showFailToast } from 'vant';
import { NewsDataEntity, ArticlePageSearchEntity, PagedEntity4NewsData, addArticle, getArticlePage } from "@/api/publish";
import { get, post } from '@/api/request.ts'
import tabBar from '@/components/nav/tabBar.vue'

const mainSearchValue = ref('');
const articleListLoading = ref(false)
const articleListFinished = ref(false)
const showPopup = ref(false)
const contentPopup = ref('')
const articleList = ref<NewsDataEntity[]>([])
const currentPage = ref<number>(1)
const pageSize = ref<number>(10)
const totalCount = ref<number>(0)
const pageCount = ref<number>(0)


const onMainSearching = async (val) => {
    clearSearch()
    getArticleList()
}
const onMainSearchCancel = () => {
    mainSearchValue.value = ""
    clearSearch()
}


const clearSearch = () => {
    articleList.value = []
    currentPage.value = 1
    totalCount.value = 0
    pageCount.value = 0
}

const articleListOnLoad = async () => {
    setTimeout(() => {
        getArticleList()
    }, 2500)

}
const onPaginationChange = async () => {
    getArticleList()
}
const goDetail = (content: string) => {
    showPopup.value = true
    contentPopup.value = content
}

/*const getArticleList = async () => {
    articleListLoading.value = true
    let param: ArticlePageSearchEntity = {
        currentPage: currentPage,
        pageSize: pageSize,
        searchValue: mainSearchValue.value
    }
    try {
        const { data, message } = await getArticlePage(param);
        console.log(data.rowDataList)
        if (!data) {
            showFailToast("分页获取News失败！");
            return
        } else if (!data.rowDataList || data.rowDataList.length == 0) {
            console.log(data.rowDataList)
            articleListFinished.value = true;
            return;
        }
        data.rowDataList.forEach(e => {
            if (e.tags) {
                e.tagList = e.tags.split(",")
            }
            articleList.value.push(e)
        });
        console.log(articleList.value)
        totalCount = data.totalCount
        pageCount = data.pageCount
        currentPage = data.currentPage + 1
        if (currentPage >= pageCount) {
            articleListFinished.value = true;
        }
    } catch (error) {
        showFailToast(error?.message);
    } finally {
        articleListLoading.value = false
    }
};*/

const getArticleList = async () => {
    articleListLoading.value = true
    let param: ArticlePageSearchEntity = {
        currentPage: currentPage,
        pageSize: pageSize,
        searchValue: mainSearchValue.value
    }
    try {
        const { data, message } = await getArticlePage(param);
        console.log(data.rowDataList)
        if (!data) {
            showFailToast("分页获取News失败！");
            return
        } else if (!data.rowDataList || data.rowDataList.length == 0) {
            console.log(data.rowDataList)
            articleListFinished.value = true;
            return;
        }
        data.rowDataList.forEach(e => {
            if (e.tags) {
                e.tagList = e.tags.split(",")
            }
            articleList.value.push(e)
        });
        console.log(articleList.value)
        totalCount.value = data.totalCount
        pageCount.value = data.pageCount
        currentPage.value = data.currentPage + 1
        if (currentPage.value >= pageCount.value) {
            articleListFinished.value = true;
        }
    } catch (error) {
        console.error(error)
        showFailToast("分页获取失败：", error?.message);
    } finally {
        articleListLoading.value = false
    }
};

onMounted(async () => {
    getArticleList()

    // let result = await get("/auth/test", "param").catch(err => {
    //     console.log(err);
    // })
    // console.log(result);
})

onBeforeMount(() => {
    if (!localStorage.getItem("accessToken")) router.push("/login");
});
</script>

<style scoped>
html,
body {
    height: 100%
}

.articleList {
    --van-list-text-align: left
}

.post-list {
    .post-item {
        list-style: none;
        border-radius: 4px;
        padding-left: 40px;
        cursor: pointer;
        border: 1px solid #ccc;
        margin-bottom: 10px;
        background-color: #fff;
        color: #878a8c;
        position: relative;
    }
}
</style>