
<template>
    <van-form @submit="onSubmit" :readonly="buttonLoading" validate-trigger="onSubmit" label-width="65px">
        <van-cell-group inset>
            <van-field v-model="username" name="username" label="用户名" placeholder="测试账号: test" clearable :rules="[
                { required: true, message: '请填写用户名' },
                {
                    validator: (val: string) => check.checkToUserName(val),
                    message: '用户名格式错误',
                },
            ]" />
            <van-field v-model="password" type="password" name="password" placeholder="测试密码: test1234" label="密码" clearable
                :rules="[
                    { required: true, message: '请填写密码' },
                    {
                        validator: (val: string) => check.checkToPassword(val),
                        message: '密码格式错误',
                    },
                ]" />
            <van-field v-model="keyword" name="keyword" placeholder="口令: 芝麻开门" label="口令" clearable
                :rules="[
                    { required: true, message: '请填写口令' }
                ]" />
        </van-cell-group>
        <div style="margin: 16px">
            <van-button round block type="primary" native-type="submit" :loading="buttonLoading">
                登录
            </van-button>
        </div>
    </van-form>
</template>

<script setup lang="ts">
import { ref } from "vue";
import router from "@/router/index";
import { showToast, showFailToast, showSuccessToast  } from "vant";
import { get, post } from '@/api/request.ts'
import { login, LoginAndRegParams } from "@/api/user";
import { check } from "@/utils/index";
export interface API {
    username: string;
}
// const router = useRouter();
const username = ref("");
const password = ref("");
const keyword = ref("");
const buttonLoading = ref(false);

const onSubmit = async (values: LoginAndRegParams) => {
    if (!values?.username?.trim() || !values?.password?.trim()) {
        showFailToast("请输入用户名或密码");
        return;
    }
    try {
        buttonLoading.value = true;
        let { data, message } = await login(values);
        console.log("data=", data, "; message=", message)
        if (data.accessToken) {
            localStorage.setItem("accessToken", data.accessToken);
            localStorage.setItem("refreshToken", data.refreshToken);
            showSuccessToast({ message: "登录成功", duration: 600 });
            setTimeout(() => {
                router.push("/");
            }, 500);
        }
    } catch (error) {
        showFailToast(error?.message);
    } finally {
        buttonLoading.value = false;
    }
};
defineExpose({
    username,
});
</script>


