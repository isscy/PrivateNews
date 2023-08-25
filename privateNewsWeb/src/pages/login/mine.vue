<template>
    <div class="content">
        <van-nav-bar title="个人中心" left-text="" left-arrow @click-left="onClickLeft" />
        <div class="userinfo">
            <div class="userinfo-avatar">
                <van-image round fit="contain" width="80" height="80" :src="userinfo.avatar" alt="头像"
                    @click="clickAvatar" />
            </div>
            <div class="userinfo-name">
                <span> {{ userinfo.username }} </span>
            </div>
        </div>

        <van-cell-group title="My News" inset>
            <van-cell title="我发布的" icon="coupon-o" is-link @click="clickEditPassword" />
            <van-cell title="我关注的" icon="star-o" is-link @click="clickAboutMe" />
        </van-cell-group>

        <van-cell-group title="消息通知" inset>
            <van-cell title="联系客服" icon="service-o" is-link @click="clickEditPassword" />
            <van-cell title="系统通知" icon="fire-o" is-link @click="clickAboutMe" />
        </van-cell-group>


        <van-cell-group title="我的信息" inset>
            <van-cell title="修改密码" icon="brush-o" is-link @click="clickEditPassword" />
            <van-cell title="关于我" icon="setting-o" is-link @click="clickAboutMe" />
        </van-cell-group>

        <van-action-bar>
            <van-action-bar-icon icon="replay" text="刷新" @click="onReplay" />
            <van-action-bar-button type="danger" text="退出登录" @click="loginOut" />
        </van-action-bar>
    </div>


    <van-dialog v-model:show="canEditUserinfo" title="修改用户信息" show-cancel-button @confirm="canEditUserinfo = false"
        :beforeClose="editSignature">
        <van-field v-model="userinfo.newSignature" placeholder="请输入个性签名" clearable />
    </van-dialog>
</template>

<script setup lang="ts">
import { showSuccessToast, Notify } from "vant";
import { ref, reactive, watchEffect, onMounted } from "vue";
import router from "@/router/index";
import { getUserInfo, editUserInfo } from "@/api/user";
import avatar_default from '@/assets/images/avatar_default.png'


const canEditUserinfo = ref(false);
const userinfo = reactive({
    username: "匿名用户",
    signature: "",
    newSignature: "",
    avatar: avatar_default
});


// 获取用户信息
onMounted(async () => {
    initUserData();
});
const initUserData = async () => {
    const { data } = await getUserInfo();
    userinfo.username = data.nickname || data.username;
    //userinfo.signature = data.signature;
    //userinfo.avatar = data.avatar;
};

// 修改密码
const clickEditPassword = () => {
    router.push({
        path: "/update-password",
    });
};


// 退出登录
const loginOut = () => {
    localStorage.removeItem("accessToken");
    showSuccessToast({ message: "退出成功", duration: 1000 });
    setTimeout(() => { router.push("/login"); }, 1000);
};

// 返回上一页
const onClickLeft = () => history.back();
</script>



<style lang="scss" scoped>
.content {
    .userinfo {
        margin-top: 35px;
        margin-bottom: 35px;
        display: grid;
        align-items: center;
        justify-items: center;

        .userinfo-avatar {

            // background: url("@/assets/images/avatar_default.png") repeat;
            div {
                border: 2px solid #fff;
            }
        }
    }
}
</style>