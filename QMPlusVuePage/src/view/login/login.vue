<template>
        <el-container class="login-box">
            <el-header>登录</el-header>
                <el-main>输入框在这里
                    <el-button>
                        登录
                    </el-button>
                </el-main>
        </el-container>
</template>

<script>
import {mapState, mapGetters, mapActions , mapMutations} from 'vuex'
export default {
    name:"Login",
    computed:{
    ...mapGetters("User",["userName"]),
    ...mapState("User",["userName"])
    },
    methods: {
        ...mapActions("User",["AsyncSetUserName"]),
        ...mapMutations("User",["setUserName"])
    },
}
</script>

<style scoped lang="scss">
    .login-box{
        background: #409EFF
    }
</style>