<template>
    <div class="wrapper">
        <van-cell-group inset class="main">
            <van-tabs v-model:active="activeTab">
                <van-tab title="登录" name="login">
                    <Login ref="refLogin" />
                </van-tab>
                <van-tab title="注册" name="register">
                    <!-- <Register @toLogin="regSuccess" /> -->
                </van-tab>
            </van-tabs>
        </van-cell-group>
    </div>
</template>

<script setup lang="ts">
import { ref, onBeforeMount } from "vue";
import { useRouter } from "vue-router";
import Login, { API as LoginAPI } from "@/components/login/LoginForm.vue";
import {useCounterStore} from '@/store/counter'
//import Register from "@/components/login/RegisterForm.vue";
const counter = useCounterStore;
const router = useRouter();
onBeforeMount(() => {
    if (localStorage.getItem("accessToken")) router.push("/");
});
const refLogin = ref<LoginAPI | undefined>();
const activeTab = ref<string>("login");
const regSuccess = (val: string) => {
    activeTab.value = "login";
    if (refLogin.value) {
        refLogin.value.username = val;
    }
};
</script>


<style lang="scss" scoped>
.wrapper {
    height: 100vh;
    background-image: url("@/assets/images/login_background.jpeg");
    padding-top: 25vh;
}

.main {
    min-height: 30vh;
    margin: 10px;
}
</style>