
import { defineStore } from 'pinia'
import { ref, watchEffect, reactive } from 'vue'
import originSetting from  "@/config.json"
import {  setBodyPrimaryColor } from '@/utils/format'
export const useAppStore = defineStore('app', () => {
  const theme = ref(localStorage.getItem('theme') || 'light')
  const device = ref("")
  const config = reactive({
    weakness: false,
    grey: false,
    primaryColor: '#79B6E7',
    showTabs: true,
    darkMode: 'light',
    layout_side_width : 256,
    layout_side_collapsed_width : 80,
    layout_side_item_height : 48,
    show_watermark: false,

    side_mode : 'normal'
  })

  // 初始化配置
  Object.keys(originSetting).forEach(key => {
      config[key] = originSetting[key]
    if(key === 'primaryColor'){
      setBodyPrimaryColor(originSetting[key],config.darkMode)
    }
  })

  if (localStorage.getItem('darkMode')) {
    config.darkMode = localStorage.getItem('darkMode')
  }


  watchEffect(() =>{
    if (theme.value === 'dark'){
      document.documentElement.classList.add('dark');
      document.documentElement.classList.remove('light');
      localStorage.setItem('theme', 'dark');
    }else{
        document.documentElement.classList.add('light');
        document.documentElement.classList.remove('dark');
        localStorage.setItem('theme', 'light');
    }
  })

  const toggleTheme = (dark) => {
    if (dark) {
      theme.value = 'dark';
    } else {
      theme.value = 'light';
    }
  }

  const toggleWeakness = (e) => {
    config.weakness = e;
    if(e) {
      document.documentElement.classList.add('html-weakenss');
    }else{
      document.documentElement.classList.remove('html-weakenss');
    }
  }

  const toggleGrey = (e) => {
    config.grey = e;
    if(e) {
      document.documentElement.classList.add('html-grey');
    }else{
        document.documentElement.classList.remove('html-grey');
    }
  }

  const togglePrimaryColor = (e) => {
    config.primaryColor = e;
    setBodyPrimaryColor(e,config.darkMode)
  }

  const toggleTabs = (e) => {
    config.showTabs = e;
  }

  const toggleDevice = (e) => {
    device.value = e;
  }

  const toggleDarkMode = (e) => {
    config.darkMode = e
    localStorage.setItem('darkMode', e)
    if(e === 'dark'){
      toggleTheme(true)
    }else{
      toggleTheme(false)
    }
  }

  const toggleDarkModeAuto = () =>{
    // 处理浏览器主题
    const darkQuery = window.matchMedia('(prefers-color-scheme: dark)')
    const dark = darkQuery.matches
    toggleTheme(dark)
    darkQuery.addEventListener('change', (e) => {
      toggleTheme(e.matches)
    })
  }

  const toggleConfigSideWidth = (e) => {
    config.layout_side_width = e;
  }

  const toggleConfigSideCollapsedWidth = (e) => {
      config.layout_side_collapsed_width = e;
  }

  const toggleConfigSideItemHeight = (e) => {
      config.layout_side_item_height = e;
  }

  const toggleConfigWatermark = (e) => {
        config.show_watermark = e;
  }

  const toggleSideModel= (e) =>{
    config.side_mode = e
  }

  if(config.darkMode === 'auto'){
    toggleDarkModeAuto()
  }

    toggleGrey(config.grey)

  return {
    theme,
    device,
    config,
    toggleTheme,
    toggleDevice,
    toggleWeakness,
    toggleGrey,
    togglePrimaryColor,
    toggleTabs,
    toggleDarkMode,
    toggleConfigSideWidth,
    toggleConfigSideCollapsedWidth,
    toggleConfigSideItemHeight,
    toggleConfigWatermark,
    toggleSideModel
  }

})
