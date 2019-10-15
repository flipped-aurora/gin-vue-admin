<template>
  <div>
    <h2 style="text-align:center">
    gin-vue-admin by 奇淼 And Spike
    </h2>
    <h4 style="text-align:center">
      微信：shouzi_1994
    </h4>
    
    <Animition />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Animition from '@/view/dashbord/component/animition.vue'
export default {
  name: 'Dashbord',
  computed: {
    ...mapGetters('user', ['userInfo'])
  },
  components: {
    Animition
  }
}
</script>

<style lang="scss" scoped>

</style>