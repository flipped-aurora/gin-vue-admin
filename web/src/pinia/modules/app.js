
import { defineStore } from 'pinia'
import { ref, onMounted, nextTick, watchEffect } from 'vue'
export const useAppStore = defineStore('app', () => {
  const theme = ref(localStorage.getItem('theme') || 'light')
  const device = ref("")

  watchEffect(() =>{
    if ( theme.value === 'dark'){
      document.documentElement.classList.add('dark');
      document.documentElement.classList.remove('light');
      localStorage.setItem('theme', 'dark');
    }else{
        document.documentElement.classList.add('light');
        document.documentElement.classList.remove('dark');
        localStorage.setItem('theme', 'light');
    }
  })

  const toggleTheme = (dark , _) => {
    if (dark) {
      theme.value = 'dark';
    } else {
      theme.value = 'light';
    }

  }


  const toggleDevice = (e) => {
      device.value = e;
  }


  return {
    theme,
    device,
    toggleTheme,
    toggleDevice
  }

})
