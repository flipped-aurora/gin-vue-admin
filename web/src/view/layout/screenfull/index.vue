<template>
  <div @click="click">
    <!-- <svg
            t="1508738709248"
            class="screenfull-svg"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="2069"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            width="32"
            height="32"
            v-if="isShow"
           >
                <path
                    d="M333.493443 428.647617 428.322206 333.832158 262.572184 168.045297 366.707916 64.444754 64.09683 64.444754 63.853283 366.570793 167.283957 262.460644Z"
                    p-id="2070"/>
                <path
                    d="M854.845439 760.133334 688.61037 593.95864 593.805144 688.764889 759.554142 854.56096 655.44604 958.161503 958.055079 958.161503 958.274066 656.035464Z"
                    p-id="2071"/>
                <path
                    d="M688.535669 428.550403 854.31025 262.801405 957.935352 366.921787 957.935352 64.34754 655.809313 64.081481 759.919463 167.535691 593.70793 333.731874Z"
                    p-id="2072"/>
                <path
                    d="M333.590658 594.033341 167.8171 759.804852 64.218604 655.67219 64.218604 958.270996 366.342596 958.502263 262.234493 855.071589 428.421466 688.86108Z"
                    p-id="2073"/>
        </svg> -->
    <svg
      v-if="isShow"
      t="1590133227479"
      class="screenfull-svg"
      viewBox="0 0 1024 1024"
      version="1.1"
      width="20"
      height="20"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path id="svg_1" d="m928.512,959.744a32,32 0 0 1 -32,-32l0,-256a32,32 0 0 1 64,0l0,256a32,32 0 0 1 -32,32z" />
      <path
        id="svg_2"
        d="m960.512,927.744a32,32 0 0 1 -32,32l-256,0a32,32 0 0 1 0,-64l256,0a32,32 0 0 1 32,32zm-864.768,-863.488a32,32 0 0 1 32,32l0,256a32,32 0 0 1 -64,0l0,-256a32,32 0 0 1 32,-32z"
      />
      <path id="svg_3" d="m63.744,96.256a32,32 0 0 1 32,-32l256,0a32,32 0 0 1 0,64l-256,0a32,32 0 0 1 -32,-32z" />
      <path id="svg_4" d="m958.030718,91.777575a32,32 0 0 1 -32,32l-256,0a32,32 0 0 1 0,-64l256,0a32,32 0 0 1 32,32z" />
      <path id="svg_5" d="m926.030718,59.777575a32,32 0 0 1 32,32l0,256a32,32 0 0 1 -64,0l0,-256a32,32 0 0 1 32,-32z" />
      <path
        id="svg_6"
        d="m940.622718,69.250038a32,32 0 0 1 0,45.248l-247.936,247.936a32,32 0 0 1 -45.248,-45.248l247.936,-247.936a31.936,31.936 0 0 1 45.248,0z"
      />
      <path id="svg_7" d="m61.649508,930.478492a32,32 0 0 1 32,-32l256,0a32,32 0 0 1 0,64l-256,0a32,32 0 0 1 -32,-32z" />
      <path id="svg_8" d="m93.649508,962.478492a32,32 0 0 1 -32,-32l0,-256a32,32 0 0 1 64,0l0,256a32,32 0 0 1 -32,32z" />
      <path
        id="svg_9"
        d="m79.121508,945.070492a32,32 0 0 1 0,-45.248l247.936,-247.936a32,32 0 0 1 45.248,45.248l-247.936,247.936a32,32 0 0 1 -45.248,0z"
      />
    </svg>
    <svg
      v-else
      t="1590133734869"
      class="screenfull-svg"
      viewBox="0 0 1024 1024"
      version="1.1"
      xmlns="http://www.w3.org/2000/svg"
      p-id="1862"
      xmlns:xlink="http://www.w3.org/1999/xlink"
      width="20"
      height="20"
    >
      <path d="M928.512 959.744a32 32 0 0 1-32-32v-256a32 32 0 0 1 64 0v256a32 32 0 0 1-32 32z" fill="" p-id="1863" />
      <path
        d="M960.512 927.744a32 32 0 0 1-32 32h-256a32 32 0 0 1 0-64h256a32 32 0 0 1 32 32zM95.744 64.256a32 32 0 0 1 32 32v256a32 32 0 0 1-64 0v-256a32 32 0 0 1 32-32z"
        fill=""
        p-id="1864"
      />
      <path d="M63.744 96.256a32 32 0 0 1 32-32h256a32 32 0 0 1 0 64h-256a32 32 0 0 1-32-32z" fill="" p-id="1865" />
      <path d="M384.064 671.744a32 32 0 0 1-32 32h-256a32 32 0 0 1 0-64h256a32 32 0 0 1 32 32z" fill="" p-id="1866" />
      <path d="M352.064 639.744a32 32 0 0 1 32 32v256a32 32 0 0 1-64 0v-256a32 32 0 0 1 32-32z" fill="" p-id="1867" />
      <path
        d="M366.656 657.216a32 32 0 0 1 0 45.248L118.72 950.4a32 32 0 0 1-45.248-45.248l247.936-247.936a31.936 31.936 0 0 1 45.248 0z"
        fill=""
        p-id="1868"
      />
      <path d="M639.616 352.512a32 32 0 0 1 32-32h256a32 32 0 0 1 0 64h-256a32 32 0 0 1-32-32z" fill="" p-id="1869" />
      <path d="M671.616 384.512a32 32 0 0 1-32-32v-256a32 32 0 0 1 64 0v256a32 32 0 0 1-32 32z" fill="" p-id="1870" />
      <path
        d="M657.088 367.104a32 32 0 0 1 0-45.248l247.936-247.936a32 32 0 0 1 45.248 45.248l-247.936 247.936a32 32 0 0 1-45.248 0z"
        fill=""
        p-id="1871"
      />
    </svg>
  </div>
</template>

<script>
import screenfull from 'screenfull' // 引入screenfull
export default {
  name: 'Screenfull',
  props: {
    width: {
      type: Number,
      default: 22
    },
    height: {
      type: Number,
      default: 22
    },
    fill: {
      type: String,
      default: '#48576a'
    }
  },
  data() {
    return {
      isShow: true
    }
  },
  mounted() {
    if (screenfull.isEnabled) {
      screenfull.on('change', this.changeFullShow)
    }
  },
  destroyed() {
    screenfull.off('change', this.changeFullShow)
  },
  methods: {
    click() {
      if (screenfull.isEnabled) {
        screenfull.toggle()
      }
    },
    changeFullShow() {
      this.isShow = !screenfull.isFullscreen
    }
  }
}
</script>

<style scoped>
.screenfull-svg {
  width: 17px;
  height: 17px;
  cursor: pointer;
  fill: #606266;
  vertical-align: middle;
  margin-right: 15px;
}
</style>
