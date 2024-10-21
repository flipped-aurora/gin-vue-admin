<template>
	<!-- #ifdef MP -->
	<view @touchstart="touch.startDrag" @touchmove.stop="touch.onDrag" @touchend="touch.endDrag" :data-prop="towxsShareData">
	<!-- #endif -->
	<!-- #ifdef H5||APP-VUE -->
	<view @touchstart="touch.startDrag" @touchmove="touch.onDrag" @touchend="touch.endDrag" :data-prop="towxsShareData">
	<!-- #endif -->
	<!-- #ifdef APP-NVUE -->
	<view  >
	<!-- #endif -->
		<view @touchstart="touchstart" :style="{ left: _offset[0] + 'rpx', top: _offset[1] + 'rpx' }"  class="div" id="adsorb" ref="adsorb">
			<view
				:eventPenetrationEnabled="true"
				:style="{
					width: props.width + 'rpx',
					height: props.height + 'rpx'
				}"
			>
				<slot></slot>
			</view>
		</view>
	</view>
	
</template>
<!-- #ifndef APP-NVUE -->
<script module="touch" lang="wxs" src="./touch.wxs"></script>
<!-- #endif -->
<script lang="ts" setup>
import { ref, inject, computed, unref, PropType,getCurrentInstance } from 'vue';
// #ifdef APP-NVUE
var dom = weex.requireModule("dom");
const Binding = uni.requireNativePlugin("bindingx");
const animation = uni.requireNativePlugin("animation");
const proxy = getCurrentInstance()?.proxy??null;
// #endif
const props = defineProps({
	/** 是否吸附边缘，关闭可以任意托动组件，开启拖动只会吸附在两边。 */
	adsorb: {
		type: Boolean,
		default: true
	},
	/** 开启吸附后，吸附到边缘的动画时间，单位ms */
	duration: {
		type: Number,
		default: 600
	},
	width: {
		type: Number,
		default: 100
	},
	height: {
		type: Number,
		default: 100
	},
	/** 默认的位置 */
	offset: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	/** 吸附的偏移量，比如向左时，是吸附侧边0还是再往左偏移多少 */
	adsorbX: {
		type: Number,
		default: 0
	}
});
const sysinfo = inject(
	'tmuiSysInfo',
	computed(() => {
		return {
			bottom: 0,
			height: 750,
			width: uni.upx2px(750),
			top: 0,
			isCustomHeader: false,
			sysinfo: null
		};
	})
);
const _offset = computed(() => props.offset);
let bindxToken:any = null;
const towxsShareData = ref({
	adsorb: props.adsorb,
	sys: sysinfo.value,
	adsorbX: props.adsorbX,
	duration: props.duration
});
let position = {x:0,y:0};
function getEl(el: any) {
  if (typeof el === "string" || typeof el === "number") return el;
  if (WXEnvironment) {
    return el.ref;
  } else {
    return el instanceof HTMLElement ? el : el.$el;
  }
}
function spinNvueAniEnd(start: number,end:number, duration = props.duration) {
  // #ifdef APP-NVUE
  if (!proxy?.$refs?.adsorb) return;
  animation.transition(
    proxy?.$refs.adsorb,
    {
      styles: {
        transform: `translate(${start}px,${end}px)`,
        transformOrigin: "center center",
      },
      duration: duration, //ms
      timingFunction: "cubicBezier(0.18, 0.89, 0.32, 1)",
      delay: 0, //ms
    },
    () => {
      
    }
  );

  // #endif
}

function touchstart(e: TouchEvent) {
  
  // #ifdef APP-NVUE
  if (!proxy?.$refs?.adsorb) return;

  let icon = getEl(proxy?.$refs.adsorb);
  let icon_bind = Binding.bind(
    {
      anchor: icon,
      eventType: "pan",
      props: [
        {
          element: icon,
          property: "transform.translateX",
          expression: `x+${position.x}`,
        },
		{
		  element: icon,
		  property: "transform.translateY",
		  expression: `y+${position.y}`,
		},
      ],
    },
    function (res) {
      if (res.state == "end") {
		position.x+=res.deltaX
		position.y+=res.deltaY
		if(towxsShareData.value.adsorb){
			dom.getComponentRect(proxy?.$refs.adsorb, function (res:UniApp.NodeInfo|UniApp.NodeField) {
			  if (res?.size) {
				let left=0
				let top=0
			    let rect = res.size;
				let x = uni.upx2px(props.offset[0]);
				let y = uni.upx2px(props.offset[1]);
				if(Math.abs((rect.left+rect.width/2))<=towxsShareData.value.sys.width/2){
					left = x - Math.abs(position.x)
					position.x = -left+position.x - towxsShareData.value.adsorbX
				}else{
					left =  towxsShareData.value.sys.width - rect.right;
					position.x =position.x + left + towxsShareData.value.adsorbX
				}
				if(rect.bottom>=towxsShareData.value.sys.height){
					position.y =position.y - (rect.bottom-towxsShareData.value.sys.height)
				}
				if(rect.top<=0){
					position.y =position.y + Math.abs(rect.top)
				}
				
				
				spinNvueAniEnd(position.x,position.y);
			  }
			});
		}
        
      } else if (res.state == "start") {
        
      }
    }
  );
  bindxToken = icon_bind.token;
  // #endif
}


defineExpose({});
</script>

<style scoped>
.div {
	position: fixed;
}
</style>
