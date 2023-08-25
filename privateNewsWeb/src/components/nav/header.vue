<template>
    <div class="header">
       <div class="content">
            <!--左侧 ： login及网站名称-->
            <div class="left">
                <img src="../../assets/images/logo_N.jpeg" alt="">
                <p>PrivateNews</p>
            </div>
            <!--右侧 ： 登录及个人中心-->
            <div class="right">
               <p class="help">帮助中心</p>     
               <p class="login">登录/注册</p>     
               <p class="userCenter">个人中心</p> 
            </div>
       </div>
    </div>
</template>

<script setup lang="ts">

</script>

<style scoped lang="scss">
.header{
    position: fixed;
    z-index: 999;
    width: 100%;
    height: 70px;
    background: #fff;
    display: flex;
    justify-content: center;
    .content{
        width: 1200px;
        height: 70px;
        background: white;
        display: flex;
        justify-content: space-between;
        .left{
            display: flex;
            justify-content: center;
            align-items: center;
            img{
                width: 50px;
                height: 50px;
                margin-right: 10px;
            }
            p{
                font-size: 20px;
                color: #55a6fe;
            }
        }
        .right{
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 14px;
            color: #ccc;
            p{
                margin-right: 10px;
            }
        }
    }
}

</style>