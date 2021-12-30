<template>
  <span class="headerAvatar">
    <template v-if="picType === 'avatar'">
      <el-avatar v-if="userInfo.headerImg" :size="30" :src="avatar" />
      <el-avatar v-else :size="30" :src="noAvatar" />
    </template>
    <template v-if="picType === 'img'">
      <img v-if="userInfo.headerImg" :src="avatar" class="avatar">
      <img v-else :src="noAvatar" class="avatar">
    </template>
    <template v-if="picType === 'file'">
      <img :src="file" class="file">
    </template>
  </span>
</template>

<script>
export default {
  name: 'CustomPic'
}
</script>

<script setup>
import noAvatarPng from '@/assets/noBody.png'
import { useStore } from 'vuex'
import { computed, ref, defineProps } from 'vue-demi'
const store = useStore()

const props = defineProps({
  picType: {
    type: String,
    required: false,
    default: 'avatar'
  },
  picSrc: {
    type: String,
    required: false,
    default: ''
  }
})

const path = ref(import.meta.env.VITE_BASE_API + '/')
const noAvatar = ref(noAvatarPng)
const userInfo = computed(() => store.getters['user/userInfo'])

const avatar = computed(() => {
  if (props.picSrc === '') {
    if (userInfo.value.headerImg !== '' && userInfo.value.headerImg.slice(0, 4) === 'http') {
      return userInfo.value.headerImg
    }
    return path.value + userInfo.value.headerImg
  } else {
    if (props.picSrc !== '' && props.picSrc.slice(0, 4) === 'http') {
      return props.picSrc
    }
    return path.value + props.picSrc
  }
})
const file = computed(() => {
  if (props.picSrc && props.picSrc.slice(0, 4) !== 'http') {
    return path.value + props.picSrc
  }
  return props.picSrc
})

</script>

<style scoped>
.headerAvatar{
    display: flex;
    justify-content: center;
    align-items: center;
    margin-right: 8px;
}
.file{
    width: 80px;
    height: 80px;
    position: relative;
}
</style>
