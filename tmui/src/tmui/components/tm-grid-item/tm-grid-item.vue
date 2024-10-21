<template>
	<view :style="wkStyle">
		<tm-sheet
			:color="props.bgColor"
			:text="props.text"
			:border="0"
			hover-class="opacity-6"
			:transprent="props.transprent"
			:height="props.height"
			:width="_colWidth - 0.5"
			:margin="[0, 0]"
			:padding="[0, 0]"
			_class="flex-col flex"
			@click="onClick"
		>
			<view class="flex-1 flex flex-col-center-center">
				<tm-badge
					:top="0"
					:userInteractionEnabled="true"
					:fontSize="20"
					:dot="props.dot"
					:count="props.count"
					:max-count="props.maxCount"
					:icon="props.icon"
					:color="props.color"
				>
					<view class="flex-col flex-col-center-center flex px-10 my-5" :class="[props.dot ? '' : 'py-5 ']">
						<slot></slot>
					</view>
				</tm-badge>
			</view>
		</tm-sheet>
	</view>
</template>
<script lang="ts" setup>
/**
 * 九宫格子项目
 * @description 注意，它只能放置在tm-grid中，且不能嵌套tmg-grid
 * @slot 默认插槽 任意内容
 */
import { computed, inject, provide, Ref, ref, watchEffect, getCurrentInstance, onBeforeUnmount, watch, nextTick, onMounted } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmBadge from '../tm-badge/tm-badge.vue'
import { cssstyle, tmVuetify } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()
const emits = defineEmits(['click'])
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	...custom_props,
	//项目的高度。如果提供为0，就表示自动高度。
	height: {
		type: Number,
		default: 100
	},
	transprent: {
		type: Boolean,
		default: true
	},
	dot: {
		type: [Boolean, String],
		default: false
	},
	icon: {
		type: [String],
		default: ''
	},
	//如果count为数字时，显示数字角标，如果为string是显示文本角标。
	count: {
		type: [Number, String],
		default: 0
	},
	maxCount: {
		type: [Number, String],
		default: 999
	},
	bgColor: {
		type: String,
		default: 'white'
	},
	// dot的主题色
	color: {
		type: String,
		default: 'red'
	},
	//如果提供了链接，当点击项目时自动跳转页面。
	url: {
		type: String,
		default: ''
	}
})
interface arrayid {
	id: string | number
	type: string | number
}
// 设置响应式全局组件库配置表。
const tmcfg = computed(() => store.tmStore)
//是否暗黑模式。
const isDark = computed(() => computedDark(props, tmcfg.value))
//计算主题
const tmcomputed = computed<cssstyle>(() => {
	return computedTheme({ ...props, color: props.bgColor }, isDark.value, tmcfg.value)
})

const _colWidth = inject('tmGridItemWidth', 0)
const _tmGridshowBorder = inject(
	'tmGridshowBorder',
	computed(() => false)
)
const tmGridshowCachList = inject(
	'tmGridshowCachList',
	computed<Array<arrayid>>(() => [])
)
const uid: Ref<arrayid> = ref({
	id: uni.$tm.u.getUid(1),
	type: ''
})
//父级方法。
let parentFormItem: any = proxy?.$parent
while (parentFormItem) {
	if (parentFormItem?.keyName == 'tmGrid' || !parentFormItem) {
		break
	} else {
		parentFormItem = parentFormItem?.$parent ?? undefined
	}
}
onMounted(() => {
	if (parentFormItem?.pushKey) {
		parentFormItem.pushKey(uid.value)
	}
}),
	onBeforeUnmount(() => {
		parentFormItem.delKey(uid.value)
	})
let wkStyle: any = ref(`width:${_colWidth}'rpx'`)

watch(
	[tmGridshowCachList, _tmGridshowBorder],
	() => {
		nextTick(() => setStyleFun())
	},
	{ deep: true }
)

function setStyleFun() {
	let ar = tmGridshowCachList.value.filter((el) => el.id == uid.value.id)
	if (ar.length == 1) {
		uid.value = ar[0]
	}
	if (!_tmGridshowBorder.value) {
		wkStyle.value = `box-sizing: border-box;border:0rpx solid rgba(0,0,0,0);width:${_colWidth}rpx`
		return
	}

	if (uid.value.type == 1) {
		wkStyle.value = `box-sizing: border-box;border:1rpx solid ${tmcomputed.value.border};width:${_colWidth - 1}rpx`
	}
	if (uid.value.type == 2) {
		wkStyle.value = `box-sizing: border-box;border-bottom:1rpx solid ${tmcomputed.value.border};border-right:1rpx solid ${
			tmcomputed.value.border
		};border-left:1rpx solid rgba(0,0,0,0);border-top:1rpx solid rgba(0,0,0,0);width:${_colWidth - 1}rpx`
	}
	if (uid.value.type == 3) {
		wkStyle.value = `box-sizing: border-box;border-top:1rpx solid rgba(0,0,0,0);border-bottom:1rpx solid ${
			tmcomputed.value.border
		};border-right:1rpx solid ${tmcomputed.value.border};border-left:1rpx solid ${tmcomputed.value.border};width:${_colWidth - 1}rpx`
	}
	if (uid.value.type == 4) {
		wkStyle.value = `box-sizing: border-box;border-left:1rpx solid rgba(0,0,0,0);border-bottom:1rpx solid ${
			tmcomputed.value.border
		};border-top:1rpx solid ${tmcomputed.value.border};border-right:1rpx solid ${tmcomputed.value.border};width:${_colWidth - 1}rpx`
	}
}
function onClick(e: Event) {
	emits('click', e)
	if (props.url !== '') {
		try {
			uni.navigateTo({
				url: props.url
			})
		} catch (e) {
			//TODO handle the exception
		}
	}
}
</script>
