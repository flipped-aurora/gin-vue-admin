<template>
  <div>
    <wfd-gva ref="wfd" :data="demoData" :height="600" :users="candidateUsers" :groups="candidateGroups" :categorys="categorys" :lang="lang" />
  </div>
</template>
<script>


import wfdGva from 'wfd-gva'
export default {
  name: 'Workflow',
  components:{
    wfdGva
  },
  data () {
    return {
      lang: "zh",
      demoData: {},
      candidateUsers: [],
      candidateGroups: []
    }
  },
  methods:{
    save(){
      console.log(this.$refs['wfd'].graph.save())
    }
  }
}
</script>