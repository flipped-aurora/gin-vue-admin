<template>
  <div>
    <h2 style="text-align:center">行星咖啡App by 林千凯</h2>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'Dashbord',
  data() {
    return {
      drawer: false
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo'])
  }
}
</script>

<style lang="scss" scoped>
</style>