
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

  const toggleTheme = (dark , event) => {
    if (dark) {
      theme.value = 'dark';
    } else {
      theme.value = 'light';
    }
    const isAppearanceTransition =
      document.startViewTransition &&
      !window.matchMedia('(prefers-reduced-motion: reduce)').matches
    if (!isAppearanceTransition || !event) {
      return
    }
    const x = event.clientX
    const y = event.clientY
    const endRadius = Math.hypot(
      Math.max(x, innerWidth - x),
      Math.max(y, innerHeight - y)
    )
    // @ts-expect-error: Transition API
    const transition = document.startViewTransition(async () => {
      await nextTick()
    })
    transition.ready.then(() => {
      const clipPath = [
        `circle(0px at ${x}px ${y}px)`,
        `circle(${endRadius}px at ${x}px ${y}px)`,
      ]
      document.documentElement.animate(
        {
          clipPath:  theme.value === 'dark' ? [...clipPath].reverse() : clipPath,
        },
        {
          duration: 400,
          easing: 'ease-in-out',
          pseudoElement: '::view-transition-new(root)'
        }
      )
    })



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
