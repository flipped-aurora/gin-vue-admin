<template>
    <span class="headerAvatar">
        <template v-if="picType === 'avatar'">
            <el-avatar :size="30" :src="avatar" v-if="userInfo.headerImg"></el-avatar>
            <el-avatar :size="30" :src="require('@/assets/noBody.png')" v-else></el-avatar>
        </template>
        <template v-if="picType === 'img'">
            <img :src="avatar" class="avatar" v-if="userInfo.headerImg" />
            <img :src="require('@/assets/noBody.png')" class="avatar" v-else/>
        </template>
        <template v-if="picType === 'file'">
            <img :src="file" class="file"/>
        </template>
    </span>
</template>

<script>
import { mapGetters } from 'vuex'
const path = process.env.VUE_APP_BASE_API
export default {
    name: "customPic",
    props: {
        picType: {
            type: String,
            required: false,
            default: "avatar"
        },
        picSrc: {
            type: String,
            required: false,
            default: ""
        }
    },
    data(){
        return{
            path: path + '/',
        }
    },
    computed:{
        ...mapGetters('user', ['userInfo']),
        avatar(){
            if(this.picSrc === ''){
 return this.userInfo.headerImg
            }else{
                return this.picSrc
            }
        },
        file(){
            return this.picSrc
        }
    }
}
</script>

<style scoped>
.headerAvatar{
    display: flex;
    justify-content: center;
    align-items: center;
}
.file{
    width: 80px;
    height: 80px;
    position: relative;
}
</style>
