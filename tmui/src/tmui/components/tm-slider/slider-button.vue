<template>
	<!-- #ifdef MP-WEIXIN -->
	<!-- <view 
    data-text="hehe"
    @touchstart.stop="test.touchstart" @mousedown.stop="movestart" @touchmove.stop="test.touchmove" @mousemove.stop="moveing" @touchend.stop="moveend" @mouseup.stop="moveend" @mouseleave.stop="moveend" 
    class="absolute movable" :style="[
    props.direction == 'horizontal'?{width:props.size+'rpx',height:props.size+'rpx',transform:`translateX(${_x}px)`,top:'0px'}:'',
    props.direction == 'vertical'?{width:props.size+'rpx',height:props.size+'rpx',transform:`translateY(${_x}px)`,left:0+'rpx',top:'0px'}:'',
    ]">
        <tm-sheet :followTheme="props.followTheme" :shadow="1" :border="4" :userInteractionEnabled="false" :color="props.color" :round="24" :width="props.size" :height="props.size" :margin="[0,0]" :padding="[0,0]"></tm-sheet>
    </view> -->

	<!-- #endif -->

	<view
		@touchstart.stop="movestart"
		@mousedown.stop="movestart"
		@touchmove.stop="moveing"
		@mousemove.stop="moveing"
		@touchend.stop="moveend"
		@mouseup.stop="moveend"
		@mouseleave.stop="moveend"
		class="absolute"
		:style="[
			props.direction == 'horizontal'
				? {
						width: props.size + 'rpx',
						height: props.size + 'rpx',
						transform: `translateX(${_x}px)`,
						top: '0px'
				  }
				: '',
			props.direction == 'vertical'
				? {
						width: props.size + 'rpx',
						height: props.size + 'rpx',
						transform: `translateY(${_x}px)`,
						left: 0 + 'rpx',
						top: '0px'
				  }
				: ''
		]"
	>
	
		<tm-sheet
			:followTheme="props.followTheme"
			:shadow="3"
			:border="4"
			:userInteractionEnabled="false"
			:color="props.color"
			:border-color="store.tmStore.dark?'black':'white'"
			:round="24"
			:width="props.size"
			:height="props.size"
			:margin="[0, 0]"
			:padding="[0, 0]"
		></tm-sheet>
	</view>
</template>

<script lang="ts" setup>
import { useTmpiniaStore } from '@/tmui/tool/lib/tmpinia';
import { computed } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'

const store = useTmpiniaStore();

const emits = defineEmits(['movestart', 'moveing', 'moveend'])

const props = defineProps({
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	size: {
		type: Number,
		default: 40
	},
	x: {
		type: Number,
		dfault: 0
	},
	color: {
		type: String,
		default: 'primary'
	},
	/**
	 * 方向
	 * horizontal:水平,
	 * vertical:竖向。
	 */
	direction: {
		type: String,
		default: 'vertical'
	}
})
const _x = computed(() => props.x)
let timerId = NaN
let timerId2 = NaN
function debounce(func: Function, wait = 500, immediate = false) {
	// 清除定时器
	if (!isNaN(timerId)) clearTimeout(timerId)
	// 立即执行，此类情况一般用不到

	if (immediate) {
		var callNow = !timerId
		timerId = setTimeout(() => {
			timerId = NaN
		}, wait)

		if (callNow) typeof func === 'function' && func()
	} else {
		// 设置定时器，当最后一次操作后，timeout不会再被清除，所以在延时wait毫秒后执行func回调方法
		timerId = setTimeout(() => {
			typeof func === 'function' && func()
		}, wait)
	}
}

function movestart(e: TouchEvent | MouseEvent) {
	let etype = e.type.toLocaleLowerCase()
	let ex = 0
	let ey = 0
	if (etype == 'mousedown') {
		ex = e.pageX
		ey = e.pageY
	} else if (etype == 'touchstart') {
		ex = e.changedTouches[0].pageX
		ey = e.changedTouches[0].pageY
	}
	emits('movestart', { x: ex, y: ey })
}
function moveing(e: TouchEvent | MouseEvent) {
	let etype = e.type.toLocaleLowerCase()
	let ex = 0
	let ey = 0
	if (etype == 'mousemove') {
		ex = e.pageX
		ey = e.pageY
	} else if (etype == 'touchmove') {
		ex = e.changedTouches[0].pageX
		ey = e.changedTouches[0].pageY
	}
	debounce(
		() => {
			emits('moveing', { x: ex, y: ey })
		},
		5,
		false
	)
	e.preventDefault()
	e.stopPropagation()
}
function moveend(e: TouchEvent | MouseEvent) {
	let etype = e.type.toLocaleLowerCase()
	let ex = 0
	let ey = 0
	if (etype == 'mouseup' || etype == 'mouseleave') {
		ex = e.pageX
		ey = e.pageY
	} else if (etype == 'touchend') {
		ex = e.changedTouches[0].pageX
		ey = e.changedTouches[0].pageY
	}
	emits('moveend', { x: ex, y: ey })
}
</script>
<script lang="ts">
export default {
	data() {
		return {
			
		}
	},
	mounted() {}
}
</script>

<script module="test" lang="wxs">
// 暂时用不了.uni官方不支持setup中使用wsx,使用了可以,但无法使用模板变量.
	var startX = 0
	var startY = 0
	var lastLeft = 0; var lastTop = 0
    function setDefaultXY(x,y){
        // lastLeft = x;
       return x
    }
	function touchstart(event, ins) {
        console.log(JSON.stringify(ins.selectComponent('.movable')))

	  var touch = event.touches[0] || event.changedTouches[0]
	  startX = touch.pageX
	  startY = touch.pageY
	}
	function touchmove(event, ins) {
	  var touch = event.touches[0] || event.changedTouches[0]
	  var pageX = touch.pageX
	  var pageY = touch.pageY
	  var left = pageX - startX + lastLeft
	  var top = pageY - startY + lastTop
	  startX = pageX
	  startY = pageY
	  lastLeft = left
	  lastTop = top

	  ins.selectComponent('.movable').setStyle({
	    left: left + 'px',
	    // top: top + 'px'
	  })
		return false
	}
    function touchend(event, ins) {

    }
	module.exports = {
        setDefaultXY:setDefaultXY,
		lastLeft: lastLeft,
	  touchstart: touchstart,
	  touchmove: touchmove
	}
</script>
