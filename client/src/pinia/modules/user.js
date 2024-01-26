import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

export const useUserStore = defineStore('user', () => {

  const userInfo = ref({
    name: 'admin',
    avatar: '头像url',
    roles: ['admin'],
    introduction: 'I am a super administrator',
  })
  const token = ref(window.localStorage.getItem('token') || '')
  const setUserInfo = (val) => {
    userInfo.value = val
  }

  const setToken = (val) => {
    token.value = val
  }

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value
    }
  }
  /* 获取用户信息*/
  const GetUserInfo = async() => {

  }

  watch(() => token.value, () => {
    window.localStorage.setItem('token', token.value)
  })

  return {
    userInfo,
    token,
    ResetUserInfo,
    GetUserInfo,
    setToken,
  }
})
