<template>
	<view
		class="tm-segtab relative flex flex-col"
		:class="[`round-${props.round}`]"
		ref="tm-segtab"
		:style="{ width: wh.outerWidth + 'rpx', height: wh.outerHeight + 'rpx' }"
	>
		<tm-sheet
			:round="props.round"
			:linear="props.linear"
			:linear-deep="props.linearDeep"
			:no-level="true"
			:color="props.bgColor"
			darkBgColor="rgba(255,255,255,0.06)"
			:height="wh.outerHeight"
			_class="flex-row relative overflow flex-1"
			class="flex-1"
			:padding="[0]"
			:margin="[0]"
			:border="0"
		>
			<!-- #ifdef APP-NVUE -->
			<view
				v-if="_cId !== '' && leftWidth >= 0"
				ref="tmBgEl"
				class="relative flex flex-row"
				:style="[
					{
						width: leftWidth + 'px',
						height: wh.innerHeight + 'rpx',
						top: `${props.gutter}rpx`,
						left: '0px'
					}
				]"
			>
				<tm-sheet
					:follow-dark="props.followDark"
					:round="props.round"
					class="flex-1"
					_class="flex-1"
					:color="props.color"
					:margin="[0]"
					:padding="[0]"
				></tm-sheet>
			</view>
			<!-- #endif -->
			<!-- #ifndef APP-NVUE -->
			<view
				v-if="_cId !== ''"
				class="absloute flex flex-row bgbtnpos tm-segtab relative"
				:style="[
					{
						transform: 'translateX(' + (leftPos + wh.gutterpx) + 'px)',
						width: leftWidth + 'px',
						height: wh.innerHieghtWeb,
						top: `${props.gutter}rpx`,
						left:'0px'
					}
				]"
			>
				<tm-sheet
					:follow-dark="props.followDark"
					:round="props.round"
					class="flex-1 flex flex-row"
					parenClass="flex-1"
					_class="flex-1 flex flex-row"
					:color="props.color"
					:margin="[0, 0]"
					:padding="[0, 0]"
					:border="0"
				></tm-sheet>
			</view>
			<!-- #endif -->
			<view
				class="absolute flex flex-row flex-row-center-start tm-segtab"
				:style="{
					width: `${wh.innerWidth}rpx`,
					height: wh.innerHeight + 'rpx',
					padding: `${props.gutter}rpx`,
					top: props.gutter+'rpx'
				}"
			>
				<view
					@click="itemClick(index, item.id)"
					:ref="'tab_'"
					:class="['tab' + index]"
					class="flex flex-row flex-row-center-center px-16 nowrap flex-1 "

					v-for="(item, index) in _list"
					:style="{ margin: `${props.gutter}rpx 0`}"
					:key="index"
				>
					<tm-text
						_style="transition: color 0.3s;"
						:color="item.id === _cId ? props.activeColor : ''"
						:lineHeight="0"
						:font-size="props.fontSize"
						:userInteractionEnabled="false"
						:label="item.text"
					></tm-text>
				</view>
			</view>
		</tm-sheet>
	</view>
</template>
<script lang="ts" setup>
/**
 * 分段器选项卡
 */
import { computed, PropType, toRaw, getCurrentInstance, ref, onMounted, nextTick, watch, Ref, inject } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import { listitem } from './interface'
import { custom_props } from '../../tool/lib/minxs'
// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
const animation = uni.requireNativePlugin('animation')
// #endif
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['update:modelValue', 'change', 'click'])

const props = defineProps({
	...custom_props,
	round: {
		type: Number,
		default: 2
	},
	width: {
		type: Number,
		default: 600
	},
	height: {
		type: Number,
		default: 64
	},
	gutter: {
		type: Number,
		default: 4
	},
	list: {
		type: Array as PropType<Array<string | listitem>>,
		default: () => [],
		required: true
	},
	//v-model可以是index索引也可是对象id
	modelValue: {
		type: [Number, String],
		default: 0
	},
	//如果想以字段id来达到index选中效果。需要list为对象，并且提供唯一标识id字段。
	defaultValue: {
		type: [Number, String],
		default: 0
	},
	//在点击切换之前执行，返回false阻止切换，可以是Promise
	beforeChange: {
		type: [Function, Boolean],
		default: () => false
	},
	color: {
		type: String,
		default: 'white'
	},
	bgColor: {
		type: String,
		default: 'grey-3'
	},
	fontSize: {
		type: Number,
		default: 24
	},
	//被选中后的文字色
	activeColor: {
		type: String,
		default: 'primary'
	}
})
const leftPos = ref(-100)
const leftWidth = ref(0)
let timid: any = uni.$tm.u.getUid()
const _list = computed(() => {
	let templist = []
	for (let i = 0, len = props.list.length; i < len; i++) {
		let al: listitem = { text: '', id: i }
		let el = props.list[i]
		if (typeof el == 'string' || typeof el == 'number') {
			al.text = el
		} else if (typeof el == 'object') {
			al.text = el?.text ?? ''

			if (typeof el?.id != 'undefined') {
				al.id = el['id']
			}
		}
		templist.push(al)
	}

	return templist
})
const firstRender = ref(true)
//当前值。
const _cId: Ref<string | number> = ref(props.defaultValue ?? 0)
const _blackValue = _cId.value
//如果list提供的是对象，想以id来选中定位，而不是inde索引则需要转换。
function zhunhuanid(val: string | number) {
	let index = _list.value.findIndex((el) => el.id == val)
	return index
}
// 外框宽度
const wh = computed(() => {
	let iw = props.width - props.gutter
	// #ifdef APP-NVUE
	iw = props.width
	// #endif
	// console.log((uni.upx2px(props.height)-uni.upx2px(props.height - props.gutter * 2))/2,uni.upx2px(props.gutter))
	let inheight = ''
	// #ifndef APP-NVUE
	inheight = `calc(100% - ${props.gutter * 2}rpx)`
	// #endif
	return {
		outerWidth: props.width,
		outerHeight: props.height,
		innerWidth: iw,
		innerHeight: props.height - props.gutter * 2,
		innerHieghtWeb:inheight,
		gutterpx: (uni.upx2px(props.height)-uni.upx2px(props.height - props.gutter * 2))/2,
		
	}
})

