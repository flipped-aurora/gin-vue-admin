<template>
	<tm-sheet
		:margin="props.margin"
		:padding="[24, 0]"
		:width="_width"
		:height="props.height"
		:color="props.color"
		:_style="props._style"
		:followTheme="props.followTheme"
		:dark="props.dark"
		:round="props.round"
		:shadow="props.shadow"
		:outlined="props.outlined"
		:border="props.border"
		:borderStyle="props.borderStyle"
		:borderDirection="props.borderDirection"
		:text="props.text"
		:transprent="props.transprent"
		:linear="props.linear"
		:linearDeep="props.linearDeep"
		_class="flex flex-row flex-between"
	>
		<view v-if="_icon" class="flex flex-row flex-row-center-start" style="width: 50rpx">
			<tm-icon :font-size="28" :name="_icon" style="line-height: normal"></tm-icon>
		</view>
		<view class="flex flex-1 flex-row overflow" style="width: 0px">
			<view
				ref="content"
				:style="{
					animationDuration: _duration + 's',
					paddingLeft: (isNvue ? 0 : _Left) + 'rpx'
				}"
				class="aniRow flex-row flex-row-center-start"
				id="wrap"
			>
				<view class="flex-row flex-row-center-start" v-for="(item, index) in _list" :key="index">
					<tm-text
						_class="pl-24 nowrap"
						:font-size="props.fontSize"
						:color="props.fontColor"
						@click="emits('click', index)"
						:label="item"
					></tm-text>
				</view>
			</view>
		</view>
		<view v-if="props.showRight" class="flex flex-row flex-row-center-end" style="width: 40rpx">
			<tm-icon :font-size="24" name="tmicon-angle-right" style="line-height: normal"></tm-icon>
		</view>
	</tm-sheet>
</template>

<script lang="ts" setup>
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import { custom_props } from '../../tool/lib/minxs'
import { ComponentInternalInstance, computed, getCurrentInstance, nextTick, onMounted, onUpdated, PropType, ref } from 'vue'
// #ifdef APP-NVUE
const animation = uni.requireNativePlugin('animation')
const dom = uni.requireNativePlugin('dom')
// #endif
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits<{
	(e: 'click', index: number): void
}>()
const props = defineProps({
	...custom_props,
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [32, 24]
	},
	transprent: {
		type: Boolean,
		default: false
	},
	text: {
		type: Boolean,
		default: true
	},
	width: {
		type: Number,
		default: 726
	},
	height: {
		type: Number,
		default: 70
	},
	fontSize: {
		type: Number,
		default: 26
	},
	color: {
		type: String,
		default: 'primary'
	},
	fontColor: {
		type: String,
		default: ''
	},
	icon: {
		type: String,
		default: 'tmicon-info-circle'
	},
	//是否显示右边图标。
	showRight: {
		type: Boolean,
		default: false
	},
	shadow: {
		type: Number,
		default: 0
	},
	list: {
		type: [Array, String, Object] as PropType<Array<string | object> | string | object>,
		default: () => ''
	},
	rangeKey: {
		type: String,
		default: 'text'
	},
	speed: {
		type: Number,
		default: 60
	},
	// duration属性在3.8.3后删除,存在这里的意义是向下兼容之前的老版本.
	duration: {
		type: Number,
		default: 0
	},
	border: {
		type: Number,
		default: 0
	}
})
const _icon = computed(() => props.icon)
const _list = computed<Array<string>>(() => {
	let listData: string[] = []
	if (typeof props.list === 'string') {
		listData.push(props.list)
	}
	if (typeof props.list === 'object') {
		// @ts-expect-error
		listData.push(props.list[props.rangeKey])
	}
	if (Array.isArray(props.list)) {
		props.list.forEach((el) => {
			if (typeof el == 'string') {
				listData.push(el)
			} else if (typeof el == 'object' && props.rangeKey) {
				// @ts-expect-error
				listData.push(el[props.rangeKey])
			}
		})
	}
	return listData
})

const _width = computed(() => Math.ceil(props.width - props.margin[0] * 2 - 24))

const _duration = ref(0)

const _Left = computed(() => {
	if (_icon.value !== '') {
		return _width.value - 124
	}
	return _width.value - 84
})

const isNvue = ref(false)
//  #ifdef APP-NVUE
isNvue.value = true
// #endif
onMounted(() => {
	//  #ifdef APP-NVUE
	setTimeout(function () {
		getNuveWrapWidth()
		nvueani()
	}, 200)
	// #endif
	//  #ifndef APP-NVUE
	getNoNvueWrapWidth()
	// #endif
})
onUpdated(() => {})

// 非nvue下获取宽度等信息.
function getNoNvueWrapWidth() {
	uni.createSelectorQuery()
		.in(proxy)
		.select('#wrap')
		.boundingClientRect((res) => {
			let totalWidth = (Number(res.width || 0) + _width.value) / props.speed
			_duration.value = Math.ceil(totalWidth)
		})
		.exec()
}
function getNuveWrapWidth() {
	var testEl = proxy?.$refs.content
	return new Promise((resj) => {
		dom.getComponentRect(testEl, (res) => {
			if (res?.size) {
				let totalWidth = (Number(res.size.width || 0) + _width.value) / props.speed
				_duration.value = Math.ceil(totalWidth)
				resj()
			}
		})
	})
}

function nvueani() {
	var testEl = proxy?.$refs.content
	getNuveWrapWidth().then(() => {
		animation.transition(
			testEl,
			{
				styles: {
					transform: 'translateX(-100%)',
					transformOrigin: 'center center'
				},
				duration: _duration.value * 1000, //ms
				timingFunction: 'linear',
				delay: 0 //ms
			},
			() => {
				animation.transition(
					testEl,
					{
						styles: {
							transform: 'translateX(' + uni.upx2px(_Left.value) + 'px)',
							transformOrigin: 'center center'
						},
						duration: 0, //ms
						timingFunction: 'linear',
						delay: 1 //ms
					},
					() => {
						nvueani()
					}
				)
			}
		)
	})
}
</script>

<style scoped>
/* #ifndef APP-NVUE */
.aniRow {
	animation-name: roll;
	animation-timing-function: linear;
	animation-iteration-count: infinite;
}
@keyframes roll {
	0% {
		transform: translateX(0);
	}
	100% {
		transform: translateX(-100%);
	}
}
/* #endif */
</style>
