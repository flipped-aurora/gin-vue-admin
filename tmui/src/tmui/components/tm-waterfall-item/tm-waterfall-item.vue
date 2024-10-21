<template>
	<view
		@click="onClick"
		ref="itemWall"
		class="absolute itemWall"
		:style="[
			!isPush ? { transform: 'translateX(-1000px)' } : '',
			isPush ? { transform: `translateX(${_nodeInfo.left}px) translateY(${_nodeInfo.top}px)` } : ''
		]"
	>
		<tm-sheet
			:margin="[0, 0]"
			:padding="[0, 0]"
			:round="props.round"
			:width="_width"
			unit="px"
			:color="props.color"
			_class="flex flex-col flex-col-top-start"
			:eventPenetrationEnabled="true"
		>
			<tm-image
				v-if="!imgerror && props.img"
				@click="onImgClick"
				:round="props.round"
				@load="imgLoadSuccess"
				@error="error"
				:src="props.img"
				unit="px"
				:height="_nodeInfo.imgHeight"
				:width="_nodeInfo.imgWidth"
			></tm-image>
			<view
				class="flex flex-row flex-row-center-center"
				:userInteractionEnabled="false"
				v-if="imgerror"
				:style="[imgerror ? { height: _nodeInfo.imgWidth + 'px', width: _nodeInfo.imgWidth + 'px' } : '']"
			>
				<tm-icon name="tmicon-exclamation-circle"></tm-icon>
			</view>

			<view class="flex flex-col flex-1 flex-col-top-start">
				<slot></slot>
			</view>
		</tm-sheet>
	</view>
</template>

<script lang="ts" setup>
/**
 * 瀑布流项目
 * @description 不可单独使用，只可放置在tm-waterfall中。
 * @slot default 默认插槽，可以放置任意内容，但不允许异步更改高度，否则瀑布失效。
 * @method img-click 封面图被点击时触发。如果不提供封面图本事件不触发。
 */
import { onMounted, nextTick, inject, computed, Ref, ref, getCurrentInstance, toRaw } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmImage from '../tm-image/tm-image.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import { itemParenSG } from '../tm-waterfall/interface'

// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
// #endif
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['img-click', 'click'])
const imgerror = ref(false)
const props = defineProps({
	//封面图片。
	img: {
		type: String,
		default: ''
	},
	color: {
		type: String,
		default: 'white'
	},
	round: {
		type: Number,
		default: 4
	}
})
const _width: ComputedRef<number> = inject(
	'tmWaterFallItemRealWidth',
	computed(() => uni.upx2px(_width.value))
)

let _nodeInfo: Ref<itemParenSG> = ref({
	id: uni.$tm.u.getUid(2), //子组件id，唯一。
	width: _width.value,
	height: _width.value,
	imgWidth: _width.value,
	imgHeight: _width.value,
	bottom: 0,
	index: NaN,
	top: 0,
	left: 0
})
const isimgLoad = ref(props.img ? false : true)
const _parentComs = getParent()
const isPush = ref(false) //是否已经推到父。
if (!_parentComs) {
	console.error('请不要单独使用此组件，请放置在:tm-waterfall内；')
}
_parentComs.sumTotal(_nodeInfo.value.id)

onMounted(() => {
	nextTick(() => {
		if (isimgLoad.value === true) {
			//没有异步内容时，直接就是加载渲染完成。
			nvuegetClientRect()
		}
	})
})

function imgLoadSuccess(e) {
	const { width, height } = e.detail
	//计算缩放的宽和高。
	let _w = _width.value
	let _height = _w / (width / height)
	_nodeInfo.value = { ..._nodeInfo.value, imgWidth: _w, imgHeight: _height }
	setTimeout(() => {
		nextTick(() => nvuegetClientRect())
	}, 50)
}
function error() {
	const wx_width = _width.value
	const wx_height = _width.value
	//计算缩放的宽和高。
	let _w = _width.value
	let _height = _w / (wx_width / wx_height)
	console.log(wx_height)
	_nodeInfo.value = { ..._nodeInfo.value, imgWidth: wx_width, imgHeight: wx_width }
	setTimeout(() => {
		nextTick(() => nvuegetClientRect())
		imgerror.value = true
	}, 50)
}
function getParent() {
	//父级方法。
	let parent = proxy.$parent

	while (parent) {
		if (parent?.parentNameId == 'tmWaterfallId' || !parent) {
			break
		} else {
			parent = parent?.$parent ?? undefined
		}
	}
	return parent
}
function nvuegetClientRect() {
	nextTick(function () {
		// #ifdef APP-PLUS-NVUE
		dom.getComponentRect(proxy.$refs.itemWall, function (res) {
			if (res?.size) {
				if (res.size.height == 0 && res.size.width == 0) {
					nvuegetClientRect()
				} else {
					isimgLoad.value = true
					const { width, height } = res.size
					_nodeInfo.value = {
						..._nodeInfo.value,
						height: height
					}
					if (isPush.value === false && isimgLoad.value) {
						pushKey()
						isPush.value = true
					}
				}
			}
		})
		// #endif
		// #ifndef APP-PLUS-NVUE
		uni.createSelectorQuery()
			.in(proxy)
			.select('.itemWall')
			.boundingClientRect((res) => {
				if (res.height === 0 && res.width === 0) {
					nvuegetClientRect()
				} else {
					isimgLoad.value = true
					const { width, height } = res
					_nodeInfo.value = {
						..._nodeInfo.value,
						height: height
					}
					if (isPush.value === false && isimgLoad.value) {
						pushKey()
						isPush.value = true
					}
				}
			})
			.exec()
		// #endif
	})
}
async function pushKey() {
	if (_parentComs) {
		let pos = await _parentComs.pushKey(toRaw(_nodeInfo.value))
		_nodeInfo.value = pos
	}
}
function onImgClick(e: TouchEvent) {
	emits('img-click', _nodeInfo.value)
}
/** 整个项目被点击。 */
function onClick(e: TouchEvent) {
	emits('click', _nodeInfo.value)
}
</script>
