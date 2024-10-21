<template>
	<!-- #ifdef MP-WEIXIN || MP-ALIPAY-->
	<root-portal>
	<!-- #endif -->
	<!-- #ifdef H5 -->
	<teleport to="body" :disabled="!props.teleport">
	<!-- #endif -->

		<view
			@touchmove.prevent=""
			v-if="showMask"
			class="l-0 t-0"
			:style="[
				_inContent && !isNvue
					? { width: '100%', height: '100%', top: '0px', position: 'absolute' }
					: {
							width: sysinfo.width + 'px',
							height: sysinfo.height + 'px',
							top: 0 + 'px',
							position: 'fixed'
					  },
				zIndex ? { zIndex: zIndex } : ''
			]"
		>
			<view
				ref="overlay"
				:class="[
					bgColor_rp && !props.transprent && ani ? 'blurOnOpacity' : 'blurOffOpacity', 
				'overlay',
				store.tmuiConfig?.themeConfig.overflowBlur&&bgColor_rp && !props.transprent && ani?'blurOn':'',
				store.tmuiConfig?.themeConfig.overflowBlur&&bgColor_rp && !props.transprent && !ani?'blurOff':'',
				]"
				:style="[
					bgColor_rp && !props.transprent ? { backgroundColor: showMask ? bgColor_rp : '' } : '',
					_inContent && !isNvue ? { width: '100%', height: '100%' } : { width: sysinfo.width + 'px', height: sysinfo.height + 'px' },
					{ transitionDuration: props.duration + 'ms' }
				]"
			></view>
			<!-- #ifndef APP-NVUE -->
			<view
				@click.stop="closeByclick"
				:class="[
					align_rpx,
					' absolute flex flex-col  l-0 t-0 ',
					customClass,

					props.contentAnimation ? 'overlay' : '',
					props.contentAnimation && ani ? 'blurOnOpacity ' : '',
					props.contentAnimation && !ani ? 'blurOffOpacity overlay' : ''
				]"
				:style="[_inContent && !isNvue ? { width: '100%', height: '100%', top: '0px' } : { width: sysinfo.width + 'px', height: sysinfo.height + 'px' }, customCSSStyle]"
			>
				<slot></slot>
			</view>
			<!-- #endif -->
			<!-- #ifdef APP-NVUE -->
			<view
				@click.stop="closeByclick"
				:class="[align_rpx, ' absolute flex flex-col  l-0 t-0 ', customClass]"
				:style="[_inContent && !isNvue ? { width: '100%', height: '100%', top: '0px' } : { width: sysinfo.width + 'px', height: sysinfo.height + 'px' }, customCSSStyle]"
			>
				<slot></slot>
			</view>
			<!-- #endif -->
		</view>

	<!-- #ifdef MP-WEIXIN || MP-ALIPAY -->
	</root-portal>
	<!-- #endif -->
	<!-- #ifdef H5 -->
	</teleport>
	<!-- #endif -->

</template>
<script lang="ts" setup>
/**
 * 遮罩层
 * @description 遮罩层全屏弹出。
 */
import {
  getCurrentInstance,
  computed,
  ref,
  provide,
  inject,
  onMounted,
  onUnmounted,
  nextTick,
  watch,
  ComponentInternalInstance,
} from "vue";
import { cssstyle, tmVuetify } from "../../tool/lib/interface";
import { custom_props, computedClass, computedStyle } from "../../tool/lib/minxs";
import { useTmpiniaStore } from "../../tool/lib/tmpinia";
import { useWindowInfo } from '../../tool/useFun/useWindowInfo'

// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin("dom");
const animation = uni.requireNativePlugin("animation");
// #endif
const store = useTmpiniaStore();
const defaultBgColor = "rgba(0,0,0,0.24)";
// 混淆props共有参数
const props = defineProps({
  ...custom_props,
  // 内容的对齐方式的类
  align: {
    type: String,
    default: "flex-col-center-center",
  },
  //当前组件的主题。可以是颜色值，也可以是主题名称。
  bgColor: {
    type: String,
    default: "rgba(0,0,0,0.24)",
  },
  zIndex: {
    type: [Number, String],
    default: 999,
  },
  show: {
    type: Boolean,
    default: false,
  },
  overlayClick: {
    type: Boolean,
    default: true,
  },
  transprent: {
    type: [Boolean, String],
    default: false,
  },
  duration: {
    type: Number,
    default: 300,
  },
  contentAnimation: {
    type: Boolean,
    default: false,
  },
  /** 是否嵌入弹层，开启后将在它的父组件内执行弹层。 */
  inContent: {
    type: Boolean,
    default: false,
  },
  /** 是否使用teleport */
  teleport:{
    type: Boolean,
    default: true,
  },
});
const emits = defineEmits(["click", "open", "close", "update:show"]);
const proxy = getCurrentInstance()?.proxy ?? null;
//自定义样式：
const customCSSStyle = computedStyle(props);
//自定类
const customClass = computedClass(props);
const sysinfo = useWindowInfo();

