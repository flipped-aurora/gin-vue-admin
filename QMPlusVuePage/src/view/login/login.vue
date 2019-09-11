<template>
        <el-container class="login-box">
            <el-header>登录</el-header>
                <el-main>输入框在这里
                    <el-button @click="login">
                        登录
                    </el-button>
                </el-main>
        </el-container>
</template>

<script>
import {mapState, mapGetters, mapActions , mapMutations} from 'vuex'
import {login} from '@/api/user'
export default {
    name:"Login",
    computed:{
    ...mapGetters("user",["userInfo"]),
    ...mapState("user",["token"])
    },
    methods: {
        ...mapActions("user",["AsyncSetUserInfo"]),
        ...mapMutations("user",["setUserInfo"]),
        login(){
            login({userName:"string1",passWord:"string"})
        }
    },
}
</script>

<style scoped lang="scss">
    .login-box{
        background: #409EFF
    }
</style>