async function itemClick(index: number, id: number | string) {
	emits('click', index)
	if (typeof props.beforeChange === 'function') {
		uni.showLoading({ title: '...', mask: true })
		let p = await props.beforeChange(index)
		if (typeof p === 'function') {
			p = await p(index)
		}
		uni.hideLoading()
		if (!p) return
	}
	if (_cId.value === id) return
	_cId.value = id
	getDomRectBound(index)
	emits('change', _cId.value, toRaw(_list.value[index]))
	emits('update:modelValue', _cId.value)
	initPos()
}
watch(
	[() => props.list],
	() => {
		initPos()
	},
	{ deep: true }
)
watch(
	[ () => props.modelValue],
	() => {
		if(_cId.value == props.modelValue) return;
		_cId.value = props.modelValue
		initPos()
	}
)
onMounted(() => {
	initPos()
})
//定位背景按钮位置。
function initPos() {
	let indexel = _list.value.findIndex((el) => el.id === _cId.value)
	clearTimeout(timid)
	let timerdur = 150
	// #ifndef APP-NVUE
	timerdur = 50
	// #endif

	timid = setTimeout(() => {
		nextTick(() => getDomRectBound(indexel))
	}, timerdur)
}
function getEl(el) {
	if (typeof el === 'string' || typeof el === 'number') return el
	if (WXEnvironment) {
		return el.ref
	} else {
		return el instanceof HTMLElement ? el : el.$el
	}
}
function getDomRectBound(idx: number) {
	// #ifdef APP-NVUE
	dom.getComponentRect(proxy?.$refs['tm-segtab'], function (PARENAREDS) {
		if (PARENAREDS?.size) {
			let parentleft = Math.floor(PARENAREDS.size.left ?? 0)
			dom.getComponentRect(proxy?.$refs['tab_'][idx], function (res) {
				if (res?.size) {
					const { left, top, width } = res.size
					let domx = getEl(proxy?.$refs['tmBgEl'])
					leftWidth.value = Math.ceil(width ?? 0)
					leftPos.value = Math.ceil((left ?? 0) - uni.upx2px(props.gutter) - parentleft)
					animation.transition(
						proxy?.$refs['tmBgEl'],
						{
							styles: {
								transform: 'translateX(' + (leftPos.value + wh.value.gutterpx) + 'px)'
							},
							duration: firstRender.value ? 1 : 200, //ms
							timingFunction: 'ease',
							delay: 0 //ms
						},
						() => {
							firstRender.value = false
						}
					)
				}
			})
		}
	})
	// #endif
	// #ifndef APP-NVUE
	uni.createSelectorQuery()
		.in(proxy)
		.select('.tm-segtab')
		.boundingClientRect((nodeParent) => {
			let parentleft = nodeParent?.left ?? 0
			uni.createSelectorQuery()
				.in(proxy)
				.select('.tab' + idx)
				.boundingClientRect((node) => {
					if (!node) return
					leftPos.value = (node?.left ?? 0) - uni.upx2px(props.gutter) - parentleft
					leftWidth.value = (node?.width ?? 0) + +uni.upx2px(props.gutter)
				})
				.exec()
		})
		.exec()
	// #endif
}
</script>
<style scoped>
.bgbtnpos {
	transition-timing-function: cubic-bezier(0.18, 0.89, 0.32, 1);
	transition-duration: 0.4s;
	transition-property: left, width, transform;
	transition-delay: 0s;
}
.tm-segtab {
	display: flex;
	/* #ifndef APP-NVUE */
	box-sizing: border-box;
	/* #endif */
}

</style>