const isAniing = ref(false);
let timids = uni.$tm.u.getUid(1);
let timerId = NaN;
const animationData = ref(null);
const showMask = ref(false);
const ani = ref(false);
onUnmounted(() => clearTimeout(timerId));
const align_rpx = computed(() => props.align);

const bgColor_rp = computed(() => {
  if (!props.bgColor || props.transprent) return "rgba(0,0,0,0)";
  return props.bgColor || defaultBgColor;
});
const _inContent = ref(props.inContent);
const isNvue = ref(false);
let timerIdth:any = NaN;
let timerIdth_flas = false;
// #ifdef APP-NVUE
_inContent.value = false;
isNvue.value = true;
// #endif
let parent: any = null;
onMounted(() => {
  if (!props.show) return;
  open(props.show);
});

function throttle(func: Function, wait = 500, immediate = true) {
  if (immediate) {
    if (!timerIdth_flas) {
      timerIdth_flas = true;
      // 如果是立即执行，则在wait毫秒内开始时执行
      typeof func === "function" && func();

      timerIdth = setTimeout(() => {
        timerIdth_flas = false;
      }, wait);
    }
  } else {
    if (!timerIdth_flas) {
      timerIdth_flas = true;
      // 如果是非立即执行，则在wait毫秒内的结束处执行
      timerIdth = setTimeout(() => {
        timerIdth_flas = false;
        typeof func === "function" && func();
      }, wait);
    }
  }
}

function close() {
  if (timerId) {
    clearTimeout(timerId);
    timerId = NaN;
  }
  
  open(false);
}

function closeByclick(e: Event) {

  try {
    e.stopPropagation();
    e.stopImmediatePropagation();
  } catch (e) {
    //TODO handle the exception
  }
  
  if (timerId) {
    clearTimeout(timerId);
    timerId = NaN;
  }
  
  emits("click", e);
  if (!props.overlayClick) return;
  open(false);
}

function open(off: boolean) {
  if (off == true) {
    uni.hideKeyboard();
  }
  // #ifdef APP-NVUE
  fadeInNvue(off);
  // #endif

  // #ifndef APP-NVUE
  fadeInVue(off);
  // #endif
}
function touchmove(e:TouchEvent){
try{
  e.preventDefault()
}catch(e){}
}
function fadeInNvue(off: boolean = false) {
  if (off == false) {
    if (showMask.value == off) return;
    var testEl = proxy?.$refs?.overlay;
    animation.transition(
      testEl,
      {
        styles: {
          backgroundColor: bgColor_rp.value,
          opacity: 0,
        },
        duration: props.duration || 1, //ms
        timingFunction: "linear",
        delay: 0, //ms
      },
      () => {
        showMask.value = off;
        emits("close");
        emits("update:show", false);
        // isAniing.vale = false;
      }
    );
  } else {
    showMask.value = off;
    emits("open");
    clearTimeout(timids);
    timids = setTimeout(function () {
      var testEl = proxy?.$refs.overlay;
      animation.transition(
        testEl,
        {
          styles: {
            backgroundColor: bgColor_rp.value,
            opacity: 1,
          },
          duration: props.duration || 1, //ms
          timingFunction: "linear",
          delay: 0, //ms
        },
        () => {}
      );
    }, 50);
  }
}
function fadeInVue(off = false) {
  if (showMask.value == off) return;
  throttle(
    function () {
		
      if (off == false) {
        ani.value = false;
        setTimeout(function () {
          showMask.value = off;
          emits("close");
          emits("update:show", false);
        }, props.duration + 10);
      } else {
        showMask.value = true;
        setTimeout(function () {
          ani.value = true;
        }, 10);
        emits("open");
        setTimeout(function () {
          emits("update:show", true);
        }, props.duration);
      }
    },
    props.duration+10,
    true
  );
}
watch(
  () => props.show,
  (newval) => {
    open(newval);
  }
);
defineExpose({
  close: close,
  open: open,
});
</script>

<style scoped="scoped">
.overlay {
	transition-timing-function: ease;
	transition-property: opacity;
	transition-delay: 0;
	opacity: 0;
}
.blur{
	
}
.blurOn {
	/* #ifndef APP-PLUS-NVUE */
	backdrop-filter: blur(2px);
	/* #endif */
	/* opacity: 1; */
}
.blurOff {
	/* #ifndef APP-PLUS-NVUE */
	backdrop-filter: blur(0px);
	/* #endif */
	/* opacity: 0; */
}

.blurOnOpacity {
	opacity: 1;
}
.blurOffOpacity {
	opacity: 0;
}
</style>
